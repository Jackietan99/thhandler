package znet

import (
	"fmt"
	"thHandler/internal/znet/ziface"
)

//ping test 自定义路由
type PingRouter struct {
	BaseRouter
}

//Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {

}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgID=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("{\"action\":\"respond\",\"recode\":0}"))
	if err != nil {
		fmt.Println("Handle SendMsg err: ", err)
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {

}

type ActionRouter struct {
	BaseRouter
}

//Test PreHandle
func (this *ActionRouter) PreHandle(request ziface.IRequest) {

}

//Test Handle
func (this *ActionRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call ActionRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgID=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("{\"action\":\"r_setstep\",\"recode\":0}"))
	if err != nil {
		fmt.Println("Handle SendMsg err: ", err)
	}
}

//Test PostHandle
func (this *ActionRouter) PostHandle(request ziface.IRequest) {

}
