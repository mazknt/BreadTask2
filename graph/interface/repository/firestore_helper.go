package repository

import (
	"BreadTask2/graph/entity/dto"
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

/**
 * dbClientの作成
 */
func createFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, createClientError := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"), option.WithCredentialsFile(os.Getenv("CREDENTIAL_OPTION")))
	if createClientError != nil {
		log.Println("failed to create firebase client", createClientError)
		return nil, createClientError
	}
	return client, nil
}

/**
 * Collectionの取得
 * breadCollection
 */
func getBreadCollection(dbClient *firestore.Client) (dto.BreadCollection, error) {
	ctx := context.Background()
	iter := dbClient.Collection("breadCollection").Documents(ctx)
	breadCollection := &dto.BreadCollection{
		BreadDocuments: make(map[string]dto.BreadDocument),
	}

	// 取得したBreadCollection → model.BreadCollection への変換
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break // 全てのドキュメントを取得した
		}
		if err != nil {
			log.Println("何らかのエラー:", err)
			return dto.BreadCollection{}, err
		}
		var breadDocument dto.BreadDocument
		if err := doc.DataTo(&breadDocument); err != nil {
			log.Printf("Failed to map document data to struct: %v", err)
			return dto.BreadCollection{}, err
		}
		breadCollection.BreadDocuments[doc.Ref.ID] = breadDocument
	}
	return *breadCollection, nil
}

/**
 * Documentの取得
 * breadDocument
 */
func getBreadDocument(dbClient *firestore.Client, id string) (dto.BreadDocument, error) {
	ctx := context.Background()
	docRef := dbClient.Collection("breadCollection").Doc(id)
	doc, err := docRef.Get(ctx)
	if err != nil {
		log.Println("failed to get firebase document", err)
		return dto.BreadDocument{}, err
	}
	if !doc.Exists() {
		return dto.BreadDocument{}, nil
	}
	var breadDocument dto.BreadDocument
	err = doc.DataTo(&breadDocument)
	log.Println(doc)
	if err != nil {
		log.Println("failed doc.DataTo(&token)", err)
		return dto.BreadDocument{}, err
	}
	return breadDocument, nil
}
