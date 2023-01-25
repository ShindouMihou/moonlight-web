package chapter

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/book"
	"server/modules"
)

func Insert(chapter Chapter) (*Chapter, error) {
	result, err := modules.GetCollection("chapters").InsertOne(context.TODO(), chapter)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	chapter.Id = id.Hex()
	return &chapter, nil
}

func Find(id string) (*Chapter, error) {
	var book Chapter
	filter := bson.D{{"_id", primitive.ObjectIDFromHex(id)}}
	err := modules.GetCollection("chapters").FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func Replace(chapter Chapter) (*Chapter, error) {
	filter := bson.D{{"_id", primitive.ObjectIDFromHex(chapter.Id)}}
	_, err := modules.GetCollection("chapters").ReplaceOne(context.TODO(), filter, chapter)
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}

func DeleteAssociated(book book.Book) error {
	filter := bson.D{{"book", primitive.ObjectIDFromHex(book.Id)}}
	_, err := modules.GetCollection("chapters").DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func Paginate(book string, next *string) ([]Chapter, error) {
	filter := bson.D{{"book", primitive.ObjectIDFromHex(book)}}
	if next != nil {
		filter = bson.D{
			{
				"$and",
				bson.A{
					filter,
					bson.D{{"_id", bson.D{{"$gt", primitive.ObjectIDFromHex(*next)}}}},
				},
			},
		}
	}
	cursor, err := modules.GetCollection("chapters").Find(context.TODO(), filter, options.Find().SetLimit(24))
	if err != nil {
		return nil, err
	}
	var chapters []Chapter
	if err := cursor.All(context.TODO(), &chapters); err != nil {
		return nil, err
	}
	return chapters, nil
}
