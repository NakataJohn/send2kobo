package repository

import (
	"context"
	"send2kobo/domain"
	"send2kobo/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bookRepository struct {
	database   mongo.Database
	collection string
}

func NewBookRepository(db mongo.Database, collection string) domain.BookRepository {
	return &bookRepository{
		database:   db,
		collection: collection,
	}
}

func (br *bookRepository) Create(c context.Context, book *domain.Book) error {
	collection := br.database.Collection(br.collection)
	_, err := collection.InsertOne(c, book)
	return err
}

func (br *bookRepository) Fetch(c context.Context) ([]domain.Book, error) {
	collection := br.database.Collection(br.collection)
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var books []domain.Book

	err = cursor.All(c, &books)
	if books == nil {
		return []domain.Book{}, err
	}
	return books, err
}

func (br *bookRepository) GetByTitle(c context.Context, title string) (domain.Book, error) {
	collection := br.database.Collection(br.collection)
	var book domain.Book
	err := collection.FindOne(c, bson.M{"title": title}).Decode(&book)
	return book, err
}

func (br *bookRepository) GetByID(c context.Context, id string) (domain.Book, error) {
	collection := br.database.Collection(br.collection)
	var book domain.Book
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}
	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&book)
	return book, err
}

func (br *bookRepository) DeleteByID(c context.Context, id string) error {
	collection := br.database.Collection(br.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(c, bson.M{"_id": idHex})
	return err
}

