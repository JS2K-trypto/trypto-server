package model

import (
	"context"
	conf "trypto-server/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client       *mongo.Client
	colAccount   *mongo.Collection
	colTripPlan  *mongo.Collection
	colDnftBadge *mongo.Collection
	colResource  *mongo.Collection
}

func NewModel() (*Model, error) {

	config := conf.GetConfig("./config/config.toml")
	r := &Model{}
	var err error
	mgUrl := config.DB["user"]["host"].(string)
	dbName := config.DB["user"]["name"].(string)
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(dbName)
		r.colAccount = db.Collection("account")
		r.colTripPlan = db.Collection("tripPlan")
		r.colDnftBadge = db.Collection("dnftbadge")
		r.colResource = db.Collection("resource")

		// indexmodel := mongo.IndexModel{Keys: bson.D{{"triptitle", "text"}, {}}}
		// name, err := r.colTripPlan.Indexes().CreateOne(context.TODO(), indexmodel)
		// if err != nil {
		// 	panic(err)
		// }
		// log.Println(name)
	}

	return r, nil
}
