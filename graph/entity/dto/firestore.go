package dto

import (
	"BreadTask2/graph/model"
	"time"
)

type BreadCollection struct {
	BreadDocuments map[string]BreadDocument
}

type BreadDocument struct {
	BreadInfo model.BreadInfo `firestore:"breadInfo"`
	CreatedAt time.Time       `firestore:"createdAt"`
	UpdatedAT time.Time       `firestore:"updatedAt"`
}
