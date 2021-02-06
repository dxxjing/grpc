package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	black "grpc-test/grpc-black/proto"
)

const Addr = ":9888"

func main() {

	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("client dial err:" + err.Error())
		return
	}
	defer conn.Close()
	//初始化客户端
	cli := black.NewBlackClient(conn)
	req := &black.ListReq{
		Index: 1,
		Count: 20,
	}
	//调用rpc方法
	rsp, err := cli.List(context.Background(), req)
	if err != nil {
		fmt.Println("get black list err:" + err.Error())
		return
	}
	//打印结果
	fmt.Println(rsp.GetInfo())
	for _, info := range rsp.GetInfo() {
		fmt.Println(info.GetUid(), info.GetName())
	}
}
