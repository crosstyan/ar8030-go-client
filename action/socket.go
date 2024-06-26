package action

import (
	"ar8030/bb"
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"net"
)

// OpenSocket implements `bb_socket_open`
//
//   - slot 目标SLOT, 如果为DEV，目标SLOT均为BB_SLOT_AP, See bb.Slot
//   - port 逻辑端口，不同端口的数据互相独立，port 的数量受 bb.BB_CONFIG_MAX_TRANSPORT_PER_SLOT 控制
//   - flags 传输标志，可以是 bb.BB_SOCK_FLAG_TX, bb.BB_SOCK_FLAG_RX, bb.BB_SOCK_FLAG_DATAGRAMSockFlag
func OpenSocket(conn net.Conn, workId uint32, slot bb.Slot, port byte, flags bb.SockFlag, opt *bb.SockOpt) error {
	defaultOpt := bb.SockOpt{
		TxBufSize: bb.BB_CONFIG_MAC_TX_BUF_SIZE,
		RxBufSize: bb.BB_CONFIG_MAC_RX_BUF_SIZE,
	}
	merge := func(old bb.SockOpt, new *bb.SockOpt) bb.SockOpt {
		if new == nil {
			return old
		}
		if new.TxBufSize != 0 {
			old.TxBufSize = new.TxBufSize
		}
		if new.RxBufSize != 0 {
			old.RxBufSize = new.RxBufSize
		}
		return old
	}
	o := merge(defaultOpt, opt)
	// actually only needs 12
	wBuf := make([]byte, 0, 16)
	wbBuf := bytes.NewBuffer(wBuf)
	err := binary.Write(wbBuf, binary.LittleEndian, flags)
	if err != nil {
		return err
	}
	err = binary.Write(wbBuf, binary.LittleEndian, o)
	if err != nil {
		return err
	}
	pack := bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.SocketRequestId(bb.SoOpen, slot, port),
		Buf:   wbBuf.Bytes(),
	}
	resp, err := RequestWithPack(conn, &pack, nil)
	if err != nil {
		return err
	}
	if resp.Sta != 0 {
		return errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	return nil
}

// WriteSocket implements `bb_socket_write`
//
//   - slot 目标SLOT, 如果为DEV，目标SLOT均为BB_SLOT_AP, See bb.Slot
//   - port 逻辑端口，不同端口的数据互相独立，port 的数量受配置宏 bb.BB_CONFIG_MAX_TRANSPORT_PER_SLOT 控制
//
// note that the socket will reply with `ReqId=0x04010003` (which has `bb.SoCmdOpt`=bb.SoRead)
// and its Sta is the number of bytes written (but buffer will be empty for sure)
// This function WON'T check the response and just fire the request
// The response should be handled by the caller (usually another goroutine for reading from the socket)
//
// TODO: find out what `datagram` means
// TODO: check for payload length (See bb.SockOpt)
// See also bb.BB_CONFIG_MAC_TX_BUF_SIZE and bb.BB_CONFIG_MAC_RX_BUF_SIZE
func WriteSocket(conn net.Conn, workId uint32, slot bb.Slot, port byte, payload []byte) error {
	pack := bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.SocketRequestId(bb.SoWrite, slot, port),
		Buf:   payload,
	}
	buf_ := make([]byte, 0, 32+len(payload))
	buf := bytes.NewBuffer(buf_)
	err := pack.Write(buf)
	if err != nil {
		return err
	}
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

type RxMessage struct {
	Slot    bb.Slot
	Port    byte
	Payload []byte
}

type SockHeader struct {
	Opt  bb.SoCmdOpt
	Slot bb.Slot
	Port byte
}

func SockHeaderFromReqId(reqId bb.RequestId) SockHeader {
	return SockHeader{
		Opt:  bb.SoCmdOpt(reqId >> 16 & 0xff),
		Slot: bb.Slot(reqId >> 8 & 0xff),
		Port: byte(reqId & 0xff),
	}
}

// UnwrapSocketRx implements `so_rpc_cb`
func UnwrapSocketRx(pack *bb.UsbPack) (*RxMessage, error) {
	h := SockHeaderFromReqId(pack.ReqId)
	opt, slot, port := h.Opt, h.Slot, h.Port
	if opt != bb.SoRead {
		return nil, errorx.IllegalState.New("not a read command; opt=0x%02x; reqid=0x%08x", opt, pack.ReqId)
	}
	return &RxMessage{
		Slot:    slot,
		Port:    port,
		Payload: pack.Buf,
	}, nil
}

// CloseSocket implements `bb_socket_close` but actually does nothing
// since you only needs to close the underlying connection
func CloseSocket(conn net.Conn, workId uint32, slot bb.Slot, port byte) error {
	return nil
}

// SocketComOpt implements `bb_socket_com_opt`
// TODO: implement this
// see also
//   - `bb_socket_ioctl`
//   - `bb_socket_com_opt`
//   - `so_rpc_cb`
//   - `BASE_FUN` (`sofun`) .rdcb
//   - `BASE_SESSION` (what kind of session we're talking about?)
//   - bb_read_thread (interesting)
func SocketComOpt(conn net.Conn, workId uint32, slot bb.Slot, port byte, opt bb.SoCmdOpt) error {
	// var reqId = bb.SocketRequestId(opt, slot, port)
	return errorx.NotImplemented.New("SocketComOpt")
}
