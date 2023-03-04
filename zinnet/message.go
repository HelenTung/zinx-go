package zinnet

import "github.com/helenvivi/zinx/zinterface"

// 定义message结构
type Message struct {
	//消息的ID
	ID uint32
	//消息的长度
	Datelen uint32
	//消息的具体内容
	Date []byte
}

// Get Message ID
func (m *Message) GetMessageID() uint32 {
	return m.ID
}

// Get Message DataLen
func (m *Message) GetMessageLen() uint32 {
	return m.Datelen
}

// Get Message Date
func (m *Message) GetMessageByte() []byte {
	return m.Date
}

// Set Message ID
func (m *Message) SetMessageID(ID uint32) {
	m.ID = ID
}

// Set Message DataLen
func (m *Message) SetMessageLen(DataLen uint32) {
	m.Datelen = DataLen
}

// Set Message Date
func (m *Message) SetMessageByte(Date []byte) {
	m.Date = Date
}

//create message

func NewMsg(ID uint32, data []byte) zinterface.Imessage {
	return &Message{
		ID:      ID,
		Datelen: uint32(len(data)),
		Date:    data,
	}
}
