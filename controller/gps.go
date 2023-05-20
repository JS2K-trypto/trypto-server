package controller

import (
	"fmt"
	"log"
	"net/http"
	"trypto-server/model"

	"github.com/gin-gonic/gin"
)

var (
	encyDnft model.EncyclopediaDNFT
)

func (p *Controller) GpsCalc(c *gin.Context) {

	fmt.Println("gpsCalc1")
	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(encyDnft)
	log.Println("ency", &encyDnft)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
