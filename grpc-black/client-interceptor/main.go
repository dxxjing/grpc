package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	black "grpc-test/grpc-black/proto"
	"os"
	"strings"
	"time"
)

const Addr = ":9888"
var (
	OpenTLS = true
)

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}
// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}


func main() {
	//初始化TLS
	//第二个参数 serverName 必须是签名公钥时 输入的服务名，否则认证失败
	cred, err := credentials.NewClientTLSFromFile(getParentDir() + "/keys/server.pem", "grpc-black")
	if err != nil {
		fmt.Println("cred err:" + err.Error())
		return
	}
	var opts []grpc.DialOption
	//TLS认证
	opts = append(opts, grpc.WithTransportCredentials(cred))
	// 指定自定义认证 token
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	//连接服务 并开启TLS
	conn, err := grpc.Dial(Addr, opts...)
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

// interceptor 客户端拦截器
func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}

func getParentDir() string {
	curDir, _ := os.Getwd()
	index := strings.LastIndex(curDir, string(os.PathSeparator))
	return curDir[:index]
}