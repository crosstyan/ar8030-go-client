package main

import (
	"ar8030/action"
	"ar8030/bb"
	"ar8030/log"
	"context"
	"fmt"
	"github.com/joomcode/errorx"
	"net"
	"os"
)

const (
	HOST = "127.0.0.1"
	PORT = 50000
)

func printDeviceInfo(conn net.Conn, selId uint32) error {
	mac, err := action.GetMac(conn, selId)
	if err != nil {
		return errorx.ExternalError.Wrap(err, "failed to get MAC")
	}
	log.Sugar().Infow("MAC", "id", selId, "mac", bb.MacLike(mac, ":"))
	// 0x3fff is a magic value, no idea why it's chosen
	st, err := action.GetStatus(conn, selId, 0x3fff)
	if err != nil {
		return errorx.ExternalError.Wrap(err, "failed to get status")
	}
	log.Sugar().Infow("status", "id", selId, "status", st)
	sysInfo, err := action.GetSysInfo(conn, selId)
	if err != nil {
		return errorx.ExternalError.Wrap(err, "failed to get system info")
	}
	log.Sugar().Infow("system info", "id", selId, "info", sysInfo)
	return nil
}

func main() {
	var err error
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", HOST, PORT))
	if err != nil {
		log.Sugar().Panicw("failed to resolve TCP address", "error", err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Sugar().Panicw("failed to close connection", "error", err.Error())
		}
	}(conn)
	if err != nil {
		log.Sugar().Panicw("failed to dial TCP", "error", err.Error())
	}
	err = action.TestServer(conn)
	if err != nil {
		log.Sugar().Panicw("failed to test server", "error", err.Error())
	}
	wrkList, err := action.GetWorkIdList(conn)
	if err != nil {
		log.Sugar().Panicw("failed to get work id list", "error", err.Error())
	}
	log.Sugar().Infow("work id list", "list", wrkList)

	ctx := context.Background()
	if err != nil {
		log.Sugar().Panicw("failed to dial TCP", "error", err.Error())
	}
	subChs := make([]<-chan action.SubscribedMessage, 0)
	handleDev := func(selId uint32) {
		conn, err := bb.NewTCPFromConn(conn)
		defer func() {
			log.Sugar().Debugw("closing query connection", "id", selId)
			err = conn.Close()
			if err != nil {
				log.Sugar().Errorw("failed to close connection", "error", err.Error())
			}
		}()
		if err != nil {
			log.Sugar().Panicw("failed to dial TCP", "error", err.Error())
		}
		// You can only select one workId for a single connection
		// which is stupid, but it's how the daemon implemented
		ok, err := action.SelectWorkId(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to test work id", "error", err.Error(), "id", selId)
		}
		if !ok {
			log.Sugar().Errorw("work id is invalid", "id", selId)
			os.Exit(1)
		}
		err = printDeviceInfo(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to print device info", "error", err.Error())
		}
		logErr := func(err error, selId uint32, event bb.Event) {
			if err != nil {
				log.Sugar().Errorw("failed to subscribe message", "id", selId, "error", err.Error(), "event", event)
			} else {
				log.Sugar().Infow("event subscribed", "id", selId, "event", event)
			}
		}
		ch1, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_LINK_STATE)
		logErr(err, selId, bb.BB_EVENT_LINK_STATE)
		ch2, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_MCS_CHANGE)
		logErr(err, selId, bb.BB_EVENT_MCS_CHANGE)
		ch3, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_CHAN_CHANGE)
		logErr(err, selId, bb.BB_EVENT_CHAN_CHANGE)
		ch4, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_OFFLINE)
		logErr(err, selId, bb.BB_EVENT_OFFLINE)
		subChs = append(subChs, ch1, ch2, ch3, ch4)
	}

	for _, selId := range wrkList {
		handleDev(selId)
	}

	chs, _ := bb.MergeChannels(subChs...)
	go func(ch <-chan action.SubscribedMessage, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Sugar().Debug("context done")
				return
			case msg, ok := <-ch:
				if !ok {
					return
				}
				log.Sugar().Infow("subscribed message", "message", msg)
			}
		}
	}(chs, ctx)

	hotPlugConn, err := bb.NewTCPFromConn(conn)
	if err != nil {
		log.Sugar().Panicw("failed to dial TCP", "error", err.Error())
	}
	ch, err := action.MoveRegisterHotPlug(hotPlugConn, ctx)
	if err != nil {
		log.Sugar().Panicw("failed to register hot plug", "error", err.Error())
	}
	for {
		pack, ok := <-ch
		if !ok {
			break
		}
		handleHotPlugEvent := func(pack *bb.UsbPack) {
			if pack.ReqId != bb.BB_RPC_GET_HOTPLUG_EVENT {
				log.Sugar().Warnw("unexpected request id", "id", pack.ReqId)
				return
			}
			// note that the bb.EventHotPlug.Id is the same as workId (selId)
			// workId will change, MacAddr will change, but DevInfo.Mac will not change
			resp := bb.UnsafeFromByteSlice[bb.EventHotPlug](pack.Buf)
			log.Sugar().Infow("hot plug event", "mac", resp.BbMac.String(), "event", resp)
		}
		handleHotPlugEvent(&pack)
	}
}
