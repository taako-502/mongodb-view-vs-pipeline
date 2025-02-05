package service

import "go.mongodb.org/mongo-driver/v2/mongo"

type service struct {
	collection     *mongo.Collection
	viewName       string
	collectionName string
}

func NewService(
	collection *mongo.Collection,
	viewName string,
	collectionName string,
) *service {
	return &service{
		collection:     collection,
		viewName:       viewName,
		collectionName: collectionName,
	}
}
