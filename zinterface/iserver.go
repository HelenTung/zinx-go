package zinterface

// 定义服务器的接口、抽象方法
type IServer interface {
	//启动服务器
	Start()
	//运行服务器
	Serve()
	//停止服务器
	Stop()
	// 路由功能：给当前服务注册路由方法、提供给客户端的链接使用
	AddRouter(msgID uint32, router IRouter)
	//获取mgr
	GetConnMgr() IConnManger
	//获取hook
	GetHook() Ihook
}
