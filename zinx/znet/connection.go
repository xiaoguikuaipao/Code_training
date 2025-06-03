package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClosed  bool
	handleAPI ziface.HandleFunc
	ExitChan  chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, callback_api ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connId,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChan:  make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	fmt.Printf("conn start(), id=%d\n", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	fmt.Printf("conn start(), id=%d\n", c.ConnID)

	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send([]byte, int) error {
	return nil
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) StartReader() {
	defer c.Stop()
	defer fmt.Printf("connID = %d, reader exits, remote addr is %s\n", c.ConnID, c.RemoteAddr().String())

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("recv buf error:%s\n", err)
			continue
		}
		err = c.handleAPI(c.Conn, buf, cnt)
		if err != nil {
			fmt.Printf("handle error:%s\n", err.Error())
			break
		}
	}
}
