package router

import (
	"fmt"
	ctl "trypto-server/controller"

	"trypto-server/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당

	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func LiteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next() // 다음 요청 진행
	}
}

// 실제 라우팅
func (p *Router) Idx() *gin.Engine {
	//~생략
	e := gin.Default() //gin 선언

	// r.Use(gin.Logger())   //gin 내부 log, logger 미들웨어 사용 선언
	// r.Use(gin.Recovery()) //gin 내부 recover, recovery 미들웨어 사용 - 패닉복구

	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS()) //crossdomain 미들웨어 사용 등록
	logger.Info("start server")

	routerAdm := e.Group("/v01/badge", LiteAuth())
	{
		fmt.Println(routerAdm)
		routerAdm.POST("/issue", p.ct.CreateBadge) // controller 패키지의 실제 처리 함수
		routerAdm.GET("/user", p.ct.GetMyBadge)
		//routerAdm.POST("/ok1", p.ct.GetOk)     // controller 패키지의 실제 처리 함수
	}

	routerAcc := e.Group("/v01/acc", LiteAuth())
	{
		routerAcc.POST("register", p.ct.UserRegisterHandler)
		routerAcc.POST("nickname", p.ct.UserEditHandler)
		routerAcc.GET("profile", p.ct.UserProfileHandler)
	}

	routerTrip := e.Group("/v01/trip", LiteAuth())
	{
		routerTrip.POST("/myplan", p.ct.CreateTripPlan)
	}

	return e
}
