package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// api handler test
func apiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}


func main(){

	ctx_, cancel := context.WithCancel(context.Background())
	errG, ctx := errgroup.WithContext(ctx_)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/test", apiHandler)

	srv := http.Server{
		Addr:              ":8081",
		Handler:           mux,
	}

	//当收到退出信号时，cancel()，此时服务shutdown
	errG.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(context.Background())
	})
	//启动服务
	errG.Go(func() error {
		return srv.ListenAndServe()
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
				cancel()
			}
		}
	})

	if err := errG.Wait(); err != nil &&  !errors.Is(err, context.Canceled)  {
		log.Fatal(err)
	}
}
