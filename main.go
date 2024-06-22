package main

import (
	"ar8030/bb"
	"ar8030/log"
	"bytes"
	"fmt"
	"net"
)

const (
	HOST = "127.0.0.1"
	PORT = 50000
)

func TestServer(conn *net.TCPConn) error {
	// length, capacity
	buf_ := make([]byte, 0, 1024)
	// It can also be used to set the initial size of the internal buffer for writing.
	// To do that, buf should have the desired capacity but a length of zero.
	buf := bytes.NewBuffer(buf_)
	pack := bb.UsbPack{
		MsgId: 0,
		Sta:   0,
		ReqId: bb.BB_RPC_TEST,
	}
	err := pack.Write(buf)
	if err != nil {
		return err
	}
	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	buf_ = make([]byte, 1024)
	sz, err := conn.Read(buf_)
	log.Sugar().Debugw("received", "size", sz)
	buf = bytes.NewBuffer(buf_[:sz])
	err = pack.Read(buf)
	if err != nil {
		return err
	}
	log.Sugar().Infow("received", "pack", pack)
	return nil
}

func main() {
	var err error
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", HOST, PORT))
	if err != nil {
		log.Sugar().Panicw("failed to resolve TCP address", "error", err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	defer func(conn *net.TCPConn) {
		err = conn.Close()
		if err != nil {
			log.Sugar().Panicw("failed to close connection", "error", err)
		}
	}(conn)
	if err != nil {
		log.Sugar().Panicw("failed to dial TCP", "error", err)
	}
	err = TestServer(conn)
	if err != nil {
		log.Sugar().Panicw("failed to test server", "error", err)
	}
}
