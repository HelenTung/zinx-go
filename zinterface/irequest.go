package zinterface

type IRequest interface {
	//获取conn
	GetRequestConn() Iconn
	//获取数据
	GetRequestData() []byte
	//获取msg ID
	GetMsgID() uint32
}
