package main

import (
	"ar8030/action"
	"ar8030/log"
	"fmt"
	"net"
	"os"
)

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
	defer func(conn net.Conn) {
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
	// 0x3fff is a magic value, no idea why it's chosen
	st, err := action.GetStatus(conn, selId, 0x3fff)
	if err != nil {
		log.Sugar().Panicw("failed to get status", "error", err.Error())
	}
	log.Sugar().Infow("status", "status", st)
	sysInfo, err := action.GetSysInfo(conn, selId)
	if err != nil {
		log.Sugar().Panicw("failed to get system info", "error", err.Error())
	}
	log.Sugar().Infow("system info", "info", sysInfo)

	oCfg, err := action.GetTotalCfg(conn, selId)
	if err != nil {
		log.Sugar().Panicw("failed to get configuration", "error", err.Error())
	}
	sCfg := string(oCfg)
	log.Sugar().Infow("configuration", "len", len(sCfg))
	f, err := os.Create("config.json")
	if err != nil {
		log.Sugar().Panicw("failed to create file", "error", err.Error())
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Sugar().Panicw("failed to close file", "error", err.Error())
		}
	}(f)
	_, err = f.WriteString(sCfg)
	if err != nil {
		log.Sugar().Panicw("failed to write to file", "error", err.Error())
	}
}
