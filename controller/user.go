package controller

import (
	"fmt"
	"log"
	"net/http"
	"trypto-server/model"

	_ "trypto-server/docs"

	"github.com/gin-gonic/gin"
)

var (
	account model.Account
)

//	@BasePath		/v01
//
// UserRegisterHandler godoc
//
//	@Summary		계정주소, 닉네임, 비밀번호를 입력합니다.
//	@Tags			UserRegisterHandler
//	@Description	유저를 등록해주는 함수
//	@name			UserRegisterHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	string	true	walletAccount
//	@Param			nickName		string	true	nickName
//	@Param			password		string	true	password
//	@Router			/v01/acc/register [post]
//	@Success		200	{object}	string
func (p *Controller) UserRegisterHandler(c *gin.Context) {

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := p.md.RegisterUser(account)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User registered failed"})
	} else if account.NickName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty NickName"})
	} else if account.WalletAccount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty walletAccount"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}

}

//	@BasePath		/v01
//
// UserEditHandler godoc
//
//	@Summary		계정주소, 닉네임을 입력합니다.
//	@Tags			UserEditHandler
//	@Description	유저 프로필 업데이트하는 함수
//	@name			UserEditHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	string	true	walletAccount
//	@Param			nickName		string	true	nickName
//	@Router			/v01/acc/nickname [post]
//	@Success		200	{object}	string
func (p *Controller) UserEditHandler(c *gin.Context) {

	account.WalletAccount = c.Query("walletAccount")

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

//	@BasePath		/v01
//
// UserProfileHandler godoc
//
//	@Summary		계정주소, 닉네임을 입력합니다.
//	@Tags			UserProfileHandler
//	@Description	유저 프로필 정보를 가져는 함수
//	@name			UserProfileHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	string	true	walletAccount
//	@Param			nickName		string	true	nickName
//	@Router			/v01/acc/profile [get]
//	@Success		200	{object}	string
func (p *Controller) UserProfileHandler(c *gin.Context) {

	account.WalletAccount = c.Query("walletAccount")

	log.Println(account)
	result := p.md.GetProfile(account)
	fmt.Println("result", result)
	c.JSON(http.StatusOK, result)
}
