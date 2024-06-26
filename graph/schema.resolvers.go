package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"BreadTask2/graph/interface/repository"
	model1 "BreadTask2/graph/model"
	"BreadTask2/graph/usecase"
	"context"
	"log"
)

// GetAllBreads is the resolver for the GetAllBreads field.
func (r *queryResolver) GetAllBreads(ctx context.Context) ([]*model1.Bread, error) {
	breadService := usecase.BreadServiceImpl{
		Firestore: &repository.FirestoreImpl{},
	}
	allBreadsList, getAllBreadsError := breadService.GetAllBreads()
	if getAllBreadsError != nil {
		log.Println("getAllBreadsError:", getAllBreadsError)
		return nil, getAllBreadsError
	}
	return allBreadsList, nil
}

// GetBread is the resolver for the getBread field.
func (r *queryResolver) GetBread(ctx context.Context, id string) (*model1.Bread, error) {
	breadService := usecase.BreadServiceImpl{
		Firestore: &repository.FirestoreImpl{},
	}
	bread, getBreadError := breadService.GetBread(id)
	if getBreadError != nil {
		log.Println("getBreadError:", getBreadError)
		return nil, getBreadError
	}
	return bread, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
