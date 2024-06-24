package action

import (
	"ar8030/bb"
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"net"
)

// SetPairMode implements the action of bb.BB_SET_PAIR_MODE
// entering or exiting the pairing mode for `workId` device
// Note that you could poll the status by GetStatus
func SetPairMode(conn net.Conn, workId uint32, ip bb.SetPairModeIn) error {
	buf_ := make([]byte, 0, 32)
	buf := bytes.NewBuffer(buf_)
	err := binary.Write(buf, binary.LittleEndian, ip)
	pack := &bb.UsbPack{
		Buf:   buf.Bytes(),
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_SET_PAIR_MODE,
	}
	opt := RequestOption{
		RequestBufferSize:  64,
		ResponseBufferSize: 32,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return err
	}
	if resp.Sta < 0 {
		return errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	return nil
}

// SetApMac implements the action of bb.BB_SET_AP_MAC
// Note that only the DEV role can set the AP MAC address.
func SetApMac(conn net.Conn, workId uint32, ip bb.SetApMacIn) error {
	buf_ := make([]byte, 0, 32)
	buf := bytes.NewBuffer(buf_)
	err := binary.Write(buf, binary.LittleEndian, ip)
	pack := &bb.UsbPack{
		Buf:   buf.Bytes(),
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_SET_AP_MAC,
	}
	opt := RequestOption{
		RequestBufferSize:  64,
		ResponseBufferSize: 32,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return err
	}
	if resp.Sta < 0 {
		return errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	return nil
}

// SetCandidates implements the action of bb.BB_SET_CANDIDATES
// Note that only the AP role can set the candidates.
func SetCandidates(conn net.Conn, workId uint32, ip bb.SetCandidatesIn) error {
	buf_ := make([]byte, 0, 32)
	buf := bytes.NewBuffer(buf_)
	err := binary.Write(buf, binary.LittleEndian, ip)
	pack := &bb.UsbPack{
		Buf:   buf.Bytes(),
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_CFG_CANDIDATES,
	}
	opt := RequestOption{
		RequestBufferSize:  64,
		ResponseBufferSize: 32,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return err
	}
	if resp.Sta < 0 {
		return errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	return nil
}
