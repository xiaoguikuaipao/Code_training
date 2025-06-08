package ziface

type IRequest interface {
	GetConn() IConnection
	GetData() []byte
	GetMsgID() uint32
}
