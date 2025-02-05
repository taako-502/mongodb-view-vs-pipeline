package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	mongoURI        = "mongodb://localhost:27017"
	databaseName    = "testdb"
	collectionName  = "testcollection"
	viewName        = "testview"
	numDocuments    = 100000
)

func main() {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(databaseName)
	collection := db.Collection(collectionName)

	insertSampleData(collection)

	createView(db)

	viewTime := benchmarkViewFind(db)

	aggTime := benchmarkAggregationFind(collection)

	fmt.Printf("View find time: %v\n", viewTime)
	fmt.Printf("Aggregation find time: %v\n", aggTime)
}

func insertSampleData(collection *mongo.Collection) {
	ctx := context.TODO()
	count, _ := collection.CountDocuments(ctx, bson.M{})
	if count >= numDocuments {
		fmt.Println("Sample data already exists, skipping insertion.")
		return
	}

	fmt.Println("Inserting sample data...")
	docs := make([]interface{}, numDocuments)
	for i := 0; i < numDocuments; i++ {
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

func createView(db *mongo.Database) {
	fmt.Println("Creating MongoDB View...")

	_ = db.Collection(viewName).Drop(context.TODO())

	// ビューの作成
	command := bson.D{
		{Key: "create", Value: viewName},
		{Key: "viewOn", Value: collectionName},
		{Key: "pipeline", Value: bson.A{
			bson.D{{Key: "$match", Value: bson.M{"score": bson.M{"$gte": 50}}}}, // スコアが50以上のデータを取得
		}},
	}

	err := db.RunCommand(context.TODO(), command).Err()
	if err != nil {
		log.Fatalf("Failed to create view: %v", err)
	}

	fmt.Println("View created successfully.")
}

func benchmarkViewFind(db *mongo.Database) time.Duration {
	ctx := context.TODO()
	collection := db.Collection(viewName)

	start := time.Now()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to find from view: %v", err)
	}
	defer cursor.Close(ctx)

	count := 0
	for cursor.Next(ctx) {
		count++
	}

	elapsed := time.Since(start)
	fmt.Printf("View Find fetched %d documents.\n", count)
	return elapsed
}

func benchmarkAggregationFind(collection *mongo.Collection) time.Duration {
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