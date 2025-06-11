package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/internal/config"
	"zinx/ziface"
)

var (
	ErrConnectLimit = errors.New("connect exceeds limit")
)

type ConnCtrl struct {
	connections map[uint32]ziface.IConnection
	mu          sync.RWMutex
}

func NewConnCtrl() *ConnCtrl {
	return &ConnCtrl{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnCtrl) Get(ID uint32) (ziface.IConnection, error) {
	c.mu.RLock()
	if conn, ok := c.connections[ID]; ok {
		return conn, nil
	}
	c.mu.RUnlock()
	return nil, fmt.Errorf("conn[%d] not exist", ID)
}

func (c *ConnCtrl) Add(conn ziface.IConnection) error {
	if c.Len() >= config.DefaultConfig.MaxConn {
		return ErrConnectLimit
	}

	c.mu.Lock()
	if conn, ok := c.connections[conn.GetConnID()]; ok {
		return fmt.Errorf("conn[%d] is already existed", conn.GetConnID())
	}
	c.connections[conn.GetConnID()] = conn
	c.mu.Unlock()
	return nil
}

func (c *ConnCtrl) Del(ID uint32) {
	c.mu.Lock()
	c.connections[ID].Stop()
	delete(c.connections, ID)
	c.mu.Unlock()
}

func (c *ConnCtrl) ClearAll() {
	c.mu.Lock()
	for ID := range c.connections {
		c.connections[ID].Stop()
		delete(c.connections, ID)
	}
	c.mu.Unlock()
}

func (c *ConnCtrl) Len() int {
	c.mu.RLock()
	l := len(c.connections)
	c.mu.RUnlock()
	return l
}
