package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	conf "trypto-server/config"
	"trypto-server/model"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"

	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var (
	encyDnft model.EncyclopediaDNFT
	location model.Location
)

// CreateBadge godoc
//
//	@BasePath	/v01
//	@Schemes
//	@Summary		지갑계정, 위도, 경도를 입력해줍니다.
//	@Tags			CreateBadge(뱃지 발급)
//	@Description	위도, 경도를 입력받고 해당하는 나라의 리소스(ipfs uri, nft metadata)를 참고해서 뱃지를 발급해줍니다.
//	@name			CreateBadge
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount",
//	@Param			latitude		path	string 	 	true	"latitude",
//	@Param			longitude		path	string	    true	"longitude"
//	@Router			/v01/badge/issue [post]
//	@Success		200	{object}	string
func (p *Controller) CreateBadge(c *gin.Context) {

	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	geocoder := openstreetmap.Geocoder()
	location, err := geocoder.ReverseGeocode(encyDnft.Latitude, encyDnft.Longitude)
	if err != nil {
		fmt.Println("Error:", err)
	}

	encyDnft.DnftCountry = location.Country
	encyDnft.DnftTime = custom
	result := p.md.CreateDNFTBadge(&encyDnft)
	log.Println("dnft", result)

	config2 := conf.GetConfig("./config/config.toml")
	contractAddress := config2.Contract.DnftContract
	sdk, err := thirdweb.NewThirdwebSDK("goerli", &thirdweb.SDKOptions{
		PrivateKey: config2.Contract.PRIVATEKEY,
	})
	if err != nil {
		panic(err)
	}
	log.Println("contractAddress", contractAddress)
	contract, err := sdk.GetContractFromAbi(contractAddress, ABI)
	if err != nil {
		panic(err)
	}

	balance, err := contract.Call(context.Background(), "balanceOf", encyDnft.WalletAccount)
	if err != nil {
		panic(err)
	}

	log.Println("balance", balance)
	metaData := []string{}
	metaData = append(metaData, encyDnft.DnftBronzeUrl)
	metaData = append(metaData, encyDnft.DnftSilverUrl)
	metaData = append(metaData, encyDnft.DnftGoldUrl)

	mint, err := contract.Call(context.Background(), "safeMint", encyDnft.WalletAccount, metaData[0])
	fmt.Println("mint", mint)

	c.JSON(http.StatusOK, mint)

}

// GetMyBadge godoc
//
//	@BasePath		/v01
//	@Summary		지갑계정을 입력해줍니다.
//	@Tags			GetMyBadge(나의 뱃지 가져오기)
//	@Description	사용자 위치를 참고해서 뱃지를 발급하는 함수
//	@name			GetMyBadge
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount"
//	@Router			/v01/badge/user [get]
//	@Success		200	{object}	string
func (p *Controller) GetMyBadge(c *gin.Context) {
	account.WalletAccount = c.Query("walletAccount")

	fmt.Println("account", account.WalletAccount)
	result := p.md.GetMyDnft(account.WalletAccount)
	c.JSON(http.StatusOK, result)
}
