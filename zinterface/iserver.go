package zinterface

// 定义服务器的接口、抽象方法
type IServer interface {
	//启动服务器
	Start()
	//运行服务器
	Serve()
	//停止服务器
	Stop()
}
