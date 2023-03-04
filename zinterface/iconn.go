package zinterface

import (
	"net"
)

type Iconn interface {
	//启动链接、让连接开始工作
	Start()
	//停止链接、结束当前连接的工作
	Stop()
	//获取当前链接的绑定socket 套接字
	GetTcpConn() *net.TCPConn
	//获取当前conn 的ID
	GetConnID() uint32
	//获取远程客户端的tcp状态、IP、Port
	RemoteAddr() net.Addr
	//发送数据
	Send(msgId uint32, data []byte) error
}

//// 处理链接业务的方法
// type HandleFunc func(*net.TCPConn, []byte, int) error
