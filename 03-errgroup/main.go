package main


import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	fContext := context.Background()
	ctx_, cancel := context.WithCancel(fContext)
	errG, ctx := errgroup.WithContext(ctx_)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	srv := http.Server{
		Addr:              ":8081",
		Handler:           handler,
	}

	errG.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(context.Background())
	})
	errG.Go(func() error {
		return srv.ListenAndServe()
	})
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	errG.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-signalC:
				cancel()
			}
		}
	})
	if err := errG.Wait(); err != nil && !errors.Is(err, context.Canceled){
		fmt.Printf("server stop: %+V\n", err)
	}
}
