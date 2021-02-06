package main

import (
	"fmt"
	"grpc-test/rpc/config"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn,err := net.Dial("tcp", config.Addr)
	if err != nil  {
		fmt.Println("dial err:" + err.Error())
		return
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var rsp string
	err = client.Call("HelloService.Hi", "rpc", &rsp)
	if  err != nil {
		fmt.Println("call server err:" + err.Error())
		return
	}
	fmt.Println("recv body:", rsp)
}
