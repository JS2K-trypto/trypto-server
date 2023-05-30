package model

import (
	"context"
	"fmt"
	"log"
	"trypto-server/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	checkCountry bson.M
	encytDnft    EncyclopediaDNFT
	encytDnfts   EncyclopediaDNFTs
	account      Account
)

func (m *Model) MatchBadgeResource(encyDnft *EncyclopediaDNFT) *EncyclopediaDNFT {

	err := m.colResource.FindOne(context.TODO(), bson.M{"Country": encyDnft.DnftCountry}).Decode(&checkCountry)

	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}

	filter := bson.D{{"walletAccount", encyDnft.WalletAccount}, {"dnftCountry", encyDnft.DnftCountry}}
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
	encyDnft.DnftId = estCount + 1
	encyDnft.IssueCount = count + 1
	
	if count < 5 {
		encyDnft.BadgeTier = "bronze"
		encyDnft.DnftImgUrl = checkCountry["bronze"].(string)
	} else if count < 10 && count >= 5 {
		encyDnft.BadgeTier = "silver"
		encyDnft.DnftImgUrl = checkCountry["silver"].(string)
	} else if count >= 10 {
		encyDnft.BadgeTier = "gold"
		encyDnft.DnftImgUrl = checkCountry["gold"].(string)
	}

	result, err := m.colDnftBadge.InsertOne(context.TODO(), encyDnft)
	if err != nil {
		log.Fatal(err)
	}
	account.MyDNFTCount = encyDnft.IssueCount
	accFilter := bson.D{{Key: "id", Value: account.WalletAccount}}
	accUpdate := bson.D{{Key: "$set", Value: account.MyDNFTCount}}

	accOpts := options.Update().SetUpsert(true)
	_, acc_err := m.colAccount.UpdateOne(context.TODO(), accFilter, accUpdate, accOpts)
	if err != nil {
		panic(acc_err)
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

// MyDnft 한개만 불러오기
func (m *Model) GetMyDnft(account string) *EncyclopediaDNFTs {
	filter := bson.M{"walletAccount": account} // 데이터를 담을 변수 선언

	res, err := m.colDnftBadge.Find(context.TODO(), filter)
	for res.Next(context.Background()) {

		if err := res.Decode(&encytDnft); err != nil {
			log.Fatal(err)
		}
		encytDnfts.Arr = append(encytDnfts.Arr, encytDnft)
	}

	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}
	return &encytDnfts

}
