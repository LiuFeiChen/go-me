package data

import (
	"04-project-layout/internal/conf"
	"database/sql"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

type Data struct {
	Db *sql.DB
}

func NewData(conf *conf.Config , logger kitlog.Logger ) (*Data, error){

	db, err := sql.Open(conf.Data.Database.Driver, conf.Data.Database.Source)
	if err != nil {
		level.Error(logger).Log("msg", err.Error())
		return  nil, err
	}

	return &Data{Db:db}, nil
}