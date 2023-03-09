package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/helenvivi/zinx/zinnet"
)

func main() {
	fmt.Println("Client1 starting...")
	time.Sleep(1 * time.Second)

	//建立链接、获取conn
	conn, err := net.Dial("tcp", "127.0.0.1:7070")
	if err != nil {
		fmt.Println("client start error ", err, " exit!")
		return
	}
	//循环
	for {
		//向conn write 数据
		// _, err := conn.Write([]byte("hello,the is zinx v4.0.."))
		// if err != nil {
		// 	fmt.Println("conn write error : ", err)
		// 	return
		// }

		// buf := make([]byte, 512)
		// cnt, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("conn read error : ", err)
		// 	return
		// }
		// fmt.Printf("server call back : %s ,cnt = %d\n", buf, cnt)

		//定义封包对象
		dp := zinnet.NewData()
		//封包
		msg, err := dp.PackMsg(zinnet.NewMsg(1, []byte("the is client1,zinx v7.0")))
		if err != nil {
			fmt.Println("pack error:", err)
			return
		}
		//发送
		if _, err := conn.Write(msg); err != nil {
			fmt.Println("conn write error")
			return
		}
		//此时受到服务端回复，msgid=1
		//读取二进制流
		msghead := make([]byte, dp.GetHeadlen())
		if _, err := io.ReadFull(conn, msghead); err != nil {
			fmt.Println("client conn server error", err)
			break
		}
		//解包
		ServerMsg, err := dp.UnPackMsg(msghead)
		if err != nil {
			fmt.Println("client dp Unpack error", err)
			break
		}
		//转换对象
		Msg := ServerMsg.(*zinnet.Message)
		//判断数据包长度
		if Msg.GetMessageLen() > 0 {
			Msg.Date = make([]byte, Msg.GetMessageLen())
			if _, err := io.ReadFull(conn, Msg.Date); err != nil {
				fmt.Println("client get conn data error ", err)
				return
			}
			fmt.Println("--->Recv MsgID:", Msg.ID, "DataLen=",
				Msg.Datelen, "Data=", string(Msg.Date))

		}

		//堵塞
		time.Sleep(1 * time.Second)
	}
}
