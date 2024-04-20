package usecase

import (
	"BreadTask2/graph/interface/repository"
	"BreadTask2/graph/model"
	"log"
)

type BreadService interface {
	GetAllBreads() ([]*model.Bread, error)
	GetBread(id string) ([]*model.Bread, error)
}

type BreadServiceImpl struct {
	Firestore repository.Firestore
}

/**
 * DB内の全てのパン情報を取得
 */
func (s BreadServiceImpl) GetAllBreads() ([]*model.Bread, error) {
	allBreadsList, getAllBreadsError := s.Firestore.GetAllBreads()
	if getAllBreadsError != nil {
		log.Println("getAllBreadsError: ", getAllBreadsError)
		return nil, getAllBreadsError
	}
	return allBreadsList, nil
}

/**
 * DB内の指定のパン情報を取得
 */
func (s BreadServiceImpl) GetBread(id string) (*model.Bread, error) {
	bread, getBreadError := s.Firestore.GetBread(id)
	if getBreadError != nil {
		log.Println("getBreadError: ", getBreadError)
		return nil, getBreadError
	}
	return bread, nil
}
