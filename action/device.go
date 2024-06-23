package action

import "net"

type Device struct {
	conn  *net.TCPConn
	selId uint32
}
