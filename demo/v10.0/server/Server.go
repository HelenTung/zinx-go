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

// 处理conn业务的主方法
func (r *PingRouter) Handle(req zinterface.IRequest) {
	fmt.Println("call PingRouter handle...")
	//读取客户端数据，再回写ping...ping...ping...
	fmt.Println("recv from client: msgID= ", req.GetMsgID(),
		"data= ", string(req.GetRequestData()))
	err := req.GetRequestConn().Send(200, []byte("the is V10.0,data = ping...ping...ping..."))
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
	err := req.GetRequestConn().Send(201, []byte("the is V10.0,data = hello "))
	if err != nil {
		fmt.Println(err)
	}
}

// 链接开始的业务
func DoConnectionBegin(c zinterface.Iconn) {
	fmt.Println("\n=========>DoConnection Begin Call")
	if err := c.Send(400, []byte("DoConnection Begin")); err != nil {
		fmt.Println(err)
	}
}

// 链接断开的业务
func DoConnectionEnd(c zinterface.Iconn) {
	fmt.Println("\n=========>DoConnection End Call")
}

// server 端
func main() {
	//实例化server
	s := zinnet.NewServer()
	//开始业务
	s.GetHook().SetOnConnStart(DoConnectionBegin)
	//断开业务
	s.GetHook().SetOnConnStop(DoConnectionEnd)

	//当用户发送0时、调度方法PingRouter对应的方法
	s.AddRouter(0, &PingRouter{})
	//当用户发送1时、调度方法PingRouter对应的方法
	s.AddRouter(1, &HelloRouter{})

	//启动server
	s.Serve()
}
