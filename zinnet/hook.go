package zinnet

import (
	"fmt"

	"github.com/helenvivi/zinx/zinterface"
)

// 钩子方法
type Hook struct {
	//在服务启动之前调用
	OnConnStart func(conn zinterface.Iconn)
	//在服务启动之后调用
	OnConnStop func(conn zinterface.Iconn)
}

// 注册onconnstart hook方法
func (h *Hook) SetOnConnStart(hf func(c zinterface.Iconn)) {
	h.OnConnStart = hf
}

// 注册onconnstop hook方法
func (h *Hook) SetOnConnStop(hf func(c zinterface.Iconn)) {
	h.OnConnStop = hf
}

// 调用onconnstart hook方法
func (h *Hook) CallOnConnStart(c zinterface.Iconn) {
	if h.OnConnStart != nil {
		fmt.Println("\n----->Call OnConnstart()...")
		h.OnConnStart(c)
	}
}

// 调用onconnstart hook方法
func (h *Hook) CallOnConnStop(c zinterface.Iconn) {
	if h.OnConnStop != nil {
		fmt.Println("\n----->Call OnConnstop()...")
		h.OnConnStop(c)
	}
}

// 实例化
func NewHook() zinterface.Ihook {
	return &Hook{}
}
