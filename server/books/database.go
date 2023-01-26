package books

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"server/modules"
)

var ErrNoUpdateContents = errors.New("cannot update a book with name empty")

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
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var book Book
	filter := bson.D{{"_id", objectId}}
	err = modules.GetCollection("books").FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func Update(id string, name string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if name == "" {
		return ErrNoUpdateContents
	}
	// :update this when there are more fields to books to make it more specific updates
	updates := bson.M{"name": name}
	filter := bson.D{{"_id", objectId}}
	result, err := modules.GetCollection("books").UpdateOne(context.TODO(), filter, bson.M{"$set": updates})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func Delete(book Book) error {
	objectId, err := primitive.ObjectIDFromHex(book.Id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	_, err = modules.GetCollection("books").DeleteOne(context.TODO(), filter)
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
	// IMPORTANT: if not created like this, empty results will result in a nil value
	//goland:noinspection GoPreferNilSlice
	books := []Book{}
	if err := cursor.All(context.TODO(), &books); err != nil {
		return nil, err
	}
	return books, nil
}
