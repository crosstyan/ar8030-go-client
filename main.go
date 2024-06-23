package main

import (
	"ar8030/action"
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"unsafe"
)

var data = []byte{
	0x01, 0x00, 0x00, 0x01, 0x01, 0x01, 0xa3, 0x35, 0xf7, 0x52, 0x00, 0x00, // h
	0x00, 0x01, 0x01, 0x01, 0x04, 0x03,
	0x2c, 0x20, // invalid gibberish only for padding
	0x20, 0x0b, 0x20, 0x00, // band in Hz (00200b20=2'100'000)
	0x19, 0x02, 0x01, 0x01, 0x04, 0x03,
	0x09, 0x09,
	0x20, 0x0b, 0x20, 0x00, // the same
	0x00, 0x01, 0x01, 0x01, 0x04, 0x03, 0x20, 0x22, 0x20, 0x0b, 0x20, 0x00,
	0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x20, 0x22, 0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x01,
	0x04, 0x03, 0x0a, 0x09, 0x20, 0x0b, 0x20, 0x00, 0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x09, 0x09,
	0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x01, 0x04, 0x03, 0x30, 0x22, 0x20, 0x0b, 0x20, 0x00,
	0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x64, 0x62, 0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x01,
	0x04, 0x03, 0x22, 0x2c, 0x20, 0x0b, 0x20, 0x00, 0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x22, 0x2c,
	0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x01, 0x04, 0x03, 0x09, 0x09, 0x20, 0x0b, 0x20, 0x00,
	0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x5b, 0x20, 0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x01,
	0x04, 0x03, 0x2c, 0x20, 0x20, 0x0b, 0x20, 0x00, 0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x5d, 0x2c,
	0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x01, 0x04, 0x03, 0x62, 0x6d, 0x20, 0x0b, 0x20, 0x00,
	0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x20, 0x22, 0x20, 0x0b, 0x20, 0x00, 0x00, 0x01, 0x01, 0x00,
	0x02, 0x01, 0x20, 0x22, 0x20, 0x0b, 0x20, 0x00, 0x19, 0x02, 0x01, 0x00, 0x02, 0x01, 0x09, 0x22,
	0xe8, 0x1f, 0x25, 0x00, 0x00, 0x01, 0x01, 0x01, 0x04, 0x03, 0x22, 0x31, 0x20, 0x0b, 0x20, 0x00,
	0x19, 0x02, 0x01, 0x01, 0x04, 0x03, 0x22, 0x31, 0x20, 0x0b, 0x20, 0x00, 0x00, 0x19, 0xa3, 0x35,
	0x87, 0x91, 0x0a, 0x09, 0x09, 0x09, 0x09, 0x09, 0x22, 0x31, 0x33, 0x64, 0x62, 0x6d, 0x22, 0x3a,
	0x20, 0x5b, 0x20, 0x22, 0x31, 0x45, 0x22, 0x2c, 0x20, 0x22, 0x34, 0x30, 0x22, 0x2c, 0x20, 0x22,
	0x31, 0x41, 0x22, 0x2c, 0x20, 0x22, 0x32, 0x43, 0x22, 0x5d, 0x2c, 0x0a,
}

const (
	HOST = "127.0.0.1"
	PORT = 50000
)

func main() {
	var err error
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", HOST, PORT))
	if err != nil {
		log.Sugar().Panicw("failed to resolve TCP address", "error", err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	defer func(conn *net.TCPConn) {
		err = conn.Close()
		if err != nil {
			log.Sugar().Panicw("failed to close connection", "error", err.Error())
		}
	}(conn)
	if err != nil {
		log.Sugar().Panicw("failed to dial TCP", "error", err.Error())
	}
	err = action.TestServer(conn)
	if err != nil {
		log.Sugar().Panicw("failed to test server", "error", err.Error())
	}
	wrkList, err := action.GetWorkIdList(conn)
	if err != nil {
		log.Sugar().Panicw("failed to get work id list", "error", err.Error())
	}
	log.Sugar().Infow("work id list", "list", wrkList)
	var selId uint32 = 0
	if len(wrkList) != 0 {
		selId = wrkList[0]
		ok, err := action.SelectWorkId(conn, selId)
		if err != nil {
			log.Sugar().Panicw("failed to test work id", "error", err.Error(), "id", selId)
		}
		if !ok {
			log.Sugar().Errorw("work id is invalid", "id", selId)
			os.Exit(1)
		}
	}
	// 0x3fff is a magic value
	st, err := action.GetStatus(conn, selId, 0x3fff)
	if err != nil {
		log.Sugar().Panicw("failed to get status", "error", err.Error())
	}
	log.Sugar().Infow("status", "status", st)

	refSt := bb.GetStatusOut{}
	alignment := unsafe.Alignof(refSt)
	log.Sugar().Infow("alignment", "alignment", alignment)
	rbuf := bytes.NewBuffer(data)
	err = binary.Read(rbuf, binary.NativeEndian, &refSt)
	if err != nil {
		log.Sugar().Panicw("failed to read binary data", "error", err.Error())
	}
	log.Sugar().Infow("status", "status", refSt)
	log.Sugar().Info("ok")
}
