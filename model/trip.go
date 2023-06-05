package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"trypto-server/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	tripPlan  TripPlan
	tripPlans []TripPlan
	encyDnft  EncyclopediaDNFT
)

func (m *Model) InsertTripPlan(tripPlan *TripPlan) *TripPlan {

	totalTripPlanCount, err := m.colTripPlan.EstimatedDocumentCount(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("insert tripplan", tripPlan)
	fmt.Println("totalTripPlanCount", totalTripPlanCount)
	tripPlan.TripId = totalTripPlanCount + 1
	filter := bson.M{"walletaccount": tripPlan.WalletAccount}
	res, _ := m.colAccount.Find(context.TODO(), filter)

	res.All(context.TODO(), &datas)
	tripPlan.NickName = datas[0]["nickname"].(string)
	result, err := m.colTripPlan.InsertOne(context.TODO(), tripPlan)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tripPlan", tripPlan)
	fmt.Println("datas", datas)

	account.AccountID = tripPlan.TripId
	account.WalletAccount = tripPlan.WalletAccount
	account.NickName = tripPlan.NickName
	fmt.Println("account", account)
	accFilter := bson.D{{Key: "walletaccount", Value: account.WalletAccount}}
	accUpdate := bson.D{{Key: "$set", Value: account}}
	accOpts := options.Update().SetUpsert(true)
	_, acc_err := m.colAccount.UpdateOne(context.TODO(), accFilter, accUpdate, accOpts)
	if err != nil {
		panic(acc_err)
	}

	fmt.Println("check encyDnft", tripPlan, result)

	return tripPlan
}

func (m *Model) PatchTripPlan(tripPlan *TripPlan) *TripPlan {

	tripFilter := bson.D{{Key: "tripid", Value: tripPlan.TripId}}
	tripUpdate := bson.D{{Key: "$set", Value: tripPlan}}

	tripOpts := options.Update().SetUpsert(true)
	_, acc_err := m.colTripPlan.UpdateOne(context.TODO(), tripFilter, tripUpdate, tripOpts)
	if acc_err != nil {
		panic(acc_err)
	}

	filter := bson.M{"tripid": tripPlan.TripId} // 데이터를 담을 변수 선언
	// 메뉴 조회
	res, err := m.colTripPlan.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	// 결과를 변수에 담기
	if err = res.All(context.TODO(), &tripPlan); err != nil {
		fmt.Println(err)
	}

	return tripPlan
}

// 여행계획 전부 가져오기
func (m *Model) SelectAllTrip() []bson.M {
	var datas []bson.M
	res, err := m.colTripPlan.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.Error(err)
	}
	fmt.Println("res", res)

	// 결과를 변수에 담기
	if err = res.All(context.TODO(), &datas); err != nil {
		fmt.Println(err)
	}
	fmt.Println("datas", datas)
	return datas
}

// 나의 여행계획 가져오기
func (m *Model) SelectMyTrip(account string) []TripPlan {
	fmt.Println("account", account)

	filter := bson.M{"walletaccount": account} // 데이터를 담을 변수 선언
	// 메뉴 조회
	res, err := m.colTripPlan.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	// 결과를 변수에 담기
	if err = res.All(context.TODO(), &tripPlans); err != nil {
		fmt.Println(err)
	}
	fmt.Println("datas", tripPlans)
	return tripPlans
}

// 디테일한 여행계획 가져오기
func (m *Model) SelectDetailTrip(tripId int64) []TripPlan {

	filter := bson.M{"tripid": tripId} // 데이터를 담을 변수 선언
	// 메뉴 조회
	res, err := m.colTripPlan.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	// 결과를 변수에 담기
	if err = res.All(context.TODO(), &tripPlans); err != nil {
		fmt.Println(err)
	}
	fmt.Println("tripPlans", tripPlans)
	return tripPlans
}

// 여행계획 검색하기
func (m *Model) SearchTrip(Query string) []TripPlan {
	fmt.Println(Query)

	matchStage := bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", Query}}}}}}

	//filter := bson.D{{"$text", bson.D{{"$search", Query}}}}
	fmt.Println("matchStage", matchStage)
	cursor, err := m.colTripPlan.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
	if err != nil {
		panic(err)
	}
	fmt.Println("cursor", cursor)
	var results []TripPlan
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}
	fmt.Println("results", results)
	return results
}
