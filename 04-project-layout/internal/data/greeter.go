package data

import (
	"04-project-layout/internal/biz"
	myerrors "04-project-layout/internal/errors"
	"context"
	"database/sql"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type greeterRepo struct {
	data *Data
	logger kitlog.Logger
}


func NewGreeterRepo(data *Data, logger kitlog.Logger) biz.GreeterRepo{
	return &greeterRepo{
		data:data,
		logger:logger,
	}
}

func (g *greeterRepo) ListGreeter(ctx context.Context) ([]*biz.Greeter, error){
	query := "SELECT * from greeter"
	rows, err := g.data.Db.QueryContext(ctx, query)
	res := make([]*biz.Greeter, 0)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		level.Error(g.logger).Log("msg", err)
		return nil ,err
	}

	defer  rows.Close()

	for rows.Next(){
		var gt biz.Greeter
		err = rows.Scan(&gt.Id, &gt.Job, &gt.Word)
		if err != nil {
			level.Error(g.logger).Log("msg", err)
			return nil, err
		}
		res = append(res, &gt)
	}

	return res, nil
}

func (g *greeterRepo) GetGreeter(ctx context.Context, id int64) (*biz.Greeter, error){
	query := "SELECT * from greeter where id=?"
	res := &biz.Greeter{}
	err := g.data.Db.QueryRowContext(ctx, query, id).Scan(&res.Id, &res.Job, &res.Word)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, myerrors.NotFoundError
		}
		level.Error(g.logger).Log("msg", err)
		return nil ,err
	}

	return res, nil
}

func (g *greeterRepo) CreateGreeter(ctx context.Context, greeter *biz.Greeter) error{
	query := fmt.Sprintf("insert into greeter (job, word) values ('%s','%s')", greeter.Job, greeter.Word)
	_, err := g.data.Db.ExecContext(ctx, query)
	if err != nil {
		level.Error(g.logger).Log("msg", err)
		return err
	}

	return nil
}

func (g *greeterRepo) UpdateGreeter(ctx context.Context, id int64, greeter *biz.Greeter) error{
	query := fmt.Sprintf("update greeter set job='%s', word='%s' where id=%d", greeter.Job, greeter.Word, id)
	_, err := g.data.Db.ExecContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return myerrors.NotFoundError
		}
		level.Error(g.logger).Log("msg", err)
		return err
	}

	return nil
}

func (g *greeterRepo) DeleteGreeter(ctx context.Context, id int64) error{
	query := fmt.Sprintf("delete from  greeter where id=%d", id)
	_, err := g.data.Db.ExecContext(ctx, query)
	if err != nil {
		level.Error(g.logger).Log("msg", err)
		return err
	}

	return nil
}