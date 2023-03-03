package zinnet

import "github.com/helenvivi/zinx/zinterface"

// 将conn 和 data绑定起来形成一个完整的对象
type Request struct {
	//客户端建立好的链接
	conn zinterface.Iconn
	//客户端请求的数据
	data []byte
}

// 获取建立好的客户端链接
func (r *Request) GetRequestConn() zinterface.Iconn {
	return r.conn
}

// 获取客户端请求的数据
func (r *Request) GetRequestByte() []byte {
	return r.data
}

// 新建request
func NewRequest(conn zinterface.Iconn, data []byte) zinterface.IRequest {
	return &Request{
		conn: conn,
		data: data,
	}
}
