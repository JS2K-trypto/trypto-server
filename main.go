package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	conf "trypto-server/config"
	ctl "trypto-server/controller"

	"trypto-server/logger"
	"trypto-server/model"
	rt "trypto-server/router"

	_ "trypto-server/docs"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

	if mod, err := model.NewModel(); err != nil {
	} else if controller, err := ctl.NewCTL(mod); err != nil {
	} else if rt, err := rt.NewRouter(controller); err != nil {
	} else {

		config := conf.GetConfig("./config/config.toml")

		log.Println("config.Server.Port", config.Server.Port)
		log.Println("config.Server.Mode", config.Server.Mode)
		log.Println("config.DB[account][pass]", config.DB["account"]["pass"])
		log.Println("work", config.Work)
		log.Println("work", config.Work[0].Desc)
		log.Println("contract", config.Contract)
		log.Println("contract", config.Contract.DnftContract)

		// 로그 초기화
		if err := logger.InitLogger(config); err != nil {
			fmt.Printf("init logger failed, err:%v\n", err)
			return
		}
		logger.Debug("ready server....")

		mapi := &http.Server{
			Addr:           config.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    0, //  5 * time.Second, 이전 값 현재 값은 테스트를 위해 설정함
			WriteTimeout:   0, // 10 * time.Second, 이전 값 현재 값은 테스트를 위해 설정함
			MaxHeaderBytes: 1 << 20,
		}

		//고루틴 서버 동작
		g.Go(func() error {

			//return mapi.ListenAndServeTLS("./127.0.0.1.pem", "./127.0.0.1-key.pem")
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit //메세지 등록

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}
		logger.Info("Server exiting")
		//우아한 종료

	}

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}
