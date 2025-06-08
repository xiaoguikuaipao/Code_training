package ziface

type IConnCtrl interface {
	Get(ID uint32) (IConnection, error)
	Add(IConnection) error
	Del(ID uint32)
	ClearAll()
	Len() int
}
