package zinnet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/helenvivi/zinx/utils"
	"github.com/helenvivi/zinx/zinterface"
)

// 定义结构
type DataPack struct {
}

// 获取msg消息头长度
func (data *DataPack) GetHeadlen() uint32 {
	//uint32(len) + uint32(ID) = 8
	return 8
}

// 封包方法，写入msgID、长度，内容、返回封装好的消息头
func (data *DataPack) PackMsg(msg zinterface.Imessage) ([]byte, error) {
	//创建buf对象
	databuf := bytes.NewBuffer([]byte{})
	//写入data len
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMessageLen()); err != nil {
		return nil, err
	}
	//写入data id
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMessageID()); err != nil {
		return nil, err
	}
	//写入data buf
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMessageByte()); err != nil {
		return nil, err
	}
	return databuf.Bytes(), nil
}

// 拆包方法、先读入一定的长度（消息头长度）、再根据长度读取，返回解包好的msg
func (data *DataPack) UnPackMsg(dp []byte) (zinterface.Imessage, error) {
	//创建buf对象
	databuf := bytes.NewReader(dp)

	//解压head信息、获取msg len、ID
	msg := &Message{}
	//msg len
	if err := binary.Read(databuf, binary.LittleEndian, &msg.Datelen); err != nil {
		return nil, err
	}

	//msg ID
	if err := binary.Read(databuf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	//msg data
	if err := binary.Read(databuf, binary.LittleEndian, &msg.Date); err != nil {
		return nil, err
	}
	//判断包是否过大
	if utils.Globa.MaxPackageSize > 0 && msg.Datelen > utils.Globa.MaxPackageSize {
		return nil, errors.New("msg too Large Pack more than MaxPackSize,Recv")
	}

	return msg, nil
}

// 实例化对象
func NewData() zinterface.IDataPack {
	return &DataPack{}
}
