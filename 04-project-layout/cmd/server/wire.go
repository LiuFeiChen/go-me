// +build wireinject

package main

import (
	"04-project-layout/app"
	"04-project-layout/internal/biz"
	"04-project-layout/internal/conf"
	"04-project-layout/internal/data"
	"04-project-layout/internal/server"
	"04-project-layout/internal/service"
	"github.com/go-kit/kit/log"
	"github.com/google/wire"
)

func createApp(cfg *conf.Config, logger log.Logger ) (*app.App, error){

	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))

}
