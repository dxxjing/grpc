package main

import (
	"fmt"
	"google.golang.org/grpc"
	black "grpc-test/grpc-black/proto"
	"grpc-test/grpc-black/service"
	"net"
)

const Addr = ":9888"

func main() {
	listener,err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println("listen err:" + err.Error())
		return
	}
	//实例化grpc服务
	srv := grpc.NewServer()
	//注册BlackService
	black.RegisterBlackServer(srv, service.BlackService{})

	fmt.Println("grpc server listen " + Addr)
	srv.Serve(listener)
}
