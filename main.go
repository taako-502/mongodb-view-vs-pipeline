package main

import (
	"context"
	"log"

	"github.com/taako-502/mongodb-view-vs-pipeline/pkg/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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
}
