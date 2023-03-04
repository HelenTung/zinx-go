package zinnet

import (
	"github.com/helenvivi/zinx/zinterface"
)

// 将conn 和 data绑定起来形成一个完整的对象
type Request struct {
	//客户端建立好的链接
	conn zinterface.Iconn
	//客户端请求的数据
	msg zinterface.Imessage
}

// 获取建立好的客户端链接
func (r *Request) GetRequestConn() zinterface.Iconn {
	return r.conn
}

// 获取msg
func (r *Request) GetRequestData() []byte {
	return r.msg.GetMessageByte()
}

// 获取msg ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMessageID()
}

// 新建request
func NewRequest(conn zinterface.Iconn, msg zinterface.Imessage) zinterface.IRequest {
	return &Request{
		conn: conn,
		msg:  msg,
	}
}
