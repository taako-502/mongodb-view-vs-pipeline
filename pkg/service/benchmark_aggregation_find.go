package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// BenchmarkAggregationFind 集計を利用した検索のベンチマークを実行する
func (s *service) BenchmarkAggregationFind() (time.Duration, error) {
	ctx := context.TODO()

	start := time.Now()
	cursor, err := s.collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"score": bson.M{"$gte": 50}}}},
	})
	if err != nil {
		log.Fatalf("Failed to execute aggregation: %v", err)
	}
	defer cursor.Close(ctx)

	count := 0
	for cursor.Next(ctx) {
		count++
	}

	elapsed := time.Since(start)
	return elapsed, nil
}
