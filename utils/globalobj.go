package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/helenvivi/zinx/zinterface"
)

// 定义有关存储zinx框架的全局参数、使用zinx.json文件进行配置
type GlobaObj struct {
	//TODO:Server
	//服务器对象
	TcpServer zinterface.IServer
	//服务器名称
	Name string
	//绑定的IP地址
	Host string
	//port 监听的端口
	TcpPort int

	//TODO:zinx
	//当前zinx版本
	Version string
	//当前服务器允许的最大链接数目
	MaxConn int
	//数据包的最大值
	MaxPackageSize uint32
}

var Globa *GlobaObj

func (g *GlobaObj) ReadZinxJson() {
	data, err := os.ReadFile("../conf/zinx.json")
	if err != nil {
		fmt.Println("conf/zinx.json Not ReadLoad!,panic")
		panic(err)
	}
	if err := json.Unmarshal(data, &Globa); err != nil {
		panic(err)
	}
}

func init() {
	//zinx.json没有初始化、默认加载值
	Globa = &GlobaObj{
		Name:           "zinx",
		Version:        "V4.0",
		TcpPort:        8080,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	Globa.ReadZinxJson()
}
