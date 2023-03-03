package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client starting...")
	time.Sleep(1 * time.Second)

	//建立链接、获取conn
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client start error ", err, " exit!")
		return
	}
	//循环
	for {
		//向conn write 数据
		_, err := conn.Write([]byte("hello,the is zinx v3.0.."))
		if err != nil {
			fmt.Println("conn write error : ", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error : ", err)
			return
		}
		fmt.Printf("server call back : %s ,cnt = %d\n", buf, cnt)
		//堵塞
		time.Sleep(1 * time.Second)
	}
}
