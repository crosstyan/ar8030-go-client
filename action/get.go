package action

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/joomcode/errorx"
	"net"
)

func GetStatus(workId uint32, conn net.Conn) (*bb.GetStatusOut, error) {
	iStatus := bb.GetStatusIn{
		UserBmp: 1,
	}
	var err error
	buf_ := make([]byte, 0, 32)
	buf := bytes.NewBuffer(buf_)
	err = binary.Write(buf, binary.NativeEndian, iStatus)
	if err != nil {
		return nil, err
	}
	log.Sugar().Infow("get status", "iStatus", hex.EncodeToString(buf.Bytes()))
	pack := &bb.UsbPack{
		Buf:   buf.Bytes(),
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_GET_STATUS,
	}
	opt := RequestOption{
		RequestBufferSize: 32,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return nil, err
	}
	if resp.Sta != 0 {
		return nil, errorx.ExternalError.New("non-zero status %d", resp.Sta)
	}
	log.Sugar().Infow("get status", "resp", hex.EncodeToString(resp.Buf))
	if resp.Sta < 0 {
		return nil, errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	rBuf := bytes.NewBuffer(resp.Buf)
	oStaus := bb.GetStatusOut{}
	err = binary.Read(rBuf, binary.NativeEndian, &oStaus)
	if err != nil {
		return nil, err
	}
	return &oStaus, nil
}
