package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandler(request ziface.IRequest) {
	_, err := request.GetConn().GetTCPConnection().Write([]byte("Pre Ping..."))
	if err != nil {
		fmt.Println("preping error")
	}
}
func (pr *PingRouter) Handler(request ziface.IRequest) {
	_, err := request.GetConn().GetTCPConnection().Write([]byte("PPing..."))
	if err != nil {
		fmt.Println("preping error")
	}
}
func (pr *PingRouter) PostHandler(_ ziface.IRequest) {

}

func main() {
	s := znet.DefaultServer
	s.AddRouter(&PingRouter{})
	s.Serve()
}
