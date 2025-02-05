package service

type service struct {
	viewName       string
	collectionName string
	numDocuments   int64
}

func NewService(
	numDocuments int64,
	viewName string,
	collectionName string,
) *service {
	return &service{
		viewName:       viewName,
		collectionName: collectionName,
		numDocuments:   numDocuments,
	}
}
