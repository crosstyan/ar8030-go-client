package action

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
	"net"
)

/**
 * See `cmdtab` (`BBIOCTL_TAB`) in `ioctl_tab.c`
 */

type RequestOption struct {
	RequestBufferSize  uint
	ResponseBufferSize uint
}

func mergeRequestOption(old RequestOption, new *RequestOption) RequestOption {
	if new == nil {
		return old
	}
	if new.RequestBufferSize != 0 {
		old.RequestBufferSize = new.RequestBufferSize
	}
	if new.ResponseBufferSize != 0 {
		old.ResponseBufferSize = new.ResponseBufferSize
	}
	return old
}

func RequestWithPack(conn net.Conn, pack *bb.UsbPack, option *RequestOption) (*bb.UsbPack, error) {
	defaultOption := RequestOption{
		RequestBufferSize:  1024,
		ResponseBufferSize: 1024,
	}
	o := mergeRequestOption(defaultOption, option)
	// It can also be used to set the initial size of the internal buffer for writing.
	// To do that, buf should have the desired capacity but a length of zero.
	buf_ := make([]byte, 0, o.RequestBufferSize)
	buf := bytes.NewBuffer(buf_)
	err := pack.Write(buf)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	buf_ = make([]byte, o.ResponseBufferSize)
	sz, err := conn.Read(buf_)
	buf = bytes.NewBuffer(buf_[:sz])
	err = pack.Read(buf)
	if err != nil {
		return nil, err
	}
	return pack, nil
}

func TestServer(conn net.Conn) error {
	opt := RequestOption{
		RequestBufferSize:  32,
		ResponseBufferSize: 32,
	}
	req := bb.UsbPack{
		MsgId: 0,
		Sta:   0,
		ReqId: bb.BB_RPC_TEST,
	}
	resp, err := RequestWithPack(conn, &req, &opt)
	if err != nil {
		return err
	}
	if resp.Sta != 0 {
		return errorx.ExternalError.New("bad status %d", req.Sta)
	}
	return nil
}

// GetWorkIdList get the list of work id i.e. the device id
func GetWorkIdList(conn net.Conn) ([]uint32, error) {
	req := bb.UsbPack{
		MsgId: 0,
		Sta:   0,
		ReqId: bb.BB_RPC_GET_LIST,
	}
	resp, err := RequestWithPack(conn, &req, nil)
	if err != nil {
		return nil, err
	}
	log.Sugar().Debugw("response", "resp", resp)
	rBuf := bytes.NewBuffer(resp.Buf)

	workIdList := make([]uint32, 0)
	for rBuf.Len() > 0 {
		var workId uint32
		// TODO: make sure the endian is correct
		// the original code use `memcpy`
		err = binary.Read(rBuf, binary.NativeEndian, &workId)
		if err != nil {
			log.Sugar().Warnw("failed to read work id", "error", err, "raw", rBuf.Bytes())
			break
		}
		workIdList = append(workIdList, workId)
	}
	// NOTE: the status might be the number of work id, but not sure
	if resp.Sta < 0 {
		return workIdList, errorx.ExternalError.New("bad status %d", req.Sta)
	}

	return workIdList, nil
}

// SelectWorkId select and test the work id, i.e. the device id to see if it is a valid one
func SelectWorkId(conn net.Conn, workId uint32) (bool, error) {
	req := bb.UsbPack{
		MsgId: workId,
		Sta:   0,
		ReqId: bb.BB_RPC_SEL_ID,
	}
	resp, err := RequestWithPack(conn, &req, nil)
	if err != nil {
		return false, err
	}
	if resp.ReqId != bb.BB_RPC_SEL_ID {
		return false, errorx.ExternalError.New("unexpected request id %d", resp.ReqId)
	}
	if resp.Sta == -1 {
		return false, nil
	} else if resp.Sta < 0 {
		return false, errorx.ExternalError.New("bad status %d", req.Sta)
	}
	return true, nil
}
