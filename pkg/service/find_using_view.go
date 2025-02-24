package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// FindUsingView ビューを利用した検索を実行する
func (s *service) FindUsingView(db *mongo.Database) error {
	ctx := context.TODO()
	collection := db.Collection(s.viewName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to find from view: %w", err)
	}
	defer cursor.Close(ctx)

	count := 0
	for cursor.Next(ctx) {
		count++
	}

	return nil
}
