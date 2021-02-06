package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	black "grpc-test/grpc-black/proto"
	"os"
	"strings"
)

const Addr = ":9888"

func main() {
	//初始化TLS
	//第二个参数 serverName 必须是签名公钥时 输入的服务名，否则认证失败
	cred, err := credentials.NewClientTLSFromFile(getParentDir() + "/keys/server.pem", "grpc-black")
	if err != nil {
		fmt.Println("cred err:" + err.Error())
		return
	}
	//连接服务 并开启TLS
	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(cred))
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

func getParentDir() string {
	curDir, _ := os.Getwd()
	index := strings.LastIndex(curDir, string(os.PathSeparator))
	return curDir[:index]
}