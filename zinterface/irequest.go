package zinterface

type IRequest interface {
	//获取conn
	GetRequestConn() Iconn
	//获取数据
	GetRequestByte() []byte
}
