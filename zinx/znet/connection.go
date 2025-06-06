package znet

import (
	"fmt"
	"io"
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

func newConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
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
	defer fmt.Printf("connID = %d, reader exit, remote addr is %s\n", c.ConnID, c.RemoteAddr().String())

	for {
		// 一次Read读取的动作对应一次从内核缓冲区到用户区的复制操作，读取长度最大取决于内核设置(Recv-Q队列) - net.ipv4.tcp_rmem
		buf := make([]byte, 4096)

		n, err := c.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("recv buf error:%s\n", err)
			continue
		}
		req := Request{
			conn: c,
			data: buf[:n],
		}
		go func() {
			// 模板设计模式
			c.Router.PreHandler(&req)
			c.Router.Handler(&req)
			c.Router.PostHandler(&req)
		}()
	}
}
