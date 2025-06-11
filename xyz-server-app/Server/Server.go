package main

import (
	"fmt"
	"time"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseHandler
}

func (pr *PingRouter) PreHandler(request ziface.IRequest) error {
	_, err := request.GetConn().GetTCPConnection().Write([]byte("Pre Ping..."))
	if err != nil {
		fmt.Println("preping error")
		return err
	}
	return nil
}
func (pr *PingRouter) Handler(request ziface.IRequest) error {
	_, err := request.GetConn().GetTCPConnection().Write([]byte("PPing..."))
	if err != nil {
		fmt.Println("preping error")
		return err
	}
	return nil
}
func (pr *PingRouter) PostHandler(_ ziface.IRequest) error {
	return nil
}

func main() {
	s := znet.DefaultServer
	_ = s.AddHandler(0, &PingRouter{})
	s.SetBeforeConnDestroy(func(ziface.IConnection) {
		fmt.Println("before destroy:", time.Now())
	})
	s.Start()
}
