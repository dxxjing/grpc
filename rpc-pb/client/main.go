package main

import (
	"fmt"
	"grpc-test/rpc-pb/pb"
	"grpc-test/rpc/config"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", config.Addr)
	if err != nil  {
		fmt.Println("dial err:" + err.Error())
		return
	}
	var rsp pb.GetUserInfoRsp
	//同步阻塞调用
	err = client.Call("UserService.UserInfo", pb.GetUserInfoReq{Id: 111}, &rsp)
	if  err != nil {
		fmt.Println("call server err:" + err.Error())
		return
	}
	fmt.Println(rsp.Id)
	fmt.Printf("recv body:%#v", rsp)


	//异步调用
	/*var req = pb.GetUserInfoReq{Id: 111}
	call := client.Go("UserService.UserInfo", req, &rsp, nil)
	call = <- call.Done
	fmt.Println(call.Reply.(*pb.GetUserInfoRsp).Id, call.Reply.(*pb.GetUserInfoRsp).Name, call.Reply.(*pb.GetUserInfoRsp).Addr)
	fmt.Println(call.Args.(pb.GetUserInfoReq).Id)
	fmt.Println(rsp.Id, rsp.Name, rsp.Addr)*/

}
