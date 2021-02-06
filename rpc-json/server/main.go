package main

import (
	"fmt"
	"grpc-test/rpc/config"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {

}

func (h HelloService) Hi (req string, rsp *string) error {
	*rsp = "hello: json-" + req
	return nil
}
//采用 json编码 跨语言
//nc -l 9800 启动rpc服务 再客户端调用该服务 将再服务端出现
//{"method":"HelloService.Hi","params":["rpc"],"id":0}
func main () {
	//注册服务
	rpc.RegisterName("HelloService", new(HelloService))
	//监听端口
	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		fmt.Println("listen err:" + err.Error())
		return
	}
	for {
		//接受请求
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:" + err.Error())
			return
		}
		//内部无休止的接受连接 并用gourtine 处理每个请求
		//采用json 服务端就这里不同
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
