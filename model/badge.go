package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	checkCountry bson.M
)

func (m *Model) MatchBadgeResource(encyDnft *EncyclopediaDNFT) *EncyclopediaDNFT {

	err := m.colResource.FindOne(context.TODO(), bson.M{"Country": encyDnft.DnftCountry}).Decode(&checkCountry)
	fmt.Println("checkCountry", checkCountry)
	if err != nil {
		log.Println(err)
		fmt.Errorf("fail to get menu detail")
	}

	// filter := bson.D{{Key: "id", Value: encyDnft.DnftId}}
	// update := bson.D{{Key: "$set", Value: encyDnft}}

	// opts := options.Update().SetUpsert(true)
	// if _, err := m.colDnftBadge.UpdateOne(context.TODO(), filter, update, opts); err != nil {
	// 	log.Println(err)
	// }
	fmt.Println(checkCountry["Country"].(string))
	fmt.Println(encyDnft.DnftCountry)
	encyDnft.DnftImgUrl = checkCountry["bronze"].(string)

	// if checkCountry["Country"].(string) == encyDnft.DnftCountry {
	// 	encyDnft.DnftImgUrl = checkCountry["BadgeUrl_bronze"].(string)
	// }
	fmt.Println("match", encyDnft)
	return encyDnft
}
