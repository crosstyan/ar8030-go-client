package action

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/joomcode/errorx"
	"net"
	"unsafe"
)

// GetStatus implements bb.BB_GET_STATUS, fetching the status of the device
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
		return nil, errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	oStaus := bb.UnsafeFromByteSlice[bb.GetStatusOut](resp.Buf)
	return &oStaus, nil
}

// GetSysInfo implements bb.BB_GET_SYS_INFO, fetching the system information
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
		return nil, errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	oSysInfo := bb.UnsafeFromByteSlice[bb.GetSysInfoOut](resp.Buf)
	return &oSysInfo, nil
}

// GetPairResult implements bb.BB_GET_PAIR_RESULT, fetching the pairing result
func GetPairResult(conn net.Conn, workId uint32) (*bb.GetPairResultOut, error) {
	pack := &bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_GET_PAIR_RESULT,
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
	oPairResult := bb.UnsafeFromByteSlice[bb.GetPairResultOut](resp.Buf)
	return &oPairResult, nil
}

// GetCfg implements bb.BB_GET_CFG, only fetching a part of the configuration
func GetCfg(conn net.Conn, workId uint32, seq uint16, offset uint16, length uint16) (*bb.GetCfgOut, error) {
	if length > bb.GetCfgInMaxLength {
		return nil, errorx.IllegalArgument.New("length %d exceeds the maximum length %d", length, bb.GetCfgInMaxLength)
	}
	ip := &bb.GetCfgIn{
		Seq:    seq,
		Mode:   0,
		Offset: offset,
		Length: length,
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
		return nil, errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	oCfg := bb.UnsafeFromByteSlice[bb.GetCfgOut](resp.Buf)
	return &oCfg, nil
}

// GetFullCfg fetches the full configuration from the device
func GetFullCfg(conn net.Conn, workId uint32) ([]byte, error) {
	var cfgBuf []byte
	var offset uint16 = 0
	var seqNum uint16 = 0
	var BatchSize uint16 = bb.GetCfgInMaxLength
	cfg, err := GetCfg(conn, workId, seqNum, offset, BatchSize)
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
	// just in case the data we got is less than the maximum length
	// it's very unlikely to happen
	if cfg.TotalLength < BatchSize {
		a := cfg.Data[:cfg.TotalLength]
		return findNull(a), nil
	}
	a := findNull(cfg.Data[:])
	cfgBuf = append(cfgBuf, a...)
	for i := int(BatchSize); i < total; i += bb.GetCfgInMaxLength {
		l := func() uint16 {
			ll := total - i
			if ll > int(BatchSize) {
				return BatchSize
			} else {
				return uint16(ll)
			}
		}()
		cfg, err = GetCfg(conn, workId, seqNum, offset, l)
		if err != nil {
			return nil, err
		}
		// discard overflown data
		if cfg.Length < BatchSize {
			a = a[:cfg.Length]
		}
		a = findNull(cfg.Data[:])
		cfgBuf = append(cfgBuf, a...)
		if len(cfgBuf) >= total {
			break
		}
		offset += cfg.Length
		seqNum += 1
	}
	return cfgBuf, nil
}

// GetMac implements the BB_RPC_GET_MAC to get the DevInfo of certain workId,
// which contains the MAC address of the device.
func GetMac(conn net.Conn, workId uint32) ([]byte, error) {
	pack := &bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_RPC_GET_MAC,
	}
	opt := RequestOption{
		ResponseBufferSize: 128,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return nil, err
	}
	if resp.Sta < 0 {
		return nil, errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	return resp.Buf, nil
}

func QuerySockBufSta(conn net.Conn, workId uint32, slot int32, port int32) (*bb.QueryTxOut, error) {
	ip := &bb.QueryTxIn{
		Slot: slot,
		Port: port,
	}
	// you can't use `int` directly for `binary.Write`, that's weird
	buf := bb.NewBuffer(32)
	err := binary.Write(buf, binary.NativeEndian, ip.Slot)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.NativeEndian, ip.Port)
	if err != nil {
		return nil, err
	}
	log.Sugar().Debugw("query sock buf sta", "ip", ip, "buf", hex.EncodeToString(buf.Bytes()))
	pack := &bb.UsbPack{
		Buf:   buf.Bytes(),
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_RPC_SOCK_BUF_STA,
	}
	opt := RequestOption{
		RequestBufferSize:  64,
		ResponseBufferSize: 64,
	}
	resp, err := RequestWithPack(conn, pack, &opt)
	if err != nil {
		return nil, err
	}
	if resp.ReqId != bb.BB_RPC_SOCK_BUF_STA {
		return nil, errorx.ExternalError.New("unexpected request id %d", resp.ReqId)
	}
	if resp.Sta < 0 {
		if resp.Sta == -1 {
			// no idea how this work
			return nil, errorx.ExternalError.New("no such work node. see `buf_query` in `rpc_debug.c` for more information.")
		}
		return nil, errorx.ExternalError.New("bad status %d", resp.Sta)
	}
	expected := int(unsafe.Sizeof(bb.QueryTxOut{}))
	if len(resp.Buf) < expected {
		return nil, errorx.ExternalError.New("unexpected response length %d, expected %d", len(resp.Buf), expected)
	}
	o := bb.UnsafeFromByteSlice[bb.QueryTxOut](resp.Buf)
	return &o, nil
}
