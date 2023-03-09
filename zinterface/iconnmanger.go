package zinterface

// conn 管理抽象接口
type IConnManger interface {
	//删除链接
	DeleteConn(conn Iconn)
	//添加链接
	AddteConn(conn Iconn)
	//connid ——> 获取conn
	GetConn(uint32) (Iconn, error)
	//得到链接总数
	GetConnNum() int
	//清空并终止所有链接
	ClearConn()
}
