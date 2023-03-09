package zinterface

// IMsgHandler 是消息处理模块的抽象层，定义了需要实现的方法。
type IMsgHandler interface {
	//调度/执行对应的router方法
	DoMsgHandler(r IRequest)
	//给路由添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)
	//开辟任务处理pool
	StartWorkPool()
	// 将消息发送到队列中
	SendMsgTaskQueue(req IRequest)
}
