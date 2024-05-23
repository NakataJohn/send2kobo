package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBook = "books"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id"`
	Hash      string             `bson:"hash"`
	Title     string             `bson:"title"`
	Path      string             `bson:"path"`
	Kepubpath string             `bson:"kpath"`
	KHash     string             `bson:"khash"`
}

type BookRepository interface {
	Create(c context.Context, book *Book) error
	Fetch(c context.Context) ([]Book, error)
	GetByTitle(c context.Context, title string) (Book, error)
	GetByID(c context.Context, id string) (Book, error)
	DeleteByID(c context.Context, id string) error
}

type BookUsecase interface {
	Create(c context.Context, book *Book) error
	Fetch(c context.Context) ([]Book, error)
	GetByID(c context.Context, id string) (Book, error)
	DeleteByID(c context.Context, id string) error
}
