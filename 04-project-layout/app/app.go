package app

import
(
	"04-project-layout/internal/conf"
	"context"
	"errors"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	ctx context.Context
	cancel func()
	logger kitlog.Logger
	server *grpc.Server
	sigs []os.Signal
	addr string
}

func New(server *grpc.Server, gconf *conf.Grpc) *App{
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx: ctx,
		cancel:cancel,
		logger: kitlog.NewLogfmtLogger(os.Stdout),
		sigs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		server: server,
		addr: gconf.Addr,
	}
}

func (a *App) Run() error{
	errG, ctx := errgroup.WithContext(a.ctx)
	//启动服务
	errG.Go(func() error {
		lis, err := net.Listen("tcp", a.addr)
		if err != nil {
			return err
		}
		return a.server.Serve(lis)
	})

	//注册信号处理
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	errG.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-signalC:
				a.server.GracefulStop()
				a.cancel()
			}
		}
	})

	if err := errG.Wait(); err != nil &&  !errors.Is(err, context.Canceled)  {
		level.Error(a.logger).Log("msg", err)
	}

	return nil
}