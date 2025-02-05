package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *service) DropSampleData() error {
	ctx := context.TODO()
	count, _ := s.collection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		fmt.Println("Sample data does not exist, skipping deletion.")
		return nil
	}

	if _, err := s.collection.DeleteMany(ctx, bson.M{}); err != nil {
		return fmt.Errorf("failed to drop sample data: %v", err)
	}

	return nil
}
