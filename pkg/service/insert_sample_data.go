package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// InsertSampleData 検証に利用するサンプルデータを挿入する
func (s *service) InsertSampleData(num int64) error {
	ctx := context.TODO()
	count, _ := s.collection.CountDocuments(ctx, bson.M{})
	if count >= num {
		fmt.Println("Sample data already exists, skipping insertion.")
		return fmt.Errorf("sample data already exists")
	}

	docs := make([]interface{}, num)
	for i := range num {
		docs[i] = bson.M{
			"name":  fmt.Sprintf("User%d", i),
			"score": i % 100,
		}
	}

	_, err := s.collection.InsertMany(ctx, docs)
	if err != nil {
		return fmt.Errorf("failed to insert sample data: %v", err)
	}

	return nil
}
