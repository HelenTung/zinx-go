package zinnet

import (
	"fmt"
	"strconv"

	"github.com/helenvivi/zinx/utils"
	"github.com/helenvivi/zinx/zinterface"
)

type MessageHandler struct {
	// 存放 msg id 对应的方法
	Api map[uint32]zinterface.IRouter
	//负责worker取任务的消息队列
	TaskQueue []chan zinterface.IRequest
	//pool中worker数量
	WorkPoolSize uint32
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
		//开辟等待队列
		TaskQueue: make([]chan zinterface.IRequest, utils.Globa.MaxPoolSize),
		//允许开辟的线程池最多的work数量
		WorkPoolSize: utils.Globa.MaxPoolSize,
	}
}

// 开辟任务处理pool
func (mh *MessageHandler) StartWorkPool() {
	fmt.Println("[Zinx],StartWorkPool is ready!...")
	for i := 0; i < int(mh.WorkPoolSize); i++ {
		//work依次启动
		//分配对应的队列空间到pool中worker、每个pool中的work独享一个队列、避免考虑共享加锁、解锁问题
		mh.TaskQueue[i] = make(chan zinterface.IRequest, utils.Globa.MaxPoolWorker)
		//启动、堵塞等待
		go mh.StarOnetWork(i, mh.TaskQueue[i])
	}
}

// 调动一个work携程
func (mh *MessageHandler) StarOnetWork(WorkID int, req chan zinterface.IRequest) {
	fmt.Println("req chan starting, Work ID == ", WorkID)
	for {
		r := <-req
		mh.DoMsgHandler(r)
	}
}

// 将消息发送到队列中、这种分配方法会有时序问题、即请求先后顺序存在一定问题
func (mh *MessageHandler) SendMsgTaskQueue(req zinterface.IRequest) {
	//按照conn ID进行分配
	WorkID := req.GetRequestConn().GetConnID() % mh.WorkPoolSize
	fmt.Println("Add Conn ID ", req.GetRequestConn().GetConnID(),
		"Req Msg ID ", req.GetMsgID(),
		"to WorkID ", WorkID)
	mh.TaskQueue[WorkID] <- req
}
