package ziface

type IHandlerGroup interface {
	GetHandler(msgID uint32) (IHandler, error)
	AddHandler(msgID uint32, handler IHandler) error
}
type IHandler interface {
	PreHandler(request IRequest) error
	Handler(request IRequest) error
	PostHandler(request IRequest) error
}
