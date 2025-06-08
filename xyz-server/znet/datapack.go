package znet

import (
	"bytes"
	"encoding/binary"
	"zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, msg.DataLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, msg.MsgID())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, msg.Data())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

func (d *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	buf := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(buf, binary.LittleEndian, &msg.dataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.id); err != nil {
		return nil, err
	}
	return msg, nil
}
