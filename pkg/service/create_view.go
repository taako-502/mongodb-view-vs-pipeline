package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateView 検証に利用するビューを作成する
func (s *service) CreateView(db *mongo.Database) error {
	fmt.Println("Creating MongoDB View...")
	_ = db.Collection(s.viewName).Drop(context.TODO())
	command := bson.D{
		{Key: "create", Value: s.viewName},
		{Key: "viewOn", Value: s.collectionName},
		{Key: "pipeline", Value: bson.A{
			bson.D{{Key: "$match", Value: bson.M{"score": bson.M{"$gte": 50}}}}, // スコアが50以上のデータを取得
		}},
	}

	if err := db.RunCommand(context.TODO(), command).Err(); err != nil {
		return fmt.Errorf("failed to create view: %v", err)
	}

	fmt.Println("View created successfully.")
	return nil
}
