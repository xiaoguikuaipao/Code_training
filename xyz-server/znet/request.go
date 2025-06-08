package znet

import (
	datastructimpl "data_struct_impl"
	"fmt"
	"strconv"
	"zinx/ziface"

	"github.com/google/uuid"
)

type Request struct {
	id uuid.UUID
	datastructimpl.BaseTask
	conn    ziface.IConnection
	msg     ziface.IMessage
	handler ziface.IHandler
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.MsgID()
}

func (r *Request) GetConn() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.Data()
}

func (r *Request) ID() uint32 {
	return r.id.ID()
}

func (r *Request) Attribute4Hash() string {
	return strconv.Itoa(int(r.conn.GetConnID()))
}

// For task concurrent model
func (r *Request) Deal() error {
	if r.handler == nil {
		return fmt.Errorf("no API for msgID[%d]", r.GetMsgID())
	}
	err := r.handler.PreHandler(r)
	if err != nil {
		return fmt.Errorf("PreHandler error:%w", err)
	}
	err = r.handler.Handler(r)
	if err != nil {
		return fmt.Errorf("Handler error:%w", err)
	}
	err = r.handler.PostHandler(r)
	if err != nil {
		return fmt.Errorf("PostHandler error:%w", err)
	}
	return nil
}
