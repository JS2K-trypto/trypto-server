package main

import (
	"context"
	"fmt"
	ctl "ginRest/go-mvc/controller"
	_ "ginRest/go-mvc/docs"
	"ginRest/go-mvc/model"
	rt "ginRest/go-mvc/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// @BasePath /v1
// swagger API 선언
func setupSwagger(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

		mapi := &http.Server{
			Addr:           ":8888",
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		// middleware 설정
		setupSwagger(rt.Idx())

		stopSig := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)
		<-stopSig //메세지 등록

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			fmt.Println("timeout 5 seconds.")
		}
		fmt.Println("Server stop")
		//우아한 종료
	}

	g.Wait()
	//~생략

}
