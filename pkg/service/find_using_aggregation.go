package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// FindUsingAggregation 集計を利用した検索のベンチマークを実行する
func (s *service) FindUsingAggregation() error {
	ctx := context.TODO()
	cursor, err := s.collection.Aggregate(ctx, mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"score": bson.M{"$gte": 50}}}},
	})
	if err != nil {
		return fmt.Errorf("failed to create aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	count := 0
	for cursor.Next(ctx) {
		count++
	}

	return nil
}
