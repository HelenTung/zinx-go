package zinnet

import (
	"fmt"
	"net"

	"github.com/helenvivi/zinx/zinterface"
)

// 定义链接的结构
type Connection struct {
	//socket 套接字
	conn *net.TCPConn
	//conn 的 ID
	ConnID uint32
	//链接状态
	IsCloesd bool
	//当前链接绑定业务处理方法
	handleAPI zinterface.HandleFunc
	//告知当前链接状态的channel
	ExitChan chan bool
}

func (c *Connection) StartRead() {
	fmt.Println("StartRead goroutine is running")
	defer fmt.Println("connid ", c.ConnID, "reader goroutine is running!", c.RemoteAddr().String())
	defer c.Stop()
	for {
		//读取
		buf := make([]byte, 512)
		//堵塞
		cnt, err := c.conn.Read(buf)
		//读取异常、跳出本次循环
		if err != nil {
			fmt.Println("conn id ", c.ConnID, "reading err : \n", err)
			continue
		}
		//conn绑定的业务
		if err := c.handleAPI(c.conn, buf, cnt); err != nil {
			fmt.Println("conn id : ", c.ConnID, "handle is error : \n", err)
			return
		}

	}
}

// start conn、启动链接、让连接开始工作
func (c *Connection) Start() {
	fmt.Println("conn start()...connid = ", c.ConnID)

	//启动当前链接的读写
	go c.StartRead()
}

// stop conn、停止链接、结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("conn stop()...connid = ", c.ConnID)
	if c.IsCloesd {
		return
	}
	c.IsCloesd = true
	//关闭socket链接
	c.conn.Close()
	//关闭channel
	close(c.ExitChan)
}

// 获取当前链接的绑定socket 套接字
func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.conn
}

// 获取当前conn 的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的tcp状态、IP、Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	fmt.Println("conn send()...connid = ", c.ConnID)
	_, err := c.conn.Write(data)
	if err != nil {
		fmt.Println("conn id ", c.ConnID, "conn write err ", err)
		return err
	}
	return nil
}

// 实例化对象conn、初始化模块的方法,向外暴露接口
func NewConn(conn *net.TCPConn, connId uint32, api zinterface.HandleFunc) zinterface.Iconn {
	c := &Connection{
		conn:      conn,
		ConnID:    connId,
		handleAPI: api,
		IsCloesd:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
