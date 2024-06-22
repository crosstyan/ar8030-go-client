package bb

import (
	"bytes"
	"encoding/binary"
	"github.com/joomcode/errorx"
)

type UsbPack struct {
	MsgId uint32
	Sta   int32
	ReqId RequestId
	Buf   []byte
}

const UsbPackHeaderSize = 19 // 1 + 4 + 4 + 4 + 4 + 1 + 1

// Write expect an empty buffer. If the given buffer is not empty,
// the calculated XOR check will be wrong.
//
// See make_usbpack2buff in usbpack.c
func (p *UsbPack) Write(buf *bytes.Buffer) error {
	if buf.Len() != 0 {
		return errorx.IllegalArgument.New("buffer is not empty, expecting an empty buffer, got %d", buf.Len())
	}
	packetSz := UsbPackHeaderSize + len(p.Buf)
	if buf.Available() < packetSz {
		return errorx.IllegalArgument.New("buffer is too small, expecting %d, got %d", UsbPackHeaderSize+len(p.Buf), buf.Len())
	}

	buf.WriteByte(0xaa)
	var err error
	if p.Buf == nil {
		// binary.BigEndian.PutUint32(buf.Next(4), 0)
		err = binary.Write(buf, binary.BigEndian, uint32(0))
	} else {
		err = binary.Write(buf, binary.BigEndian, uint32(len(p.Buf)))
	}
	if err != nil {
		return errorx.Decorate(err, "failed to write buffer size")
	}
	err = binary.Write(buf, binary.BigEndian, p.ReqId)
	if err != nil {
		return errorx.Decorate(err, "failed to write request id")
	}
	err = binary.Write(buf, binary.BigEndian, p.MsgId)
	if err != nil {
		return errorx.Decorate(err, "failed to write message id")
	}
	err = binary.Write(buf, binary.BigEndian, p.Sta)
	if err != nil {
		return errorx.Decorate(err, "failed to write status")
	}

	// a byte xor
	buf.WriteByte(xorCheck(buf.Bytes()))
	if p.Buf != nil {
		buf.Write(p.Buf)
	}
	buf.WriteByte(0xbb)
	return nil
}

func (p *UsbPack) Read(buf *bytes.Buffer) error {
	if buf.Len() < UsbPackHeaderSize {
		return errorx.IllegalArgument.New("buffer is too small, expecting %d, got %d", UsbPackHeaderSize, buf.Len())
	}
	aa, err := buf.ReadByte()
	if err != nil {
		return errorx.Decorate(err, "failed to read 0xaa header")
	}
	if aa != 0xaa {
		return errorx.IllegalArgument.New("expecting 0xaa header, got %x", aa)
	}
	var bufSz uint32 = 0
	err = binary.Read(buf, binary.BigEndian, &bufSz)
	if err != nil {
		return errorx.Decorate(err, "failed to read buffer size")
	}
	err = binary.Read(buf, binary.BigEndian, &p.ReqId)
	if err != nil {
		return errorx.Decorate(err, "failed to read request id")
	}
	err = binary.Read(buf, binary.BigEndian, &p.MsgId)
	if err != nil {
		return errorx.Decorate(err, "failed to read message id")
	}
	err = binary.Read(buf, binary.BigEndian, &p.Sta)
	if err != nil {
		return errorx.Decorate(err, "failed to read status")
	}
	_, err = buf.ReadByte()
	if err != nil {
		return errorx.Decorate(err, "failed to read xor")
	}
	// we don't do xor check... TCP is reliable
	p.Buf = make([]byte, bufSz)
	_, err = buf.Read(p.Buf)
	if err != nil {
		return errorx.Decorate(err, "failed to read buffer")
	}
	bb, err := buf.ReadByte()
	if err != nil {
		return errorx.Decorate(err, "failed to read 0xbb footer")
	}
	if bb != 0xbb {
		return errorx.IllegalArgument.New("expecting 0xbb footer, got %x", bb)
	}
	return nil
}
