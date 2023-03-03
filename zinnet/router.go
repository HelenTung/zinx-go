package zinnet

import "github.com/helenvivi/zinx/zinterface"

// 基类router、实现router时、可以根据request需要对baserouter的方法进行重写
type BaseRouter struct {
}

// 处理业务之前的hook（钩子）方法
func (r *BaseRouter) PreHandle(req zinterface.IRequest) {
}

// 处理conn业务的主方法
func (r *BaseRouter) Handle(req zinterface.IRequest) {
}

// 处理业务之后的hook（钩子）方法
func (r *BaseRouter) PostHandle(req zinterface.IRequest) {
}
