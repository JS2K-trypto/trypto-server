package controller

import (
	"fmt"
	"ginRest/go-mvc/model"
	"net/http"

	geo "github.com/kellydunn/golang-geo"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok Test"})
	return
}

func (p *Controller) Gps(c *gin.Context) {

	// Make a few points
	p1 := geo.NewPoint(42.25, 120.2)
	p2 := geo.NewPoint(30.25, 112.2)

	// find the great circle distance between them
	dist := p1.GreatCircleDistance(p2)
	fmt.Printf("great circle distance: %d\n", dist)
	return
}
