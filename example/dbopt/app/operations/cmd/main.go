package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/go-kratos/kratos/pkg/log"

	v1 "dbopt/app/operations/api/v1"
	"dbopt/app/operations/conf"
	"dbopt/app/operations/internal/server/http"
	"dbopt/app/operations/internal/service"
)

func main() {
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// 注册message
	ecode.Register(v1.Message)
	// init log
	log.Init(conf.Conf.Log)
	defer log.Close()

	log.Info("=== start ===")
	//service init
	svc := service.New(conf.Conf)
	httpSrv := http.New(conf.Conf, svc)
	//init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, err := context.WithTimeout(context.Background(), 35*time.Second)
			if err != nil {
				log.Error("WithTimeout Shutdown %v", err)
			}
			svc.Close()
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv Shutdown %v", err)
			}
			log.Info("service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
