package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test/grpc/config"
	"grpc-test/grpc/pb"
)

func main() {
	//连接
	conn,err := grpc.Dial(config.Addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial err " + err.Error())
		return
	}
	defer conn.Close()
	//初始化客户端
	cli := pb.NewUserServiceClient(conn)
	//调用方法
	req := &pb.GetUserInfoReq{Id: 112}
	rsp, err := cli.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("get user info err" + err.Error())
		return
	}
	fmt.Println(rsp.GetId(), rsp.GetName(), rsp.GetAddr())
}
