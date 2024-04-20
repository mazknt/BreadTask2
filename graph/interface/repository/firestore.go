package repository

import (
	"BreadTask2/graph/model"
)

type Firestore interface {
	// 取得
	GetAllBreads() ([]*model.Bread, error)
	GetBread(id string) (*model.Bread, error)
}

type FirestoreImpl struct{}

func (f *FirestoreImpl) GetAllBreads() ([]*model.Bread, error) {
	// クライアントの生成
	dbClient, createDBClientError := createFirestoreClient()
	if createDBClientError != nil {
		return nil, createDBClientError
	}
	defer dbClient.Close()

	// パン情報の取得
	breadCollection, getTokenCollectionError := getBreadCollection(dbClient)
	if getTokenCollectionError != nil {
		return nil, getTokenCollectionError
	}
	var allBreadsList []*model.Bread
	for id, breadDocument := range breadCollection.BreadDocuments {
		bread := &model.Bread{
			ID: id,
			BreadInfo: &model.BreadInfo{
				Name:      breadDocument.Name,
				CreatedAt: breadDocument.CreatedAt,
			},
		}
		allBreadsList = append(allBreadsList, bread)
	}
	return allBreadsList, nil
}

func (f *FirestoreImpl) GetBread(id string) (*model.Bread, error) {
	// クライアントの生成
	dbClient, createDBClientError := createFirestoreClient()
	if createDBClientError != nil {
		return nil, createDBClientError
	}
	defer dbClient.Close()

	// 指定のパン情報の取得
	breadDocument, getBreadDocumentError := getBreadDocument(dbClient, id)
	if getBreadDocumentError != nil {
		return nil, getBreadDocumentError
	}
	bread := &model.Bread{
		ID: id,
		BreadInfo: &model.BreadInfo{
			Name:      breadDocument.Name,
			CreatedAt: breadDocument.CreatedAt,
		},
	}
	return bread, nil
}
