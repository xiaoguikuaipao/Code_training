package znet

import (
	"errors"
	"fmt"
	"io"
	"zinx/ziface"
)

type Message struct {
	dataLen uint32
	id      uint32
	data    []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		dataLen: uint32(len(data) + 8),
		id:      id,
		data:    data,
	}
}

func ReadFrom(c ziface.IConnection, dp ziface.IDataPack) (msg *Message, err error) {
	// 一次Read读取的动作对应一次从内核缓冲区到用户区的复制操作，读取长度最大取决于内核设置(Recv-Q队列) - net.ipv4.tcp_rmem
	buf := make([]byte, dp.GetHeadLen())

	_, err = io.ReadFull(c.GetTCPConnection(), buf)
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		fmt.Printf("recv head error:%s\n", err)
		return nil, err
	}
	imsg, err := dp.UnPack(buf)
	if err != nil {
		fmt.Printf("unpack error:%s", err.Error())
		return nil, err
	}
	var ok bool
	msg, ok = imsg.(*Message)
	if !ok {
		return nil, errors.New("Assert Message error")
	}
	if msg.DataLen() > dp.GetHeadLen() {
		buf = make([]byte, msg.DataLen()-dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), buf)
		if err != nil {
			if err == io.EOF {
				return nil, err
			}
			fmt.Printf("recv body error:%s\n", err)
			return nil, err
		}
		msg.SetData(buf)
	}
	return msg, nil
}

func (m *Message) MsgID() uint32 {
	return m.id
}

func (m *Message) DataLen() uint32 {
	return m.dataLen
}

func (m *Message) Data() []byte {
	return m.data
}

func (m *Message) SetMsgID(id uint32) {
	m.id = id
}

func (m *Message) SetData(data []byte) {
	m.data = data
}

func (m *Message) SetDataLen(dl uint32) {
	m.dataLen = dl
}
