package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// InsertSampleData 検証に利用するサンプルデータを挿入する
func (s *service) InsertSampleData(collection *mongo.Collection) {
	ctx := context.TODO()
	count, _ := collection.CountDocuments(ctx, bson.M{})
	if count >= s.numDocuments {
		fmt.Println("Sample data already exists, skipping insertion.")
		return
	}

	fmt.Println("Inserting sample data...")
	docs := make([]interface{}, s.numDocuments)
	for i := range s.numDocuments {
		docs[i] = bson.M{
			"name":  fmt.Sprintf("User%d", i),
			"score": i % 100,
		}
	}

	_, err := collection.InsertMany(ctx, docs)
	if err != nil {
		log.Fatalf("Failed to insert sample data: %v", err)
	}
	fmt.Println("Sample data inserted successfully.")
}
