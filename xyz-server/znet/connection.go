package znet

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"
	"zinx/ziface"

	"github.com/google/uuid"
)

type Connection struct {
	conn     *net.TCPConn
	connID   uint32
	isClosed bool
	exitChan chan bool
	msgChan  chan []byte
	server   ziface.IServer
}

func newConnection(conn *net.TCPConn, connID uint32, server *Server) (*Connection, error) {
	c := &Connection{
		conn:     conn,
		connID:   connID,
		isClosed: false,
		exitChan: make(chan bool, 1),
		msgChan:  make(chan []byte),
		server:   server,
	}
	err := server.ConnCtr.Add(c)
	if err != nil {
		c.Stop()
		return nil, err
	}
	return c, nil
}

func (c *Connection) Start() {
	fmt.Printf("conn start(), id=%d\n", c.connID)
	go c.StartReader()
	go c.StartWriter()
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	fmt.Printf("conn stop(), id=%d\n", c.connID)

	c.isClosed = true
	c.conn.Close()
	close(c.exitChan)
	close(c.msgChan)
	c.server.GetConnCtrl().Del(c.connID)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnID() uint32 {
	return c.connID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) Send(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection is closed")
	}
	dp := NewDataPack()
	bs, err := dp.Pack(NewMessage(msgID, data))
	if err != nil {
		return err
	}
	c.msgChan <- bs
	return err
}

func (c *Connection) StartReader() {
	defer c.Stop()
	var err error
	defer fmt.Printf("connID = %d, reader exit, remote addr is %s, error:%v \n", c.connID, c.RemoteAddr().String(), err)

	dp := NewDataPack()
	for {
		var msg *Message
		msg, err = ReadFrom(c, dp)
		if err != nil {
			return
		}
		h, err := c.server.GetHandler(msg.MsgID())
		if err != nil {
			fmt.Println("no API for msgID:", msg.MsgID())
			return
		}
		id, err := uuid.NewV7()
		if err != nil {
			fmt.Println("gen uuid error:", err)
			return
		}
		req := Request{
			id:      id,
			conn:    c,
			msg:     msg,
			handler: h,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = c.server.GetTaskPool().AddRobin(ctx, &req)
		if err != nil {
			fmt.Printf("handle req[%s]\n", id)
			return
		}
	}
}

func (c *Connection) StartWriter() {
	for {
		select {
		case data := <-c.msgChan:
			_, err := c.GetTCPConnection().Write(data)
			if err != nil {
				fmt.Println("write error, data: ", data)
			}
		case <-c.exitChan:
			return
		}
	}
}
