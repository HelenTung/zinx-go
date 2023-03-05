package zinnet

import (
	"fmt"
	"strconv"

	"github.com/helenvivi/zinx/zinterface"
)

type MessageHandler struct {
	// 存放 msg id 对应的方法
	Api map[uint32]zinterface.IRouter
}

// 执行对应路由的方法
func (mh *MessageHandler) DoMsgHandler(r zinterface.IRequest) {
	handler, ok := mh.Api[r.GetMsgID()]
	if !ok {
		//msg id未注册
		panic("MSG API NOT REGISTER" + strconv.Itoa(int(r.GetMsgID())))
	}
	handler.PreHandle(r)
	handler.Handle(r)
	handler.PostHandle(r)
}

// 添加处理逻辑到路由
func (mh *MessageHandler) AddRouter(msgID uint32, router zinterface.IRouter) {
	if _, ok := mh.Api[msgID]; ok {
		//msg id已经注册
		panic("msg api log in" + strconv.Itoa(int(msgID)))
	}
	//绑定
	mh.Api[msgID] = router
	fmt.Println("add api MsgID = ", msgID, "router = ", router)
}

// 返回一个新的 MsgHandler 实例
func NewMessageHandler() zinterface.IMsgHandler {
	return &MessageHandler{
		Api: make(map[uint32]zinterface.IRouter),
	}
}
