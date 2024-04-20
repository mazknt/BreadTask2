package dto

import "time"

type BreadCollection struct {
	BreadDocuments map[string]BreadDocument
}

type BreadDocument struct {
	Name      string    `firestore:"name"`
	CreatedAt time.Time `firestore:"createdAt"`
}
