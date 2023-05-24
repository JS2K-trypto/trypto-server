package model

import (
	"context"
	"fmt"
	"log"
	"trypto-server/logger"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	tripPlan TripPlan
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

// MyDnft 여러 개 불러오기
func (m *Model) GetAllTripPlan(account string) []bson.M {
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
