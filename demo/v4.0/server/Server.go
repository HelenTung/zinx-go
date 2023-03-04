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

// 处理业务之前的hook（钩子）方法
func (r *PingRouter) PreHandle(req zinterface.IRequest) {
	fmt.Println("call router prehandle...")
	_, err := req.GetRequestConn().GetTcpConn().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error\n", err)
	}
}

// 处理conn业务的主方法
func (r *PingRouter) Handle(req zinterface.IRequest) {
	fmt.Println("call router handle...")
	_, err := req.GetRequestConn().GetTcpConn().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("call back ping error\n", err)
	}
}

// 处理业务之后的hook（钩子）方法
func (r *PingRouter) PostHandle(req zinterface.IRequest) {
	fmt.Println("call router posthandle...")
	_, err := req.GetRequestConn().GetTcpConn().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error\n", err)
	}
}

// server 端
func main() {
	//实例化server
	s := zinnet.NewServer()
	s.AddRouter(&PingRouter{})
	//启动server
	s.Serve()
}
