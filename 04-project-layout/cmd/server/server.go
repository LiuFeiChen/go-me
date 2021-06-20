package main

import (
	"04-project-layout/app"
	"04-project-layout/internal/conf"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func parseConfig() *conf.Config{
	f,err := os.Open("./configs/config.yaml")
	if err != nil{
		panic(err)
	}
	defer  f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	conf := conf.Config{}
	yaml.Unmarshal(data, &conf)
	return &conf
}

func newApp(server *grpc.Server, cfg *conf.Config) (*app.App, error) {

	return app.New(server, cfg.Server.Grpc), nil
}

func main(){

	// log setup
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	//config setup
	cfg := parseConfig()

	//data,err := data2.NewData(cfg, logger)
	//if err != nil {
	//	os.Exit(-1)
	//}
	//
	//greeterRepo := data2.NewGreeterRepo(data, logger)
	//bizGreeterUc := biz.NewGreeterUsecase(greeterRepo, logger)
	//
	//greeterSvc := service.NewGreeterService(bizGreeterUc, logger)
	//gserver := server.NewGRPCServer(cfg, greeterSvc)
	//
	//appServer, _ := newApp(gserver, cfg)
	//appServer.Run()

	appserver,_ := createApp(cfg, logger)
	appserver.Run()
}
