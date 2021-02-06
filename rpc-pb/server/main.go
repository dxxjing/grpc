package main

import (
	"fmt"
	user "grpc-test/rpc-pb/pb"
	"grpc-test/rpc/config"
	"net"
	"net/rpc"
)

type UserService struct {

}

func (u UserService) UserInfo (req user.GetUserInfoReq, rsp *user.GetUserInfoRsp) error {

	userId := req.GetId()
	//todo select * from user where id = userId
	fmt.Println(userId)
	//不能使用如下方式 使用后rsp 地址被改变
	//fmt.Printf("begin-%p\n" , rsp)
	//rsp = &(user.GetUserInfoRsp{
	//	Id: userId,
	//	Name: "jdx",
	//	Addr: "anhui",
	//})
	//fmt.Printf("end-%p\n" , rsp)
	rsp.Id = userId
	rsp.Name = "jdx"
	rsp.Addr = "anhui"

	return nil
}

func main() {
	rpc.RegisterName("UserService", new(UserService))

	listener, err := net.Listen("tcp", config.Addr)
	if  err != nil {
		fmt.Println("listen err:" + err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:" + err.Error())
			return
		}
		go rpc.ServeConn(conn)
	}
}
