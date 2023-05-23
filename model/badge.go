package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	fmt.Println("estCount", estCount)
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

	fmt.Println("check encyDnft", encyDnft, result)

	return encyDnft
}

// MyDnft 여러 개 불러오기
func (m *Model) GetMyAllDnft(account string) []bson.M {
	// Aggregation 파이프라인 생성
	pipeline := mongo.Pipeline{
		// Group 스테이지: DnftCountry로 그룹화하고 최대 issueCount 계산
		{
			{"$group", bson.D{
				{"_id", "$DnftCountry"},
				{"maxIssueCount", bson.D{{"$max", "$IssueCount"}}},
			}},
		},
		// Match 스테이지: 최대 issueCount가 0보다 큰 문서만 필터링
		{
			{"$match", bson.D{{"maxIssueCount", bson.D{{"$gt", 0}}}}},
		},
		// Project 스테이지: 필요한 필드만 선택하여 결과 형식 조정
		{
			{"$project", bson.D{
				{"DnftCountry", "$_id"},
				{"maxIssueCount", 1},
				{"_id", 0},
			}},
		},
	}

	// Aggregation 실행
	cursor, err := m.colDnftBadge.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	fmt.Println("cursor", cursor)

	// 결과 처리
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	fmt.Println("results", results)

	// 결과 출력
	for _, result := range results {
		dnftCountry := result["DnftCountry"].(string)
		maxIssueCount := result["maxIssueCount"].(int64)
		fmt.Printf("DnftCountry: %s, Max IssueCount: %d\n", dnftCountry, maxIssueCount)
	}

	return results
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
