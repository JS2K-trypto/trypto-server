package controller

import (
	"fmt"
	"net/http"
	"time"
	"trypto-server/model"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
)

var (
	encyDnft model.EncyclopediaDNFT
	location model.Location
)

// latitude := 37.7749
// longitude := -122.4194

func (p *Controller) CreateBadge(c *gin.Context) {

	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("encyDnft", encyDnft)
	geocoder := openstreetmap.Geocoder()
	location, err := geocoder.ReverseGeocode(encyDnft.Latitude, encyDnft.Longitude)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	encyDnft.DnftCountry = location.Country
	encyDnft.DnftTime = custom

	result := p.md.MatchBadgeResource(&encyDnft)
	fmt.Println("dnft", result)
	fmt.Println("location.Country:", location.Country)

	//나라를 계산한 후 DB에 적재

	//동시에 DNFT 발급하는 식으로 진행

	c.JSON(http.StatusOK, encyDnft)
}

func (p *Controller) GetMyBadge(c *gin.Context) {
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("account", account.WalletAccount)
	result := p.md.GetMyAllDnft(account.WalletAccount)
	c.JSON(http.StatusOK, result)
}



