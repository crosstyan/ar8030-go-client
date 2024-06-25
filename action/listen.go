package action

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"context"
	"errors"
	"github.com/joomcode/errorx"
	"io"
	"net"
)

// SubscribeHotPlugMoveConn will TAKE OVER the connection (conn) i.e. the ownership of the connection is transferred to this function
func SubscribeHotPlugMoveConn(conn *net.TCPConn, ctx context.Context) (<-chan bb.UsbPack, error) {
	var TCPReadBufferSize = 4096
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
				// TODO: we don't need such a big buffer
				// https://mostafa.dev/why-do-tcp-connections-in-go-get-stuck-reading-large-amounts-of-data-f490a26a605e
				rBuf := make([]byte, TCPReadBufferSize)
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

type SubscribedMessage struct {
	Event   bb.Event
	WorkId  uint32
	Payload *bb.UsbPack
}

// SubscribeMessageWithNewConn will subscribe to a message to current connection
// see `cb_bb_ioctl` and `create_new_cb` in `session_callback.c`
//
// Please Note that this function will create a new connection based on the passed argument
// The original will not be used or closed
func SubscribeMessageWithNewConn(conn *net.TCPConn, ctx context.Context, workId uint32, event bb.Event) (<-chan SubscribedMessage, error) {
	var TCPReadBufferSize = 1024 * 16
	reqId := bb.SubscribeRequestId(event)
	pack := bb.UsbPack{
		MsgId: workId,
		ReqId: reqId,
		Sta:   0,
	}
	wbBuf := bb.NewBuffer(32)
	err := pack.Write(wbBuf)
	if err != nil {
		return nil, err
	}
	newConn, err := bb.NewTCPFromConn(conn)
	if err != nil {
		return nil, err
	}
	_, err = newConn.Write(wbBuf.Bytes())
	if err != nil {
		_ = newConn.Close()
		return nil, err
	}
	ch := make(chan SubscribedMessage)
	go func(conn *net.TCPConn, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				err = conn.Close()
				if err != nil {
					log.Sugar().Errorw("failed to close connection", "error", err.Error())
				}
				close(ch)
				return
			default:
				rBuf := make([]byte, TCPReadBufferSize)
				_, err := conn.Read(rBuf)
				if err != nil {
					if errors.Is(err, io.EOF) {
						log.Sugar().Infow("connection closed", "id", workId, "event", event)
						close(ch)
					} else {
						log.Sugar().Errorw("failed to read from connection, closing channel", "id", workId, "event", event, "error", err.Error())
						err = conn.Close()
						if err != nil {
							log.Sugar().Errorw("failed to close connection", "id", workId, "event", event, "error", err.Error())
						}
						close(ch)
					}
					return
				}
				rbBuf := bytes.NewBuffer(rBuf)
				resp := &bb.UsbPack{}
				err = resp.Read(rbBuf)
				if err != nil {
					log.Sugar().Errorw("failed to unmarshal response", "id", workId, "event", event, "error", err.Error(), "payload", rBuf)
					continue
				}
				m := SubscribedMessage{
					Event:   event,
					Payload: resp,
					WorkId:  workId,
				}
				ch <- m
			}
		}
	}(newConn, ctx)
	return ch, nil
}
