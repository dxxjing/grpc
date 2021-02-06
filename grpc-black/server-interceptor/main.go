package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	black "grpc-test/grpc-black/proto"
	"grpc-test/grpc-black/service"
	"net"
	"net/http"
	"os"
	"strings"
)

const Addr = ":9888"
//interceptor拦截器 类似 middleware 用于处理 TLS token 等认证
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
	//实例化grpc服务 并开启TLS认证、开启interceptor
	srv := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
	//注册BlackService
	black.RegisterBlackServer(srv, service.BlackService{})
	//启动trace
	go startTrace()

	fmt.Println("grpc server listen " + Addr + " with TLS + token")
	srv.Serve(listener)
}

//token interceptor = middleware
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	if err := auth(ctx); err != nil {
		fmt.Println("111")
		return nil, err
	}
	//继续处理请求
	return handler(ctx, req)

}

//自定义token认证
func auth(ctx context.Context) (err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New(fmt.Sprintf("无token认证信息:%d", codes.Unauthenticated))
	}

	var (
		appid string
		appkey string
	)
	if val,ok := md["appid"]; ok {
		appid = val[0]
	}
	if val,ok := md["appkey"]; ok {
		appkey = val[0]
	}
	//比对token
	if appid != "101010" || appkey != "i am key" {
		return errors.New(fmt.Sprintf("token 无效 appid = %s, appkey = %s", appid, appkey))
	}
	return nil
}
//追踪请求
//服务端事件查看 localhost:50051/debug/events
//请求日志信息查看 localhost:50051/debug/requests
func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":50051", nil)
	fmt.Println("Trace listen on 50051")
}

func getParentDir() string {
	curDir, _ := os.Getwd()
	index := strings.LastIndex(curDir, string(os.PathSeparator))
	return curDir[:index]
}
