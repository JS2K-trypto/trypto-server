package controller

import (
	"fmt"
	"net/http"
	"time"
	"trypto-server/model"

	"github.com/gin-gonic/gin"
)

var (
	tripPlan model.TripPlan
)

func (p *Controller) CreateTripPlan(c *gin.Context) {

	if err := c.ShouldBindJSON(&tripPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("tripPlan", tripPlan)
	tripPlan.TripTime = custom

	res := p.md.InsertTripPlan(&tripPlan)

	c.JSON(http.StatusOK, res)
}
