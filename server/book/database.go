package book

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/chapter"
	"server/modules"
)

func Insert(book Book) (*Book, error) {
	result, err := modules.GetCollection("books").InsertOne(context.TODO(), book)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	book.Id = id.Hex()
	return &book, nil
}

func Find(id string) (*Book, error) {
	var book Book
	filter := bson.D{{"_id", primitive.ObjectIDFromHex(id)}}
	err := modules.GetCollection("books").FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func Replace(book Book) (*Book, error) {
	filter := bson.D{{"_id", primitive.ObjectIDFromHex(book.Id)}}
	_, err := modules.GetCollection("books").ReplaceOne(context.TODO(), filter, book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func Delete(book Book) error {
	filter := bson.D{{"_id", primitive.ObjectIDFromHex(book.Id)}}
	_, err := modules.GetCollection("books").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	err = chapter.DeleteAssociated(book)
	if err != nil {
		return err
	}
	return nil
}

func All() ([]Book, error) {
	cursor, err := modules.GetCollection("books").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var books []Book
	if err := cursor.All(context.TODO(), &books); err != nil {
		return nil, err
	}
	return books, nil
}
