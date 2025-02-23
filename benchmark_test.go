package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/taako-502/mongodb-view-vs-pipeline/pkg/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	mongoURI       = "mongodb://localhost:27017"
	databaseName   = "testdb"
	collectionName = "testcollection"
	viewName       = "testview"
)

func BenchmarkMongoDBViewVSPipeline(b *testing.B) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		b.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(databaseName)
	collection := db.Collection(collectionName)

	s := service.NewService(collection, viewName, collectionName)
	if err := s.CreateView(db); err != nil {
		b.Fatalf("Failed to create view: %v", err)
	}
	if err := s.DropSampleData(); err != nil {
		b.Fatalf("Failed to drop sample data: %v", err)
	}

	test := []int64{10000, 100000, 1000000, 10000000}
	for _, num := range test {
		b.Run(fmt.Sprintf("Documents_%d", num), func(b *testing.B) {
			if err := s.InsertSampleData(num); err != nil {
				b.Fatalf("Failed to insert sample data: %v", err)
			}
			b.ResetTimer()
			for b.Loop() {
				if _, err := s.BenchmarkViewFind(db); err != nil {
					b.Fatalf("Failed to find view: %v", err)
				}
				if _, err = s.BenchmarkAggregationFind(); err != nil {
					b.Fatalf("Failed to find aggregation: %v", err)
				}
			}
			b.StopTimer()
			if err := s.DropSampleData(); err != nil {
				b.Fatalf("Failed to drop sample data: %v", err)
			}
		})
	}
}
