package controller

import (
	"fmt"
	"net/http"
	"time"
	"trypto-server/model"

	_ "trypto-server/docs"

	"github.com/gin-gonic/gin"
)

var (
	tripPlan model.TripPlan
)

// CreateTripPlan godoc
//
//	@BasePath				/v01
//	@Summary				계정주소, 제목, 설명, 메모, 여행사진등을 입력해서 여행계획을 만들어줍니다.
//	@Tags					CreateTripPlan
//	@Description			사용자 위치를 참고해서 뱃지를 발급하는 함수
//	@name					CreateTripPlan
//	@Accept					json
//	@Produce				json
//	@Param					walletAccount	string	true	"walletAccount"
//	@Param					travelTitle		string	true	"travelTitle"
//	@Param					tripDescription	string	true	"tripDescription"
//	@Param					tripMemo		string	true	"tripMemo"
//	@Param					tripImgSrc		string	true	"tripImgSrc"
//	@Router/v01/trip/myplan	[post]
//	@Success				200	{object}	string
func (p *Controller) CreateTripPlan(c *gin.Context) {

	tripPlan.WalletAccount = c.Query("walletAccount")

	fmt.Println("tripPlan", tripPlan)
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("tripPlan", tripPlan)
	tripPlan.TripDeparture = custom
	tripPlan.TripArrival = custom

	res := p.md.InsertTripPlan(&tripPlan)

	c.JSON(http.StatusOK, res)
	//c.JSON(http.StatusOK, "hello ")
}

func (p *Controller) GetMyTrip(c *gin.Context) {

	tripPlan.WalletAccount = c.Query("walletAccount")

	fmt.Println("tripPlan", tripPlan)
	res := p.md.SelectMyTrip(tripPlan.WalletAccount)

	c.JSON(http.StatusOK, res)

}

func (p *Controller) GetAllTrip(c *gin.Context) {

	res := p.md.SelectAllTrip()
	fmt.Println(len(res))
	if len(res) > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty TripPlan"})
	}

}
