package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thHandler/internal/proto"
	"thHandler/internal/znet"
	"thHandler/pkg"
)

func Init() *gin.Engine {

	r := gin.Default()

	register(r)

	return r

}

func register(router *gin.Engine) {

	// 自定义添加 middleware
	router.Use(
		RecoveryMiddleware(),
	)

	router.GET("/p", p)

}

///p?serverid=1111&user=11111&money=10
//解析收到数据以后 TCP的全局变量查找对应serverid的链接操作发送json文本数据带入参数
//
//“{"chargetype":3,"flag":-1620723712,"userid":参数user ,"reason":1,"point":参数money,"uuid":-1620723712,"permillage":参数money,"action":"payment"}”

func p(x *gin.Context) {

	id, exists := x.GetQuery("serverid")
	if !exists {
		x.JSON(http.StatusBadRequest, "")
		return
	}

	via, err := znet.Meta.GetConnMgr().GetVia(id)
	if err != nil {
		x.JSON(http.StatusBadRequest, "")
		return
	}

	defResp := &proto.HResp{
		Userid:     x.Query("user"),
		Point:      x.Query("money"),
		Permillage: x.Query("money"),
	}

	marshal, err := pkg.Cjson.Marshal(defResp)
	if err != nil {
		x.JSON(http.StatusBadRequest, "")
		return
	}
	err = via.SendMsg(0, marshal)
	if err != nil {
		fmt.Println("connection sending error")
	}
	return

}
