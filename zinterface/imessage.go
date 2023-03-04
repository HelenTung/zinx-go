package zinterface

// 定义message抽象接口
type Imessage interface {
	//Get Message ID
	GetMessageID() uint32
	//Get Message DataLen
	GetMessageLen() uint32
	//Get Message Date
	GetMessageByte() []byte
	//Set Message ID
	SetMessageID(uint32)
	//Set Message DataLen
	SetMessageLen(uint32)
	//Set Message Date
	SetMessageByte([]byte)
}
