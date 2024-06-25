package main

import (
	"ar8030/action"
	"ar8030/bb"
	"ar8030/log"
	"context"
	"fmt"
	"net"
	"os"
)

const (
	HOST = "127.0.0.1"
	PORT = 50000
)

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
	for _, selId := range wrkList {
		ok, err := action.SelectWorkId(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to test work id", "error", err.Error(), "id", selId)
		}
		if !ok {
			log.Sugar().Errorw("work id is invalid", "id", selId)
			os.Exit(1)
		}
		mac, err := action.GetMac(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to get MAC", "id", selId, "error", err.Error())
		}
		log.Sugar().Infow("MAC", "id", selId, "mac", bb.MacLike(mac, ":"))
		logErr := func(err error, event bb.Event) {
			if err != nil {
				log.Sugar().Errorw("failed to subscribe message", "error", err.Error(), "event", event)
			}
		}
		ch1, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_LINK_STATE)
		logErr(err, bb.BB_EVENT_LINK_STATE)
		ch2, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_MCS_CHANGE)
		logErr(err, bb.BB_EVENT_MCS_CHANGE)
		ch3, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_CHAN_CHANGE)
		logErr(err, bb.BB_EVENT_CHAN_CHANGE)
		ch4, err := action.SubscribeMessage(conn, ctx, selId, bb.BB_EVENT_OFFLINE)
		logErr(err, bb.BB_EVENT_OFFLINE)
		subChs = append(subChs, ch1, ch2, ch3, ch4)
	}

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

	// useless... for now
	_ = func(selId uint32) {
		// 0x3fff is a magic value, no idea why it's chosen
		st, err := action.GetStatus(conn, selId, 0x3fff)
		if err != nil {
			log.Sugar().Panicw("failed to get status", "error", err.Error())
		}
		log.Sugar().Infow("status", "status", st)
		sysInfo, err := action.GetSysInfo(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to get system info", "error", err.Error())
		}
		log.Sugar().Infow("system info", "info", sysInfo)

		oCfg, err := action.GetFullCfg(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to get configuration", "error", err.Error())
		}
		sCfg := string(oCfg)
		log.Sugar().Infow("configuration", "len", len(sCfg))
		f, err := os.Create("config.json")
		if err != nil {
			log.Sugar().Panicw("failed to create file", "error", err.Error())
		}
		defer func(f *os.File) {
			err = f.Close()
			if err != nil {
				log.Sugar().Panicw("failed to close file", "error", err.Error())
			}
		}(f)
		_, err = f.WriteString(sCfg)
		if err != nil {
			log.Sugar().Panicw("failed to write to file", "error", err.Error())
		}
	}
}
