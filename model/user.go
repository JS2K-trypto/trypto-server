package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Model) RegisterUser(account Account) error {
	filter := bson.D{{Key: "id", Value: account.WalletAccount}}
	update := bson.D{{Key: "$set", Value: account}}

	opts := options.Update().SetUpsert(true)
	if _, err := m.colAccount.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		log.Println("Failed to insert data in user_account")
		return fmt.Errorf("fail to register user: %w", err)
	}
	return nil
}

func (m *Model) GetProfile(account Account) error {

	// 메뉴 조회
	err := m.colAccount.FindOne(context.TODO(), bson.M{"walletAccount": account.WalletAccount}).Decode(&account)
	fmt.Println("account.WalletAccount", account.WalletAccount)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")

	}
	return &account
}
