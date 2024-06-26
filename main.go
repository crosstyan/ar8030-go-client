package main

import (
	"ar8030/action"
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"context"
	"fmt"
	"github.com/joomcode/errorx"
	"net"
	"os"
	"time"
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
	// port 0 & 1 can't open both (-259)
	// port 2 can only be open at DEV role, AP can't read from it (-259)
	// port 3 seems to work
	var comPort byte = 3
	recvLoop := func(sktConn net.Conn, selId uint32, role bb.Role) {
		for {
			rBuf := make([]byte, 2048)
			rbBuf := bytes.NewBuffer(rBuf)
			_, err := sktConn.Read(rBuf)
			if err != nil {
				log.Sugar().Errorw("failed to read from socket", "id", selId, "role", role, "error", err.Error())
				return
			}
			pack := bb.UsbPack{}
			err = pack.Read(rbBuf)
			if err != nil {
				log.Sugar().Errorw("failed to read pack", "id", selId, "role", role, "error", err.Error())
				return
			}
			hdr := action.SockHeaderFromReqId(pack.ReqId)
			if hdr.Opt == bb.SoRead {
				msg, _ := action.UnwrapSocketRx(&pack)
				if msg != nil {
					log.Sugar().Infow("socket rx", "id", selId, "message", string(msg.Payload))
				}
			}
		}
	}
	go func() {
		if st.Role == bb.BB_ROLE_DEV {
			sktConn, err := bb.NewTCPFromConn(conn.(*net.TCPConn))
			defer func() {
				log.Sugar().Debugw("closing socket connection", "id", selId, "role", st.Role)
				err = sktConn.Close()
				if err != nil {
					log.Sugar().Errorw("failed to close connection", "error", err.Error())
				}
			}()
			if err != nil {
				log.Sugar().Errorw("failed to create socket connection", "id", selId, "error", err.Error())
				return
			}
			var slot = bb.BB_SLOT_AP
			var port byte = comPort
			err = action.OpenSocket(sktConn, selId, slot, port, bb.BB_SOCK_FLAG_TX|bb.BB_SOCK_FLAG_RX, nil)
			if err != nil {
				log.Sugar().Errorw("failed to open socket", "id", selId, "error", err.Error())
				return
			}
			log.Sugar().Infow("socket opened", "id", selId, "slot", slot, "port", port, "role", st.Role)
			go recvLoop(sktConn, selId, st.Role)
			for {
				var m = []byte("hello from dev")
				err = action.WriteSocket(sktConn, selId, slot, port, m)
				if err != nil {
					log.Sugar().Errorw("failed to write to socket", "id", selId, "role", st.Role, "error", err.Error())
				}
				log.Sugar().Infow("socket tx", "id", selId, "role", st.Role, "message", string(m))
				time.Sleep(1 * time.Second)
			}
		} else if st.Role == bb.BB_ROLE_AP {
			sktConn, err := bb.NewTCPFromConn(conn.(*net.TCPConn))
			defer func() {
				log.Sugar().Debugw("closing socket connection", "id", selId, "role", st.Role)
				err = sktConn.Close()
				if err != nil {
					log.Sugar().Errorw("failed to close connection", "error", err.Error())
				}
			}()
			if err != nil {
				log.Sugar().Errorw("failed to create socket connection", "id", selId, "error", err.Error())
				return
			}
			var slot = bb.BB_SLOT_0
			var port byte = comPort
			err = action.OpenSocket(sktConn, selId, slot, port, bb.BB_SOCK_FLAG_TX|bb.BB_SOCK_FLAG_RX, nil)
			if err != nil {
				log.Sugar().Errorw("failed to open socket", "id", selId, "error", err.Error())
				return
			}
			log.Sugar().Infow("socket opened", "id", selId, "slot", slot, "port", port, "role", st.Role)
			go recvLoop(sktConn, selId, st.Role)
			for {
				var m = []byte("hello from ap")
				err = action.WriteSocket(sktConn, selId, slot, port, m)
				if err != nil {
					log.Sugar().Errorw("failed to write to socket", "id", selId, "role", st.Role, "error", err.Error())
				}
				log.Sugar().Infow("socket tx", "id", selId, "role", st.Role, "message", string(m))
				time.Sleep(2*time.Second + 500*time.Millisecond)
			}
		}
	}()
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
	subChs := make([]<-chan action.SubscribedMessage, 0)
	// handleDevice prints device info and subscribes to events
	handleDevice := func(selId uint32) {
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
		ch1, err := action.SubscribeMessageWithNewConn(conn, ctx, selId, bb.BB_EVENT_LINK_STATE)
		logErr(err, selId, bb.BB_EVENT_LINK_STATE)
		ch2, err := action.SubscribeMessageWithNewConn(conn, ctx, selId, bb.BB_EVENT_MCS_CHANGE)
		logErr(err, selId, bb.BB_EVENT_MCS_CHANGE)
		ch3, err := action.SubscribeMessageWithNewConn(conn, ctx, selId, bb.BB_EVENT_CHAN_CHANGE)
		logErr(err, selId, bb.BB_EVENT_CHAN_CHANGE)
		ch4, err := action.SubscribeMessageWithNewConn(conn, ctx, selId, bb.BB_EVENT_OFFLINE)
		logErr(err, selId, bb.BB_EVENT_OFFLINE)
		subChs = append(subChs, ch1, ch2, ch3, ch4)
	}

	for _, selId := range wrkList {
		handleDevice(selId)
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
	ch, err := action.SubscribeHotPlugMoveConn(hotPlugConn, ctx)
	if err != nil {
		log.Sugar().Panicw("failed to subscribe hot plug event", "error", err.Error())
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
