package main

import (
	"context"
	"fmt"
	"rpcServer/controller"

	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
	pb "rpcServer/pb"
)

const (
	port = ":50051"
)

// auth 验证Token
func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appuid string
		appkey string
	)
	if val, ok := md["appuid"]; ok {
		appuid = val[0]
	}
	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}
	if appuid != "100" || appkey != "i am key" {
		return status.Errorf(codes.Unauthenticated, "Token认证失败")
	}
	return nil
}

func main() {
	grpc.EnableTracing = true
	var (
		opts []grpc.ServerOption
	)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		grpclog.Fatal(err)
	}

	// 添加TLS认证
	creds, err := credentials.NewServerTLSFromFile("./ssl/server.pem", "./ssl/server.key")
	if err != nil {
		grpclog.Fatal(err)
	}
	opts = append(opts, grpc.Creds(creds))

	// 注册interceptor
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		// 继续处理其他请求
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	s := grpc.NewServer(opts...)
	// 注册 GreeterService
	pb.RegisterGreeterServer(s, &controller.Server{})
	// 注册 UserService
	pb.RegisterUserServer(s, &controller.User{})

	pb.RegisterStreamServiceServer(s, &controller.StreamService{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	fmt.Printf("Listen on %s with TLS\n", port)

	// 开启trace
	go startTrace()
	err = s.Serve(lis)
	if err != nil {
		grpclog.Fatal(err)
	}
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":50052", nil)
	fmt.Println("Trace listen on 50052")
}
