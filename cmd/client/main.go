package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8787")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}

	for {

		////发封包message消息
		//dp := znet.NewDataPack()
		//msg, _ := dp.Pack(znet.NewMsgPackage(1, []byte("{\"action\":\"r_setstep\",\"recode\":0}")))
		//
		//e := []byte{'0','0'}
		//
		//h = append(h, msg)
		//_, err := conn.Write(msg)
		//if err != nil {
		//	fmt.Println("write error err ", err)
		//	return
		//}
		//
		////先读出流中的head部分
		//headData := make([]byte, dp.GetHeadLen())
		//_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		//if err != nil {
		//	fmt.Println("read head error")
		//	break
		//}
		////将headData字节流 拆包到msg中
		//msgHead, err := dp.Unpack(headData)
		//if err != nil {
		//	fmt.Println("server unpack err:", err)
		//	return
		//}
		//
		//if msgHead.GetDataLen() > 0 {
		//	//msg 是有data数据的，需要再次读取data数据
		//	msg := msgHead.(*znet.Message)
		//	msg.Data = make([]byte, msg.GetDataLen())
		//
		//	//根据dataLen从io中读取字节流
		//	_, err := io.ReadFull(conn, msg.Data)
		//	if err != nil {
		//		fmt.Println("server unpack data err:", err)
		//		return
		//	}
		//
		//	fmt.Println("==> Test Router:[Ping] Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
		//}
		//
		//time.Sleep(1 * time.Second)

		m := []byte("2A000000007B22616374696F6E223A2022725F73657473746570222C20227265636F6465223A20307D00")

		_, _ = conn.Write(m)

		time.Sleep(1 * time.Second)

	}
}
