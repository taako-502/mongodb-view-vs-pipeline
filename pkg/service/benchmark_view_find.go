package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// BenchmarkViewFind ビューを利用した検索のベンチマークを実行する
func (s *service) BenchmarkViewFind(db *mongo.Database) (time.Duration, error) {
	ctx := context.TODO()
	collection := db.Collection(s.viewName)

	start := time.Now()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to find from view: %v", err)
	}
	defer cursor.Close(ctx)

	count := 0
	for cursor.Next(ctx) {
		count++
	}

	elapsed := time.Since(start)
	return elapsed, nil
}
