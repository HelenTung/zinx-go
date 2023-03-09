package zinnet

import (
	"errors"
	"fmt"
	"sync"

	"github.com/helenvivi/zinx/zinterface"
)

type ConnManger struct {
	connections map[uint32]zinterface.Iconn
	connLock    sync.RWMutex
}

// 删除链接
func (cm *ConnManger) DeleteConn(conn zinterface.Iconn) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	//删除
	delete(cm.connections, conn.GetConnID())
	fmt.Println("Delete ConnID ", conn.GetConnID())
}

// 添加链接
func (cm *ConnManger) AddteConn(conn zinterface.Iconn) {
	//加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//加入
	cm.connections[conn.GetConnID()] = conn
	fmt.Println("Add ConnID ", conn.GetConnID())
}

// connid ——> 获取conn
func (cm *ConnManger) GetConn(ConnID uint32) (zinterface.Iconn, error) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	if conn, ok := cm.connections[ConnID]; ok {
		fmt.Println("ConnID ", conn.GetConnID())
		return conn, nil
	} else {
		return nil, errors.New("connection NOT FOUND")
	}
}

// 得到链接总数
func (cm *ConnManger) GetConnNum() int {
	return len(cm.connections)
}

// 清空并终止所有链接
func (cm *ConnManger) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for k, v := range cm.connections {
		//stop
		v.Stop()
		//delete
		delete(cm.connections, k)
		fmt.Println("Clear All connection!")

	}
}

func NewConnManger() zinterface.IConnManger {
	return &ConnManger{
		connections: make(map[uint32]zinterface.Iconn),
		connLock:    sync.RWMutex{},
	}
}
