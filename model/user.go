package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	checkUser bson.M
	datas     []bson.M
)

func (m *Model) RegisterUser(account Account) error {
	filter := bson.D{{Key: "id", Value: account.WalletAccount}}
	update := bson.D{{Key: "$set", Value: account}}

	opts := options.Update().SetUpsert(true)
	result, err := m.colAccount.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println("Failed to insert data in user_account")

		return fmt.Errorf("fail to register user: %w", err)
	}
	fmt.Println("result", result.MatchedCount, result.UpsertedCount)

	return nil
}

func (m *Model) UpdateUser(account Account) error {
	filter := bson.D{{Key: "id", Value: account.WalletAccount}}
	update := bson.D{{Key: "$set", Value: account}}

	opts := options.Update().SetUpsert(true)
	if _, err := m.colAccount.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		log.Println("Failed to update data in user_account")
		return fmt.Errorf("fail to update user: %w", err)
	}
	return nil
}

func (m *Model) GetProfile(account Account) map[string]interface{} {
	var profile []map[string]interface{}

	fmt.Println("account", account.WalletAccount)
	filter := bson.M{"walletaccount": account.WalletAccount}
	res, err := m.colAccount.Find(context.TODO(), filter)
	if err != nil {
		log.Print(err)
	}
	// 결과를 변수에 담기
	if err = res.All(context.TODO(), &datas); err != nil {
		fmt.Println(err)
	}
	fmt.Println("datas", datas)

	dnftFilter := bson.D{{"walletAccount", account.WalletAccount}}
	tripFilter := bson.D{{"walletaccount", account.WalletAccount}}
	dnftCount, err := m.colDnftBadge.CountDocuments(context.TODO(), dnftFilter)
	tripCount, err := m.colTripPlan.CountDocuments(context.TODO(), tripFilter)

	for _, data := range datas {
		item := make(map[string]interface{})

		item["commentcount"] = data["commentcount"]
		item["likecount"] = data["likecount"]
		item["nickname"] = data["nickname"]
		item["mydnftcount"] = dnftCount
		item["mytravelcount"] = tripCount
		profile = append(profile, item)
	}

	return profile[0]
}

func (m *Model) MatchUser(account string) bool {
	result := false
	fmt.Println("match", account)
	// 메뉴 조회
	err := m.colAccount.FindOne(context.TODO(), bson.M{"walletaccount": account}).Decode(&checkUser)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}
	fmt.Println("checkUser", checkUser)
	fmt.Println("account", account)
	fmt.Println("checkUser", checkUser["walletaccount"])
	if account == checkUser["walletaccount"] {
		result = true
	} else {
		result = false
	}
	return result
}
