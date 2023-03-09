package main

import (
	"fmt"

	"github.com/helenvivi/zinx/zinnet"
	"github.com/helenvivi/zinx/zinterface"
)

// 编写自定义路由、重写方法
type PingRouter struct {
	zinnet.BaseRouter
}

// 编写自定义路由、重写方法
type HelloRouter struct {
	zinnet.BaseRouter
}

// // 处理业务之前的hook（钩子）方法
// func (r *PingRouter) PreHandle(req zinterface.IRequest) {
// 	fmt.Println("call router prehandle...")
// 	_, err := req.GetRequestConn().GetTcpConn().Write([]byte("before ping...\n"))
// 	if err != nil {
// 		fmt.Println("call back before ping error\n", err)
// 	}
// }

// 处理conn业务的主方法
func (r *PingRouter) Handle(req zinterface.IRequest) {
	fmt.Println("call PingRouter handle...")
	//读取客户端数据，再回写ping...ping...ping...
	fmt.Println("recv from client: msgID= ", req.GetMsgID(),
		"data= ", string(req.GetRequestData()))
	err := req.GetRequestConn().Send(200, []byte("the is V8.0,data = ping...ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

// 处理conn业务的主方法
func (r *HelloRouter) Handle(req zinterface.IRequest) {
	fmt.Println("call HelloRouter handle...")
	//读取客户端数据，再回写ping...ping...ping...
	fmt.Println("recv from client: msgID= ", req.GetMsgID(),
		"data= ", string(req.GetRequestData()))
	err := req.GetRequestConn().Send(201, []byte("the is V8.0,data = hello "))
	if err != nil {
		fmt.Println(err)
	}
}

// // 处理业务之后的hook（钩子）方法
// func (r *PingRouter) PostHandle(req zinterface.IRequest) {
// 	fmt.Println("call router posthandle...")
// 	_, err := req.GetRequestConn().GetTcpConn().Write([]byte("after ping...\n"))
// 	if err != nil {
// 		fmt.Println("call back after ping error\n", err)
// 	}
// }

// server 端
func main() {
	//实例化server
	s := zinnet.NewServer()
	//当用户发送0时、调度方法PingRouter对应的方法
	s.AddRouter(0, &PingRouter{})
	//当用户发送1时、调度方法PingRouter对应的方法
	s.AddRouter(1, &HelloRouter{})
	//启动server
	s.Serve()
}
