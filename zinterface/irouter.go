package zinterface

// 路由抽象接口、数据类型为irequest、对请求进行处理
type IRouter interface {
	//处理业务之前的hook（钩子）方法
	PreHandle(r IRequest)
	//处理conn业务的主方法
	Handle(r IRequest)
	//处理业务之后的hook（钩子）方法
	PostHandle(r IRequest)
}
