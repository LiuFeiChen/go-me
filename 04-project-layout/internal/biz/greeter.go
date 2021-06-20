package biz

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
)

type Greeter struct {
	Id int64
	Word string
	Job string
}

type GreeterRepo interface {
	ListGreeter(ctx context.Context) ([]*Greeter, error)
	GetGreeter(ctx context.Context, id int64) (*Greeter, error)
	CreateGreeter(ctx context.Context, article *Greeter) error
	UpdateGreeter(ctx context.Context, id int64, article *Greeter) error
	DeleteGreeter(ctx context.Context, id int64) error
}

type GreeterUsecase struct {
	repo GreeterRepo
	logger kitlog.Logger
}


func NewGreeterUsecase(repo GreeterRepo,  logger kitlog.Logger) *GreeterUsecase{
	return &GreeterUsecase{repo: repo, logger:logger}
}

func (uc *GreeterUsecase)ListGreeter(ctx context.Context) ([]*Greeter, error){
	return uc.repo.ListGreeter(ctx)
}

func (uc *GreeterUsecase)GetGreeter(ctx context.Context, id int64) (*Greeter, error){
	return uc.repo.GetGreeter(ctx, id)
}

func (uc *GreeterUsecase)CreateGreeter(ctx context.Context, greeter *Greeter) error{
	return uc.repo.CreateGreeter(ctx, greeter)
}

func (uc *GreeterUsecase)UpdateGreeter(ctx context.Context, id int64, greeter *Greeter) error{
	return uc.repo.UpdateGreeter(ctx, id, greeter)
}

func (uc *GreeterUsecase)DeleteGreeter(ctx context.Context, id int64) error{
	return uc.repo.DeleteGreeter(ctx, id)
}