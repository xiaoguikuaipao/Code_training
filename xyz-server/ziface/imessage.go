package ziface

type IMessage interface {
	MsgID() uint32
	DataLen() uint32
	Data() []byte

	SetMsgID(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
