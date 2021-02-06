package main

import (
	"fmt"
	"grpc-test/rpc/config"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", config.Addr)
	if err != nil  {
		fmt.Println("dial err:" + err.Error())
		return
	}
	var rsp string
	err = client.Call("HelloService.Hi", "rpc", &rsp)
	if  err != nil {
		fmt.Println("call server err:" + err.Error())
		return
	}
	fmt.Println("recv body:", rsp)
}
