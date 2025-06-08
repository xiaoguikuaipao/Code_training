package main

import (
	"fmt"
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
	s.AddHandler(0, &PingRouter{})
	s.Serve()
}
