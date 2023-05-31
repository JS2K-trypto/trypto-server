package controller

import (
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
//	@Summary		계정주소, 닉네임을 입력합니다.
//	@Tags			UserRegisterHandler(닉네임 등록/수정)
//	@Description	유저 닉네임을 등록 및 수정 해주는 함수로 지갑계정으로 연결 후 사용자가 닉네임을 입력할 수 있다. 이후 닉네임 수정은 자유롭게 가능하다.
//	@name			UserRegisterHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount 	path	string	true	"walletAccount",
//	@Param			nickName		path	string	true	"nickName"
//	@Router			/v01/acc/register [post]
//	@Success		200	{object}	string
func (p *Controller) UserRegisterHandler(c *gin.Context) {

	account.WalletAccount = c.Query("walletAccount")
	account.NickName = c.Query("nickName")

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
//	@Summary		계정주소를 입력합니다.
//	@Tags			UserProfileHandler(나의 프로필 정보 가져오기)
//	@Description	유저 프로필 정보를 가져는 함수다. 다음과 같은 정보를 가져온다. [닉네임, 나의 여행계획 카운트, 나의 Dynamic NFT 카운트, 좋아요 카운트 , 댓글 카운트]
//	@name			UserProfileHandler
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string	true	"walletAccount",
//	@Param			nickName		path	string	true	"nickName"
//	@Router			/v01/acc/profile [get]
//	@Success		200	{object}	string
func (p *Controller) UserProfileHandler(c *gin.Context) {

	account.WalletAccount = c.Query("walletAccount")
	if account.WalletAccount == " " {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Empty walletAccount"})
	} else {
		result := p.md.GetProfile(account)
		c.JSON(http.StatusOK, result)
		//c.JSON(http.StatusOK, gin.H{"Success": result})

		if len(result) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "You must be a registered user."})
		}
	}
	log.Println(account)

}
