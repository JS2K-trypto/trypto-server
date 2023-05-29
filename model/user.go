package model

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	checkUser bson.M
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

func (m *Model) GetProfile(account Account) []string {

	// 메뉴 조회
	err := m.colAccount.FindOne(context.TODO(), bson.M{"walletAccount": account.WalletAccount}).Decode(&account)
	fmt.Println("account.WalletAccount", account.WalletAccount)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}
	fmt.Println("account", account)
	fmt.Println("account", account.NickName)
	profile := []string{}
	// = append(profile, account.ID.String())
	profile = append(profile, account.NickName)
	profile = append(profile, strconv.Itoa(account.MyTravelCount))
	profile = append(profile, strconv.Itoa(account.MyDNFTCount))
	profile = append(profile, strconv.Itoa(account.LikeCount))
	profile = append(profile, strconv.Itoa(account.CommentCount))

	return profile
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
