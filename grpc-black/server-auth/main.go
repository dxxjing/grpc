package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	black "grpc-test/grpc-black/proto"
	"grpc-test/grpc-black/service"
	"net"
	"os"
	"strings"
)

const Addr = ":9888"

func main() {
	listener,err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println("listen err:" + err.Error())
		return
	}

	parentDir := getParentDir()

	//TLS认证初始化
	creds, err := credentials.NewServerTLSFromFile(parentDir + "/keys/server.pem", parentDir + "/keys/server.key")
	if err != nil {
		fmt.Println("credential err:" + err.Error())
		return
	}

	//实例化grpc服务 并开启TLS认证
	srv := grpc.NewServer(grpc.Creds(creds))
	//注册BlackService
	black.RegisterBlackServer(srv, service.BlackService{})

	fmt.Println("grpc server listen " + Addr + " with TLS")
	srv.Serve(listener)
}

func getParentDir() string {
	curDir, _ := os.Getwd()
	index := strings.LastIndex(curDir, string(os.PathSeparator))
	return curDir[:index]
}
