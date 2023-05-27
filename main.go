package main

import (
	"context"
	"flag"
	"fmt"
	"io"
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

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// @BasePath /v1
// swagger API 선언
// func setupSwagger(r *gin.Engine) {
// 	r.GET("/", func(c *gin.Context) {
// 		c.Redirect(http.StatusFound, "/swagger/index.html")
// 	})

// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// }

const (
	htmlIndex    = `<html><body>Welcome!</body></html>`
	inProduction = true
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, htmlIndex)
}

func main() {

	//model 모듈 선언
	if mod, err := model.NewModel(); err != nil {
		//~생략
	} else if controller, err := ctl.NewCTL(mod); err != nil { //controller 모듈 설정
		//~생략
	} else if rt, err := rt.NewRouter(controller); err != nil { //router 모듈 설정
		//~생략
	} else {
		var configFlag = flag.String("config", "./config/.config.toml", "toml file to use for configuration")
		flag.Parse()
		cf := conf.GetConfig(*configFlag)
		//config := conf.GetConfig("./config/.config.toml")
		//cf := conf.NewConfig(*configFlag)
		// fmt.Println("config.Server.Port", config.Server.Port)
		// fmt.Println("config.Server.Mode", config.Server.Mode)
		// fmt.Println("config.DB[account][pass]", config.DB["account"]["pass"])
		// fmt.Println("work", config.Work)
		// fmt.Println("work", config.Work[0].Desc)

		// 로그 초기화
		if err := logger.InitLogger(cf); err != nil {
			fmt.Printf("init logger failed, err:%v\n", err)
			return
		}

		logger.Debug("ready server....")
		//dataDir := "."
		//tlsConfig := &tls.Config{
		// 	ClientAuth: tls.RequireAnyClientCert,
		// }
		//http 서버 설정 변수
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    0, //  5 * time.Second, 이전 값 현재 값은 테스트를 위해 설정함
			WriteTimeout:   0, // 10 * time.Second, 이전 값 현재 값은 테스트를 위해 설정함
			MaxHeaderBytes: 1 << 20,
			//TLSConfig:      tlsConfig,
		}
		// m := &autocert.Manager{
		// 	Prompt:     autocert.AcceptTOS,
		// 	HostPolicy: autocert.HostWhitelist("example.com", "example2.com"),
		// 	Cache:      autocert.DirCache(dataDir),
		// }
		// mapi.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		//고루틴 서버 동작
		g.Go(func() error {
			//return mapi.ListenAndServeTLS("cert.pem", "key.pem")
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
