package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	ExitChan chan bool
	Router   ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
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

func (c *Connection) Send([]byte) error {
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
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf("recv buf error:%s\n", err)
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		go func() {
			c.Router.PreHandler(&req)
		}()
	}
}
