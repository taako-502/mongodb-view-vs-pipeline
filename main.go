package main

import (
	"context"
	"fmt"
	"log"

	"github.com/taako-502/mongodb-embedded-array-size-impact/pkg/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	mongoURI       = "mongodb://localhost:27017"
	databaseName   = "testdb"
	collectionName = "testcollection"
	viewName       = "testview"
)

// go run main.go
func main() {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(databaseName)
	collection := db.Collection(collectionName)

	s := service.NewService(collection, viewName, collectionName)
	if err := s.CreateView(db); err != nil {
		log.Fatalf("Failed to create view: %v", err)
	}
	if err := s.DropSampleData(); err != nil {
		log.Fatalf("Failed to drop sample data: %v", err)
	}

	// データ量の増減による影響を検証
	// 10,000件, 100,000件, 1,000,000件, 10,000,000件
	test := []int64{10000, 100000, 1000000, 10000000}
	for _, num := range test {
		if err := s.InsertSampleData(num); err != nil {
			log.Fatalf("Failed to insert sample data: %v", err)
		}
		viewTime, err := s.BenchmarkViewFind(db)
		if err != nil {
			log.Fatalf("Failed to drop sample data: %v", err)
		}
		aggTime, err := s.BenchmarkAggregationFind()
		if err != nil {
			log.Fatalf("Failed to drop sample data: %v", err)
		}
		fmt.Println("--------------------------------------------------")
		fmt.Printf("Number of documents: %d\n", num)
		fmt.Printf("View find time: %v\n", viewTime)
		fmt.Printf("Aggregation find time: %v\n", aggTime)
		if err := s.DropSampleData(); err != nil {
			log.Fatalf("Failed to drop sample data: %v", err)
		}
	}
}
