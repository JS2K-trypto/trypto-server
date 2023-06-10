package controller

import (
	"log"
	"net/http"

	"trypto-server/model"

	_ "trypto-server/docs"

	"github.com/gin-gonic/gin"
)

var account model.Account

//	@BasePath		/v01
//
// UserRegisterHandler godoc
//
//	@Summary		Enter your account address and nickname.
//	@Tags			UserRegisterHandler( You can register or edit your nickname. )
//	@Description	This function allows you to register and edit user nicknames, and you can do so by entering your wallet account and entering the desired nickname.
//	@name			UserRegisterHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount 	path	string	true	"walletAccount",
//	@Param			nickName		path	string	true	"nickName"
//	@Router			/v01/acc/register [post]
//	@Success		200	{array} model.Account
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
// UserProfileHandler godoc
//
//	@Summary		Enter your account address to get your profile information.
//	@Tags			UserProfileHandler(Get my profile information)
//	@Description	Fetches user profile information. Gets the following information [nickname, number of travel plans, number of Dynamic NFTs, number of likes, number of comments].
//	@name			UserProfileHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string	true	"walletAccount",
//	@Param			nickName		path	string	true	"nickName"
//	@Router			/v01/acc/profile [get]
//	@Success		200	{array} model.Account
func (p *Controller) UserProfileHandler(c *gin.Context) {
	account.WalletAccount = c.Query("walletAccount")
	if account.WalletAccount == " " {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Empty walletAccount"})
	} else {
		result := p.md.GetProfile(account)
		c.JSON(http.StatusOK, result)
		// c.JSON(http.StatusOK, gin.H{"Success": result})

		if len(result) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "You must be a registered user."})
		}
	}
	log.Println(account)
}
