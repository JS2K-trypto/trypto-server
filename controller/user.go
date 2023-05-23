package controller

import (
	"fmt"
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

	result := p.md.RegisterUser(account)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User registered failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func (p *Controller) UserEditHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("account.WalletAccount", account.WalletAccount)
	fmt.Println("account", account)
	checkUser := p.md.MatchUser(account.WalletAccount)
	fmt.Println("checkUser", checkUser)
	if checkUser == true {
		p.md.UpdateUser(account)
		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User updated failed"})
	}

}

func (p *Controller) UserProfileHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	log.Println(account)
	result := p.md.GetProfile(account)
	fmt.Println("result", result)
	c.JSON(http.StatusOK, result)
}
