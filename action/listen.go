package action

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"context"
	"github.com/joomcode/errorx"
	"net"
)

// MoveRegisterHotPlug will TAKE OVER the connection (conn) i.e. the ownership of the connection is transferred to this function
func MoveRegisterHotPlug(conn *net.TCPConn, ctx context.Context) (<-chan bb.UsbPack, error) {
	wbBuf := bb.NewBuffer(32)
	pack := bb.UsbPack{
		MsgId: 0,
		ReqId: bb.BB_RPC_GET_HOTPLUG_EVENT,
		Sta:   0,
	}
	err := pack.Write(wbBuf)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write(wbBuf.Bytes())
	if err != nil {
		return nil, err
	}
	rBuf := make([]byte, 32)
	sz, err := conn.Read(rBuf)
	if err != nil {
		return nil, err
	}
	rBuf = rBuf[:sz]
	rbBuf := bytes.NewBuffer(rBuf)
	resp := bb.UsbPack{}
	err = resp.Read(rbBuf)
	if err != nil {
		return nil, err
	}
	if resp.Sta != 0 {
		return nil, errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	ch := make(chan bb.UsbPack)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Sugar().Debug("context done")
				close(ch)
				return
			default:
				rBuf := make([]byte, 4096)
				_, err := conn.Read(rBuf)
				if err != nil {
					log.Sugar().Errorw("failed to read from connection", "error", err.Error())
					close(ch)
					return
				}
				rbBuf := bytes.NewBuffer(rBuf)
				resp := bb.UsbPack{}
				err = resp.Read(rbBuf)
				if err != nil {
					log.Sugar().Errorw("failed to unmarshal response", "error", err.Error())
					close(ch)
					return
				}
				ch <- resp
			}
		}
	}()
	return ch, nil
}
