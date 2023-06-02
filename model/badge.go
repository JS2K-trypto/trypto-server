package model

import (
	"context"
	"fmt"
	"log"
	conf "trypto-server/config"
	"trypto-server/logger"

	"github.com/thirdweb-dev/go-sdk/thirdweb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	checkCountry bson.M
	encytDnft    EncyclopediaDNFT
	encytDnfts   []EncyclopediaDNFT
	account      Account
	num          int64
)

func increase(num int64) int64 {
	return num + 1
}

func upgrade(count int64) interface{} {
	increaseId := int(count)
	config2 := conf.GetConfig("./config/config.toml")
	contractAddress := config2.Contract.DnftContract
	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
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
	log.Println("contract", contract)
	increase, err := contract.Call(context.Background(), "increasebadgeLevel", increaseId)
	log.Println("increase", increase)
	return increase
}

func (m *Model) CreateDNFTBadge(encyDnft *EncyclopediaDNFT) *EncyclopediaDNFT {

	fmt.Println("start encyDnft", encyDnft)
	bronzeUp := 2
	silverUp := 3
	goldUp := 4

	err := m.colResource.FindOne(context.TODO(), bson.M{"Country": encyDnft.DnftCountry}).Decode(&checkCountry)

	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}

	filter := bson.D{{"walletaccount", encyDnft.WalletAccount}, {"dnftcountry", encyDnft.DnftCountry}}
	count, err := m.colDnftBadge.CountDocuments(context.TODO(), filter)

	estCount, estCountErr := m.colDnftBadge.EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		panic(estCountErr)
	}

	//wallet account가발급한 dnft 개수가 3개 미만이면 브론즈 티어, 브론즈 이미지URL은 브론즈 URL
	//wallet account가발급한 dnft 개수가 5개 미만이면 실버
	//wallet account가발급한 dnft 개수가 7개 미만이면 골드

	encyDnft.DnftBronzeUrl = checkCountry["bronze"].(string)
	encyDnft.DnftSilverUrl = checkCountry["silver"].(string)
	encyDnft.DnftGoldUrl = checkCountry["gold"].(string)
	encyDnft.DnftId = increase(estCount)
	encyDnft.IssueCount = increase(count)
	account.MyDNFTCount = encyDnft.IssueCount

	res, _ := m.colAccount.Find(context.TODO(), filter)
	res.All(context.TODO(), &datas)

	accFilter := bson.D{{Key: "walletaccount", Value: encyDnft.WalletAccount}}
	accUpdate := bson.D{{Key: "$set", Value: account.MyDNFTCount}}
	accOpts := options.Update().SetUpsert(true)
	_, acc_err := m.colAccount.UpdateOne(context.TODO(), accFilter, accUpdate, accOpts)
	if err != nil {
		panic(acc_err)
	}

	if count <= int64(bronzeUp) {
		encyDnft.BadgeTier = "bronze"
		encyDnft.DnftImgUrl = checkCountry["bronze"].(string)
	} else if count >= int64(silverUp) && count < int64(goldUp) {
		encyDnft.BadgeTier = "silver"
		encyDnft.DnftImgUrl = checkCountry["silver"].(string)
		upgrade(count)
	} else if count >= int64(goldUp) {
		encyDnft.BadgeTier = "gold"
		encyDnft.DnftImgUrl = checkCountry["gold"].(string)
		upgrade(count)
		upgrade(count)
	}

	result, err := m.colDnftBadge.InsertOne(context.TODO(), encyDnft)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("check encyDnft", encyDnft, result)
	return encyDnft
}

// Dnft 뱃지 여러 개 불러오기
func (m *Model) GetAllDnft(account string) []bson.M {
	var datas []bson.M
	res, err := m.colDnftBadge.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.Error(err)
	}

	// 결과를 변수에 담기
	if err = res.All(context.TODO(), &datas); err != nil {
		fmt.Println(err)
	}
	return datas
}

// 나의 Dnft 불러오기
func (m *Model) GetMyDnft(account string) []EncyclopediaDNFT {
	var maxIssueCount int64

	filter := bson.M{"walletaccount": account}
	fmt.Println("account", account)
	res, err := m.colDnftBadge.Find(context.TODO(), filter)
	for res.Next(context.Background()) {

		if err := res.Decode(&encytDnft); err != nil {
			log.Fatal(err)
		}

		if encytDnft.IssueCount > maxIssueCount {
			maxIssueCount = encytDnft.IssueCount
			encytDnfts = []EncyclopediaDNFT{encytDnft}
		} else if encytDnft.IssueCount == maxIssueCount {
			encytDnfts = append(encytDnfts, encytDnft)
		}

	}

	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}
	return encytDnfts

}
