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
	numDocuments   = 100000
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

	s := service.NewService(numDocuments, viewName, collectionName)
	s.InsertSampleData(collection)
	s.CreateView(db)
	viewTime := s.BenchmarkViewFind(db)
	aggTime := s.BenchmarkAggregationFind(collection)

	fmt.Printf("View find time: %v\n", viewTime)
	fmt.Printf("Aggregation find time: %v\n", aggTime)
}
