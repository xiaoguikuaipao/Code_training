package znet

import (
	"fmt"
	"net"
	"zinx/internal/config"
	"zinx/ziface"

	"go.uber.org/zap"
)

// Server iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	Name    string
	Version string
	IP      string
	Port    int
	Router  ziface.IRouter
}

var DefaultServer = &Server{}

func init() {
	DefaultServer = NewServer(config.DefaultConfig).(*Server)
}

func (s *Server) Start() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer sugar.Sync() //nolint
	sugar.Infof("[Start] Server listening at IP:%s, Port:%d, is starting", s.IP, s.Port)
	addr, err := net.ResolveTCPAddr(s.Version, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		sugar.Errorf("resolve tcp addr error:%s", err)
		return
	}

	listener, err := net.ListenTCP(s.Version, addr)
	if err != nil {
		sugar.Errorf("listen tcp error:%s", err)
		return
	}
	sugar.Infof("[Start] Server %s success...", s.Name)

	var cid uint32
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			sugar.Error("accept error:%s", err)
			continue
		}

		dealConn := newConnection(conn, cid, s.Router)
		cid++
		go dealConn.Start()
	}
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("AddRouter success")
}

func NewServer(config *config.ServerConfig) ziface.IServer {
	return &Server{
		Name:    config.Name,
		Version: "tcp4",
		IP:      config.Host,
		Port:    config.Port,
		Router:  nil,
	}
}
