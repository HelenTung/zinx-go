package zinnet

import (
	"fmt"
	"net"

	"github.com/helenvivi/zinx/utils"
	"github.com/helenvivi/zinx/zinterface"
)

// 服务器实例化、定义server模块、接口实现
type Server struct {
	//服务器名称
	Name string
	//所绑定的IP 版本
	IPVersion string
	//绑定的IP地址
	IP string
	//port 监听的端口
	Port int
	// //路由功能、server注册的conn
	// router     zinterface.IRouter
	//server 的消息管理模块
	MsgHandler zinterface.IMsgHandler
}

// 定义计数器
var cnt uint32

// // conn 需要的操作函数、定义为当前客户端发起conn的绑定函数
// func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
// 	//回显业务
// 	fmt.Println("[Conn Handle] CallBackClient...")
// 	_, err := conn.Write(data)
// 	if err != nil {
// 		fmt.Println("write back data err", err)
// 		return errors.New("CallBackClient error")
// 	}
// 	return nil
// }

// 初始化Server对象
func NewServer() zinterface.IServer {
	s := &Server{
		Name:       utils.Globa.Name,
		IPVersion:  "tcp4",
		IP:         utils.Globa.Host,
		Port:       utils.Globa.TcpPort,
		MsgHandler: NewMessageHandler(),
	}
	return s
}

// 启动
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s,listenner at ip : %s,Port : %d is starting\n",
		utils.Globa.Name, utils.Globa.Host, utils.Globa.TcpPort)
	fmt.Printf("[Zinx] Version : %s, MaxConn : %d, MaxPackSize : %d \n",
		utils.Globa.Version, utils.Globa.MaxConn, utils.Globa.MaxPackageSize)
	//放到携程处理、避免因为读取数据堵塞
	go func() {
		//获取tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tap addr error : ", err)
			return
		}
		//监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("ListenTCP is error,listen", s.IPVersion, "err : ", err)
			return
		}
		fmt.Println("start zinx server success,", s.Name, "success,listenning...")

		cnt = 0
		//阻塞等待client连接、处理客户端业务（读写）
		for {
			//堵塞，等待客户端链接、如果客户端有链接、则堵塞返回
			TCPconn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("listenner.AcceptTCP error : ", err)
				continue
			}

			// // 与client已经建立连接、进行业务处理、处理512byte的回显业务

			// go func() {
			// 	for {
			// 		//read
			// 		buff := make([]byte, 512)
			// 		//堵塞、等待数据到达
			// 		cnt, err := TCPconn.Read(buff)
			// 		if err != nil {
			// 			fmt.Println("TCPconn.Read error :", err)
			// 			continue
			// 		}
			// 		//write、回显业务
			// 		if _, err := TCPconn.Write(buff[:cnt]); err != nil {
			// 			fmt.Println("write back buff err : ", err)
			// 			continue
			// 		}

			// 	}
			// }()
			//将处理新链接的方法与conn进行绑定、得到新的conn模块
			conn := NewConn(TCPconn, cnt, s.MsgHandler)
			//计数器+1
			cnt++
			go conn.Start()
		}

	}()
}

// 停止
func (s *Server) Stop() {
	//TODO:将服务器的资源、状态、与链接信息进行停止与回收释放
}

// 运行
func (s *Server) Serve() {
	//start server 服务功能
	s.Start()
	//TODO:启动服务之后的额外业务

	//堵塞
	select {}

}

// server 注册 路由
func (s *Server) AddRouter(msgID uint32, router zinterface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add router to conn success!")
}
