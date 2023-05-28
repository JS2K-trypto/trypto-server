package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	conf "trypto-server/config"
	_ "trypto-server/docs"
	"trypto-server/model"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var (
	encyDnft model.EncyclopediaDNFT
	location model.Location
)

// 뱃지를 생성해주는 함수
// CreateBadge godoc
// @Summary 위도, 경도를 입력받고 해당하는 나라의 리소스를 참고해서 뱃지를 발급해줍니다.
// @Tags CreateBadge
// @Description 사용자 위치를 참고해서 뱃지를 발급하는 함수
// @name CreateBadge
// @Accept  json
// @Produce  json
// @Param walletAccount  string true "walletAccount"
// @Param latitude float true "latitude"
// @Param longitude float true "longitude"
// @Router v01/badge/issue [post]
// @Success 200 {object} string
func (p *Controller) CreateBadge(c *gin.Context) {

	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("encyDnft", encyDnft)
	geocoder := openstreetmap.Geocoder()
	location, err := geocoder.ReverseGeocode(encyDnft.Latitude, encyDnft.Longitude)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	encyDnft.DnftCountry = location.Country
	encyDnft.DnftTime = custom

	result := p.md.MatchBadgeResource(&encyDnft)
	log.Println("dnft", result)

	config2 := conf.GetConfig("./config/.config.toml")
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

// 뱃지를 가져오는 함수
// CreateBadge godoc
// @Summary 나의 뱃지를 가져오는 함수
// @Tags GetMyBadge
// @Description 사용자 위치를 참고해서 뱃지를 발급하는 함수
// @name GetMyBadge
// @Accept  json
// @Produce  json
// @Param walletAccount  string true "walletAccount"
// @Router v01/badge/user [get]
// @Success 200 {object} string
func (p *Controller) GetMyBadge(c *gin.Context) {
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("account", account.WalletAccount)
	result := p.md.GetMyAllDnft(account.WalletAccount)
	c.JSON(http.StatusOK, result)
}
