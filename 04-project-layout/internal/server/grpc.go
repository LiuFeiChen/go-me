package server

import (
	"04-project-layout/internal/conf"
	"04-project-layout/internal/service"
	"google.golang.org/grpc"
	pb "04-project-layout/api/greeter/v1"
)

func NewGRPCServer(conf *conf.Config, greeter *service.GreeterService) *grpc.Server{
	svr := grpc.NewServer()
	pb.RegisterGreeterServer(svr, greeter)
	return svr
}