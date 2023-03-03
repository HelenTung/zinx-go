package main

import "github.com/helenvivi/zinx/zinnet"

// server 端
func main() {
	//实例化server
	s := zinnet.NewServer()
	//启动server
	s.Serve()
}
