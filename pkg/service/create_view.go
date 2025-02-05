package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *service) CreateView(db *mongo.Database) {
	fmt.Println("Creating MongoDB View...")

	_ = db.Collection(s.viewName).Drop(context.TODO())

	// ビューの作成
	command := bson.D{
		{Key: "create", Value: s.viewName},
		{Key: "viewOn", Value: s.collectionName},
		{Key: "pipeline", Value: bson.A{
			bson.D{{Key: "$match", Value: bson.M{"score": bson.M{"$gte": 50}}}}, // スコアが50以上のデータを取得
		}},
	}

	err := db.RunCommand(context.TODO(), command).Err()
	if err != nil {
		log.Fatalf("Failed to create view: %v", err)
	}

	fmt.Println("View created successfully.")
}
