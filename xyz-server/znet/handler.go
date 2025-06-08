package znet

import (
	"fmt"
	"sync"
	"zinx/ziface"
)

type HandlerGroup struct {
	Handlers map[uint32]ziface.IHandler
	mu       sync.RWMutex
}

func NewHandlerGroup() *HandlerGroup {
	return &HandlerGroup{
		Handlers: make(map[uint32]ziface.IHandler),
	}
}

func (h *HandlerGroup) GetHandler(msgID uint32) (ziface.IHandler, error) {
	if r, ok := h.Handlers[msgID]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("handler[%d] not exist", msgID)
}

func (h *HandlerGroup) AddHandler(msgID uint32, router ziface.IHandler) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.Handlers[msgID]; ok {
		return fmt.Errorf("handler[%d] already exist", msgID)
	}
	h.Handlers[msgID] = router
	return nil
}

type BaseHandler struct {
}

var _ ziface.IHandler = (*BaseHandler)(nil)

func (br *BaseHandler) PreHandler(_ ziface.IRequest) error {
	return nil
}
func (br *BaseHandler) Handler(_ ziface.IRequest) error {
	return nil
}

func (br *BaseHandler) PostHandler(_ ziface.IRequest) error {
	return nil
}
