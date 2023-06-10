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
	empty    []string
)

// CreateBadge godoc
//
//	@BasePath	/v01
//	@Schemes
//	@Summary		Enter your mumbai wallet account, latitude, and longitude.
//	@Tags			CreateBadge (issuing a badge)
//	@Description	It takes the latitude, longitude, and references the country's resources (IPFS URI, NFT METADATA) to issue the badge.
//	@name			CreateBadge
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount",
//	@Param			latitude		path	string 	 	true	"latitude",
//	@Param			longitude		path	string	    true	"longitude"
//	@Router			/v01/badge/issue [post]
//	@Success		200	 {array} model.EncyclopediaDNFT
func (p *Controller) CreateBadge(c *gin.Context) {
	if err := c.ShouldBindJSON(&encyDnft); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(&encyDnft)
	now := time.Now()
	custom := now.Format("2006-01-02 15:04:05")
	geocoder := openstreetmap.Geocoder()
	location, err := geocoder.ReverseGeocode(encyDnft.Latitude, encyDnft.Longitude)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("나라", location.Country)
	fmt.Println("location", location)
	encyDnft.DnftCountry = location.Country
	encyDnft.DnftTime = custom
	result := p.md.CreateDNFTBadge(&encyDnft)

	config2 := conf.GetConfig("./config/config.toml")
	contractAddress := config2.Contract.DnftContract
	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		PrivateKey: config2.Contract.PRIVATEKEY,
	})
	if err != nil {
		panic(err)
	}
	contract, err := sdk.GetContractFromAbi(contractAddress, model.ABI)
	if err != nil {
		panic(err)
	}

	mint, err := contract.Call(context.Background(), "safeMint", encyDnft.WalletAccount, encyDnft.DnftImgUrl)
	log.Println("mint", mint)
	increaseId := int(result.DnftId)
	increase, err := contract.Call(context.Background(), "increasebadgeLevel", increaseId)
	log.Println("increase", increase)

	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	// tokeURI = 2 > 3  리믹스
	// tokenid > 디비 2
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty walletAccount"})
	} else {
		fmt.Println("mint", mint)
		c.JSON(http.StatusOK, mint)

	}
}

// GetMyBadge godoc
//
//	@BasePath		/v01
//	@Summary		Entering your My Wallet account will bring up your most issued badges.
//	@Tags			GetMyBadge(Get my dynamic nft badges)
//	@Description	Function to issue badges based on user location
//	@name			GetMyBadge
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount"
//	@Router			/v01/badge/user [get]
//	@Success		200	{array} model.EncyclopediaDNFT
func (p *Controller) GetMyBadge(c *gin.Context) {
	account.WalletAccount = c.Query("walletAccount")
	empty = []string{}
	fmt.Println("account", account.WalletAccount)
	result := p.md.GetMyDnft(account.WalletAccount)

	if len(result) > 0 {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty.": empty})
	}
}

// GetMyBadgeAll godoc
//
//	@BasePath		/v01
//	@Summary
//	@Tags			GetMyBadgeAll(get all my dynamic nft badges)
//	@Description	Get all my DYNAMIC NFT badges
//	@name			GetMyBadgeAll
//	@Accept			json
//	@Produce		json
//	@Param			walletAccount	path	string 	 	true	"walletAccount"
//	@Router			/v01/badge/user [get]
//	@Success		200	{array} model.EncyclopediaDNFT
func (p *Controller) GetMyBadgeAll(c *gin.Context) {
	account.WalletAccount = c.Query("walletAccount")

	fmt.Println("account", account.WalletAccount)
	result := p.md.GetMyDnftAll(account.WalletAccount)
	empty = []string{}
	if len(result) > 0 {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"The result is empty.": empty})
	}
}
