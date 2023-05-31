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
	tripPlan    model.TripPlan
	searchQuery string
)

// CreateTripPlan godoc
//
//	@BasePath				/v01
//	@Summary				사용자가 입력한 여행계획표 데이터를 입력받고 그대로 여행계획을 생성해준다.
//	@Tags					CreateTripPlan
//	@Description			계정주소, 제목, 설명, 메모, 여행사진등을 입력해서 여행계획을 만들어줍니다.
//	@name					CreateTripPlan(나의 여행계획표 생성하기)
//	@Accept					json
//	@Produce				json
//	@Param					walletAccount	path	string 		true	"walletAccount",
//	@Param					travelTitle		path	string	 	true	"travelTitle",
//	@Param					tripDescription	path	string		true	"tripDescription",
//	@Param					tripMemo		path	string		true	"tripMemo",
//	@Param					tripImgSrc		path	string		true	"tripImgSrc"
//	@Router/v01/trip/myplan	[post]
//	@Success				200	{object}	string
func (p *Controller) CreateTripPlan(c *gin.Context) {

	if err := c.ShouldBindJSON(&tripPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("tripPlan", tripPlan)
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("tripPlan", tripPlan)
	tripPlan.TripDeparture = custom
	tripPlan.TripArrival = custom

	res := p.md.InsertTripPlan(&tripPlan)

	c.JSON(http.StatusOK, res)

}

// GetMyTrip godoc
//
//	@BasePath		/v01
//	@Summary		나의 여행계획을 가져오는 함수
//	@Tags			GetMyTrip(나의 여행계획 가져오기)
//	@Description	나의 여행계획을 MongoDB에서 가져오는 함수, 계정주소로 파악한 후 가져온다.
//	@name			GetMyTrip
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount"
//	@Router			/v01/trip/myplan [get]
//	@Success		200	{object}	string
func (p *Controller) GetMyTrip(c *gin.Context) {

	tripPlan.WalletAccount = c.Query("walletAccount")

	fmt.Println("tripPlan", tripPlan)
	res := p.md.SelectMyTrip(tripPlan.WalletAccount)
	//fmt.Println("len, len(res.Arr)", len(res.Arr))
	//c.JSON(http.StatusOK, res)

	if len(res.Arr) > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty TripPlan"})
	}

}

// GetAllTrip godoc
//
//	@BasePath		/v01
//	@Summary		모든 여행계획을 가져오는 함수
//	@Tags			GetAllTrip(전체 여행계획 가져오기)
//	@Description    모든 여행계획을  MongoDB에서 가져오는 함수. 아무 파라미터가 없다 전체를 조회한다.
//	@name			GetAllTrip
//	@Accept			json
//	@Produce		json
//	@Router			/v01/trip/allplan [get]
//	@Success		200	{object}	string
func (p *Controller) GetAllTrip(c *gin.Context) {

	res := p.md.SelectAllTrip()
	fmt.Println(len(res))
	if len(res) > 0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty TripPlan"})
	}

}

// SearchTrip godoc
//
//	@BasePath		/v01
//	@Summary		여행계획을 검색하는 API
//	@Tags			SearchTrip(여행계획 단어단위 검색하기)
//	@Description	여행계획의 제목 중 일치하는 문자열에 대해 콘텐츠를 리스폰스해주는 검색 API, 단어 단위로 구현, 예를 들어 Paris로 무작정이라고 하면 "Paris로" 까지 입력해야된다. q="Paris로" 이런식으로 입력하면 된다.
//	@name			SearchTrip
//	@Accept			json
//	@Produce		json
//	@Param			q	path	string 	 	true	"q"
//	@Router			/v01/trip/search [get]
//	@Success		200	{object}	string
func (p *Controller) SearchTrip(c *gin.Context) {

	searchQuery = c.Query("q")

	res := p.md.SearchTrip(searchQuery)
	fmt.Println("res", res)
	c.JSON(http.StatusOK, res)

}
