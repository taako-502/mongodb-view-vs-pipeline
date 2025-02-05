package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// BenchmarkAggregationFind 集計を利用した検索のベンチマークを実行する
func (s *service) BenchmarkAggregationFind(collection *mongo.Collection) time.Duration {
	ctx := context.TODO()

	start := time.Now()
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
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
	fmt.Printf("Aggregation Find fetched %d documents.\n", count)
	return elapsed
}
