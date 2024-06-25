package action

import (
	"ar8030/bb"
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"net"
)

// OpenSocket implements `bb_socket_open`
// note that flags are bb.BB_SOCK_FLAG_TX, bb.BB_SOCK_FLAG_RX and bb.BB_SOCK_FLAG_DATAGRAM (not sure about DATAGRAM)
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
// slot 目标SLOT, 如果为DEV，目标SLOT均为BB_SLOT_AP, See also bb.Slot
// port 逻辑端口，不同端口的数据互相独立，port 的数量受配置宏 bb.BB_CONFIG_MAX_TRANSPORT_PER_SLOT 控制
//
// TODO: find out what `datagram` means
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

func CloseSocket(conn net.Conn, workId uint32, slot bb.Slot, port byte) error {
	// dummy implementation
	return nil
}

// SocketComOpt implements `bb_socket_com_opt`
// see also
//   - `bb_socket_ioctl`
//   - `bb_socket_com_opt`
//   - `so_rpc_cb`
//   - `BASE_FUN` (`sofun`) .rdcb
//   - `BASE_SESSION` (what kind of session we're talking about?)
//   - bb_read_thread (interesting)
func SocketComOpt(conn net.Conn, workId uint32, slot bb.Slot, port byte, opt *bb.SockOpt) error {
	// dummy implementation
	return nil
}
