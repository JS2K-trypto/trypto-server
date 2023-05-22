package controller

import (
	"fmt"
	"log"
	"net/http"
	"trypto-server/model"

	"github.com/codingsince1985/geo-golang/openstreetmap"

	"github.com/gin-gonic/gin"
)

var (
	encyDnft model.EncyclopediaDNFT
)

// latitude := 37.7749
// longitude := -122.4194

func (p *Controller) CreateBadge(c *gin.Context) {

	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	geocoder := openstreetmap.Geocoder()

	location, err := geocoder.ReverseGeocode(encyDnft.Latitude, encyDnft.Longitude)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	country := location.Country
	region := location.State

	fmt.Println("Country:", country)
	fmt.Println("Region:", region)

	log.Println(encyDnft)

	//나라를 계산한 후 DB에 적재

	//동시에 DNFT 발급하는 식으로 진행

	c.JSON(http.StatusOK, location)
}

func SelectBadgeResource(){
	
}