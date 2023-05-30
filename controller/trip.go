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
//	@Summary				계정주소, 제목, 설명, 메모, 여행사진등을 입력해서 여행계획을 만들어줍니다.
//	@Tags					CreateTripPlan
//	@Description			사용자 위치를 참고해서 뱃지를 발급하는 함수
//	@name					CreateTripPlan
//	@Accept					json
//	@Produce				json
<<<<<<< Updated upstream
//	@Param					walletAccount	string 	path	true	walletAccount
//	@Param					travelTitle		string	path 	true	travelTitle
//	@Param					tripDescription	string	path	true	tripDescription
//	@Param					tripMemo		string	path	true	tripMemo
//	@Param					tripImgSrc		string	path	true	tripImgSrc
=======
//	@Param					walletAccount path  string true "walletAccount"
//	@Param					travelTitle	path	string	true	"travelTitle"
//	@Param					tripDescription	path string	true	"tripDescription"
//	@Param					tripMemo	path	string	true	"tripMemo"
//	@Param					tripImgSrc	path	string	true	"tripImgSrc"
>>>>>>> Stashed changes
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
//	@Tags			GetMyTrip
//	@Description	나의 여행계획을  DB에서 가져오는 함수, 계정주소로 파악한다.
//	@name			GetMyTrip
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	string 	 path	true	walletAccount
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

// GetMyTrip godoc
//
//	@BasePath		/v01
//	@Summary		모든 여행계획을 가져오는 함수
//	@Tags			GetAllTrip
//	@Description	나의 여행계획을  DB에서 가져오는 함수, 계정주소로 파악한다.
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
//	@Tags			GetAllTrip
//	@Description	triptitle중 일치하는 문자열에 대해 콘텐츠를 리스폰스해주는 검색 API, 단어 단위로 구현
//	@name			GetAllTrip
//	@Accept			json
//	@Produce		json
//	@Router			/v01/trip/search [get]
//	@Success		200	{object}	string
func (p *Controller) SearchTrip(c *gin.Context) {

	searchQuery = c.Query("q")

	res := p.md.SearchTrip(searchQuery)
	fmt.Println("res", res)
	c.JSON(http.StatusOK, res)

}
