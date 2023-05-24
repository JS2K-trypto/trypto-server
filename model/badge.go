package model

import (
	"context"
	"fmt"
	"log"
	"trypto-server/logger"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	checkCountry bson.M
	encytDnft    EncyclopediaDNFT
)

func (m *Model) MatchBadgeResource(encyDnft *EncyclopediaDNFT) *EncyclopediaDNFT {

	err := m.colResource.FindOne(context.TODO(), bson.M{"Country": encyDnft.DnftCountry}).Decode(&checkCountry)
	log.Println("checkCountry", checkCountry)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}

	// 컬렉션의 도큐먼트 카운트 세기
	///filter := bson.D{{"walletAccount": encyDnft.WalletAccount}}
	// count, err := m.colDnftBadge.CountDocuments(context.TODO(), bson.M{"walletAccount": encyDnft.WalletAccount})
	// fmt.Println("count", count)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	filter := bson.D{{"walletAccount", encyDnft.WalletAccount}, {"dnftCountry", encyDnft.DnftCountry}}
	count, err := m.colDnftBadge.CountDocuments(context.TODO(), filter)
	fmt.Println("count", count)

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
	//이전 Dnft 발급한 것을 삭제하기

	//한달에 한 번만 발급할 수 있게 조정하기

	fmt.Println("check encyDnft", encyDnft, result)

	return encyDnft
}

// MyDnft 여러 개 불러오기
func (m *Model) GetMyAllDnft(account string) []bson.M {
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
func (m *Model) GetMyDnft(account string) *EncyclopediaDNFT {

	// return list
	//issue count값이 가장 크고 country별로 그룹화해서 보여줘야함
	err := m.colDnftBadge.FindOne(context.TODO(), bson.M{"walletAccount": account}).Decode(&encytDnft)
	fmt.Println("account.WalletAccount", account)
	fmt.Println(encytDnft)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")

	}
	return &encytDnft

}
