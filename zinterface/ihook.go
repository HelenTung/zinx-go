package zinterface

type Ihook interface {
	//注册onconnstart hook方法
	SetOnConnStart(func(c Iconn))
	//注册onconnstop hook方法
	SetOnConnStop(func(c Iconn))
	//调用onconnstart hook方法
	CallOnConnStart(c Iconn)
	//调用onconnstart hook方法
	CallOnConnStop(c Iconn)
}
