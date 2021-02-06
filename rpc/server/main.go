package main

import (
	"fmt"
	"grpc-test/rpc/config"
	"net"
	"net/rpc"
)

type HelloService struct {

}

func (h HelloService) Hi (req string, rsp *string) error {
	*rsp = "hello: " + req
	return nil
}
//默认采用 go gob编码 不利于跨语言
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
		go rpc.ServeConn(conn)
	}

}
