package ziface

import (
	datastructimpl "data_struct_impl"
)

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve() error
	//路由功能
	AddHandler(msgID uint32, handler IHandler) error
	GetHandler(msgID uint32) (IHandler, error)

	GetTaskPool() datastructimpl.ITaskPool[datastructimpl.Itask]
	GetConnCtrl() IConnCtrl

	SetAfterConnCreate(func(IConnection))
	CallAfterConnCreate(IConnection)
	SetBeforeConnDestroy(func(IConnection))
	CallBeforeConnDestroy(IConnection)
}
