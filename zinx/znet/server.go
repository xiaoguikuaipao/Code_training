package znet

import (
	"fmt"
	"net"
	"zinx/ziface"

	"go.uber.org/zap"
)

// Server iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	Name    string
	Version string
	IP      string
	Port    int
}

func callBack(conn *net.TCPConn, data []byte, cnt int) error {
	_, err := conn.Write(data[:cnt])
	if err != nil {
		err = fmt.Errorf("callBack error:%w", err)
		return err
	}
	return nil
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

	var cid uint32 = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			sugar.Error("accept error:%s", err)
			continue
		}

		dealConn := NewConnection(conn, cid, callBack)
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

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:    name,
		Version: "tcp4",
		IP:      "0.0.0.0",
		Port:    8080,
	}
}
