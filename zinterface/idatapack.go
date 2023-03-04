package zinterface

// 定义datapack抽象层
type IDataPack interface {
	//获取msg消息头长度
	GetHeadlen() uint32
	//封包方法，写入msgID、长度，内容、返回封装好的消息头
	PackMsg(Imessage) ([]byte, error)
	//拆包方法、先读入一定的长度（消息头长度）、再根据长度读取，返回解包好的msg
	UnPackMsg([]byte) (Imessage, error)
}
