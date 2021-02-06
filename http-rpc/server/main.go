package main

import (
	"grpc-test/rpc/config"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {

}

func (h HelloService) Hi (req string, rsp *string) error {
	*rsp = "hello: json-" + req
	return nil
}
//http rpc
func main () {
	//注册服务
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, req *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer : w,
			ReadCloser : req.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(config.Addr, nil)
}
