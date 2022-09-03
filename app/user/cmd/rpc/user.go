package main

import (
	"flag"
	"fmt"
	"quora/app/user/cmd/rpc/internal/config"
	"quora/app/user/cmd/rpc/internal/server"
	"quora/app/user/cmd/rpc/internal/svc"
	"quora/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")
var c config.Config

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//s.AddUnaryInterceptors(MetadataInterceptors)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

//func MetadataInterceptors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//	if md, ok := metadata.FromIncomingContext(ctx); ok {
//		if len(md["username"]) > 0 && md["username"][0] == c.MetaData.UserName &&
//			len(md["password"]) > 0 && md["password"][0] == c.MetaData.PassWord {
//			return handler(ctx, req)
//		}
//
//		return nil, errors.New("密钥配置错误")
//
//	}
//	fmt.Scanln()
//	return req, err
//}
