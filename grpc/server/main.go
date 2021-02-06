package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test/grpc/config"
	"grpc-test/grpc/pb"
	"net"
)

//实现 UserServiceServer接口
type user struct {

}

func (u user) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoRsp, error) {
	userid := req.GetId()
	//
	return &pb.GetUserInfoRsp{
		Id: userid,
		Name: "jdx",
		Addr: "shanghai",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		fmt.Println("listen err" + err.Error())
		return
	}
	//实例化grpc
	s := grpc.NewServer()
	//注册user
	pb.RegisterUserServiceServer(s, user{})
	fmt.Println("Listen on " + config.Addr)
	s.Serve(listener)
}
