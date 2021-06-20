package service

import (
	v1 "04-project-layout/api/greeter/v1"
	"04-project-layout/internal/biz"
	"context"
	kitlog "github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/log/level"
	)

type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
	logger kitlog.Logger
}

func NewGreeterService(uc *biz.GreeterUsecase, logger kitlog.Logger) *GreeterService{
	return &GreeterService{uc: uc, logger:logger}
}


func (gs *GreeterService) GetGreeter(ctx context.Context, in * v1.GetGreeterRequest) (* v1.GetGreeterReply, error){
	greeter, err := gs.uc.GetGreeter(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetGreeterReply{Id: greeter.Id, Job: greeter.Job, Word: greeter.Word}, nil
}

func (gs *GreeterService) ListGreeter(ctx context.Context, in * v1.ListGreeterRequest) (* v1.ListGreeterReply, error){
	greeters, err := gs.uc.ListGreeter(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*v1.ListGt,0)
	for _,item := range greeters {
		res = append(res, &v1.ListGt{Id: item.Id, Job: item.Job, Word: item.Word})
	}
	return &v1.ListGreeterReply{Data: res}, nil
}
func (gs *GreeterService) CreateGreeter(ctx context.Context, in * v1.CreateGreeterRequest) (* v1.CreateGreeterReply, error){
	err := gs.uc.CreateGreeter(ctx, &biz.Greeter{
		Word: in.Word,
		Job:  in.Job,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateGreeterReply{}, nil
}

func (gs *GreeterService) UpdateGreeter(ctx context.Context, in * v1.UpdateGreeterRequest) (* v1.UpdateGreeterReply, error){
	err := gs.uc.UpdateGreeter(ctx, in.Id, &biz.Greeter{
		Word: in.Word,
		Job:  in.Job,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateGreeterReply{}, nil
}

func (gs *GreeterService) DeleteGreeter(ctx context.Context, in * v1.DeleteGreeterRequest) (* v1.DeleteGreeterReply, error){
	err := gs.uc.DeleteGreeter(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteGreeterReply{}, nil
}

