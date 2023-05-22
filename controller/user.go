package controller

import (
	"log"
	"net/http"
	"trypto-server/model"

	"github.com/gin-gonic/gin"
)

var (
	account model.Account
)

func (p *Controller) UserRegisterHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	log.Println(account)
	p.md.RegisterUser(account)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (p *Controller) UserProfileHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	log.Println(account)
	result := p.md.GetProfile(account)
	c.JSON(http.StatusOK, result)
}
