package znet

import (
	datastructimpl "data_struct_impl"
	"fmt"
	"net"
	"zinx/internal/config"
	"zinx/ziface"

	"go.uber.org/zap"
)

// Server iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	Name             string
	Version          string
	IP               string
	Port             int
	Handlers         ziface.IHandlerGroup
	ConnCtr          ziface.IConnCtrl
	QueueGroup       datastructimpl.TaskPool[datastructimpl.Itask]
	beforeConnCreate func(ziface.IConnection)
	afterConnCreate  func(ziface.IConnection)
}

var DefaultServer = &Server{}

func init() {
	DefaultServer = NewServer(config.DefaultConfig).(*Server)
}

func NewServer(config *config.ServerConfig) ziface.IServer {
	return &Server{
		Name:             config.Name,
		Version:          "tcp4",
		IP:               config.Host,
		Port:             config.Port,
		Handlers:         NewHandlerGroup(),
		ConnCtr:          NewConnCtrl(),
		afterConnCreate:  func(ziface.IConnection) {},
		beforeConnCreate: func(ziface.IConnection) {},
	}
}

func (s *Server) Serve() error {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	defer sugar.Sync() //nolint
	sugar.Infof("[Start] Server listening at IP:%s, Port:%d, is starting", s.IP, s.Port)
	addr, err := net.ResolveTCPAddr(s.Version, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		sugar.Errorf("resolve tcp addr error:%s", err)
		return err
	}

	listener, err := net.ListenTCP(s.Version, addr)
	if err != nil {
		sugar.Errorf("listen tcp error:%s", err)
		return err
	}
	sugar.Infof("[Start] Server %s success...", s.Name)

	s.QueueGroup = datastructimpl.NewQueueGroup[datastructimpl.Itask](
		int(config.DefaultConfig.MaxQueuePoolSize),
		int(config.DefaultConfig.MaxQueueSize),
		int(config.DefaultConfig.MaxQueueWorker),
	)
	go func() {
		var cid uint32
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				sugar.Error("accept error:%s", err)
				continue
			}
			cid++
			dealConn, err := newConnection(conn, cid, s)
			if err != nil {
				fmt.Printf("newConnection[%d] error:%v\n", cid, err)
				continue
			}
			s.CallAfterConnCreate(dealConn)
			go dealConn.Start()
		}
	}()
	return nil
}

func (s *Server) Stop() {
	fmt.Println("server stop, clear all conn")
	s.ConnCtr.ClearAll()
}

func (s *Server) Start() {
	if err := s.Serve(); err != nil {
		panic(err)
	}
	select {}
}

func (s *Server) AddHandler(msgID uint32, router ziface.IHandler) error {
	return s.Handlers.AddHandler(msgID, router)
}

func (s *Server) GetHandler(msgID uint32) (ziface.IHandler, error) {
	return s.Handlers.GetHandler(msgID)
}

func (s *Server) GetTaskPool() datastructimpl.ITaskPool[datastructimpl.Itask] {
	return s.QueueGroup
}

func (s *Server) GetConnCtrl() ziface.IConnCtrl {
	return s.ConnCtr
}

func (s *Server) SetAfterConnCreate(f func(ziface.IConnection)) {
	s.afterConnCreate = f
}

func (s *Server) CallAfterConnCreate(c ziface.IConnection) {
	s.afterConnCreate(c)
}

func (s *Server) SetBeforeConnDestroy(f func(ziface.IConnection)) {
	s.beforeConnCreate = f
}

func (s *Server) CallBeforeConnDestroy(c ziface.IConnection) {
	s.beforeConnCreate(c)
}
