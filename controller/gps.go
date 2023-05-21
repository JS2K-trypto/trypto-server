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

func (p *Controller) GpsCalc(c *gin.Context) {

	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	geocoder := openstreetmap.Geocoder()

	// latitude := 37.7749
	// longitude := -122.4194

	location, err := geocoder.ReverseGeocode(encyDnft.Latitude, encyDnft.Longitude)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	country := location.Country
	region := location.State

	fmt.Println("Country:", country)
	fmt.Println("Region:", region)

	fmt.Println("gpsCalc1")

	log.Println(encyDnft)
	log.Println("ency", &encyDnft)

	c.JSON(http.StatusOK, location)
}
