package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"trypto-server/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	tripPlan  TripPlan
	tripPlans TripPlans
)

func (m *Model) InsertTripPlan(tripPlan *TripPlan) *TripPlan {

	result, err := m.colTripPlan.InsertOne(context.TODO(), tripPlan)
	if err != nil {
		log.Fatal(err)
	}
	//이전 Dnft 발급한 것을 삭제하기

	//한달에 한 번만 발급할 수 있게 조정하기

	fmt.Println("check encyDnft", tripPlan, result)

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
func (m *Model) SelectMyTrip(account string) *TripPlans {

	fmt.Println("account", account)

	filter := bson.M{"walletaccount": account} // 데이터를 담을 변수 선언
	// 메뉴 조회
	res, err := m.colTripPlan.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next(context.Background()) {

		if err := res.Decode(&tripPlan); err != nil {
			log.Fatal(err)
		}
		tripPlans.Arr = append(tripPlans.Arr, tripPlan)
	}

	return &tripPlans
}

// 나의 여행계획 가져오기
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
