package action

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"net"
)

func GetStatus(conn net.Conn, workId uint32, userBmp uint16) (*bb.GetStatusOut, error) {
	iStatus := bb.GetStatusIn{
		UserBmp: userBmp,
	}
	var err error
	buf_ := make([]byte, 0, 32)
	buf := bytes.NewBuffer(buf_)
	err = binary.Write(buf, binary.LittleEndian, iStatus)
	if err != nil {
		return nil, err
	}
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
	if resp.Sta < 0 {
		return nil, errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	oStaus := bb.UnsafeFromByteSlice[bb.GetStatusOut](resp.Buf)
	return &oStaus, nil
}

func GetSysInfo(conn net.Conn, workId uint32) (*bb.GetSysInfoOut, error) {
	pack := &bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_GET_SYS_INFO,
	}
	opt := RequestOption{
		RequestBufferSize: 32,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return nil, err
	}
	if resp.Sta < 0 {
		return nil, errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	oSysInfo := bb.UnsafeFromByteSlice[bb.GetSysInfoOut](resp.Buf)
	return &oSysInfo, nil
}

// GetCfg fetches the configuration from the device, but only with one segment
func GetCfg(conn net.Conn, workId uint32, seq uint16, offset uint16) (*bb.GetCfgOut, error) {
	ip := &bb.GetCfgIn{
		Seq:    seq,
		Mode:   0,
		Offset: offset,
		Length: bb.GetCfgInMaxLength,
	}
	var err error
	buf_ := make([]byte, 0, 32)
	buf := bytes.NewBuffer(buf_)
	err = binary.Write(buf, binary.LittleEndian, ip.Seq)
	err = binary.Write(buf, binary.LittleEndian, ip.Mode)
	err = binary.Write(buf, binary.LittleEndian, byte(0)) // padding
	err = binary.Write(buf, binary.LittleEndian, ip.Offset)
	err = binary.Write(buf, binary.LittleEndian, ip.Length)
	if err != nil {
		return nil, err
	}
	ipBuf := buf.Bytes()
	pack := &bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_GET_CFG,
		Buf:   ipBuf,
	}
	opt := RequestOption{
		ResponseBufferSize: 2048,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return nil, err
	}
	if resp.Sta < 0 {
		return nil, errorx.ExternalError.New("negative status %d", resp.Sta)
	}
	oCfg := bb.UnsafeFromByteSlice[bb.GetCfgOut](resp.Buf)
	return &oCfg, nil
}

// GetFullCfg fetches the full configuration from the device
func GetFullCfg(conn net.Conn, workId uint32) ([]byte, error) {
	var cfgBuf []byte
	var offset uint16 = 0
	var seqNum uint16 = 0
	cfg, err := GetCfg(conn, workId, seqNum, offset)
	if err != nil {
		return nil, err
	}
	total := int(cfg.TotalLength)
	offset += cfg.Length
	seqNum += 1
	log.Sugar().Debugw("total cfg length", "length", total)
	// totalCrc16, not used
	_ = cfg.TotalCrc16
	findNull := func(data []byte) []byte {
		i := bytes.IndexByte(data, 0x00)
		if i == -1 {
			return data
		} else {
			return data[:i]
		}
	}
	a := findNull(cfg.Data[:])
	cfgBuf = append(cfgBuf, a...)
	for i := bb.GetCfgInMaxLength; i < total; i += bb.GetCfgInMaxLength {
		cfg, err = GetCfg(conn, workId, seqNum, offset)
		if err != nil {
			return nil, err
		}
		a = findNull(cfg.Data[:])
		// discard overflown data
		//
		// the daemon (server) should have rejected my request
		// or discard the data
		if cfg.Length < bb.GetCfgInMaxLength {
			a = a[:cfg.Length]
		}
		cfgBuf = append(cfgBuf, a...)
		if len(cfgBuf) >= total {
			break
		}
		offset += cfg.Length
		seqNum += 1
	}
	return cfgBuf, nil
}
