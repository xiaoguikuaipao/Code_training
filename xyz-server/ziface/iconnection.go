package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	Send(msgID uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
