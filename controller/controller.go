package controller

import (
	"net/http"
	"trypto-server/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

func (p *Controller) GetOk(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Testing successfully"})
}
