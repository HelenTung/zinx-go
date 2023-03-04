package zinnet_test

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/helenvivi/zinx/zinnet"
)

func TestDataPack(t *testing.T) {
	//创建socketTCP
	listenner, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("server  listen err: ", err)
		return
	}

	//协程处理客户端业务、服务端
	go func() {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("server listen accept err: ", err)
		}
		go func(conn net.Conn) {
			//定义拆包对象dp
			dp := zinnet.NewData()
			for {
				//从conn读head msg
				headData := make([]byte, dp.GetHeadlen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head error")
					return
				}
				//从msg解包、得到head信息、data因为读到流末尾、没有更多数据、返回为nil
				msghead, err := dp.UnPackMsg(headData)
				if err != nil {
					fmt.Println("server unpack error", err)
					return
				}
				//如果msghead > 0 ,则进行msg data读取
				if msghead.GetMessageLen() > 0 {
					//msg中存在数据、二次从conn读取msg data
					msg := msghead.(*zinnet.Message)
					msg.Date = make([]byte, msg.GetMessageLen())
					_, err := io.ReadFull(conn, msg.Date)
					if err != nil {
						fmt.Println("server unpack data error", err)
						return
					}
					fmt.Println("--->Recv MsgID:", msg.ID, "DataLen=", msg.Datelen,
						"Data=", string(msg.Date))
				}
			}
		}(conn)
	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("client listen error", err)
		return
	}
	//创建封包对象
	dp := zinnet.NewData()

	//模拟msg粘包现象、封装两个msg一起发送
	//msg1
	msg1 := &zinnet.Message{
		ID:      1,
		Datelen: 3,
		Date:    []byte{'a', 'b', 'c'},
	}
	sendData1, err := dp.PackMsg(msg1)
	if err != nil {
		fmt.Println("Client pack msg1 error", err)
		return
	}
	//msg2
	msg2 := &zinnet.Message{
		ID:      2,
		Datelen: 10,
		Date:    []byte("helloworld"),
	}
	sendData2, err := dp.PackMsg(msg2)
	if err != nil {
		fmt.Println("Client pack msg2 error", err)
		return
	}
	SendData := append(sendData1, sendData2...)

	//发送
	conn.Write(SendData)

	//堵塞
	select {}
}
