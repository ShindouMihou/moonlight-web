package chapters

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/books"
	"server/modules"
)

var ErrNoUpdateContents = errors.New("cannot update a chapter with both title and contents empty")

func Insert(chapter Chapter) (*Chapter, error) {
	result, err := modules.GetCollection("chapters").InsertOne(context.TODO(), chapter)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	chapter.Id = id.Hex()
	return &chapter, nil
}

func Update(id string, title string, contents string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if title == "" && contents == "" {
		return ErrNoUpdateContents
	}
	updates := bson.M{"contents": contents}
	if title != "" {
		updates["title"] = title
	}
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", updates}}
	result, err := modules.GetCollection("chapters").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func Find(id string) (*Chapter, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var book Chapter
	filter := bson.D{{"_id", objectId}}
	err = modules.GetCollection("chapters").FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func DeleteAssociated(book books.Book) error {
	filter := bson.D{{"book", book}}
	_, err := modules.GetCollection("chapters").DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func Paginate(book string, next string) ([]Chapter, error) {
	filter := bson.D{{"book", book}}
	if next != "" {
		nextId, err := primitive.ObjectIDFromHex(next)
		if err != nil {
			return nil, err
		}
		filter = bson.D{
			{
				"$and",
				bson.A{
					filter,
					bson.D{{"_id", bson.D{{"$gt", nextId}}}},
				},
			},
		}
	}
	cursor, err := modules.GetCollection("chapters").Find(context.TODO(), filter, options.Find().SetLimit(24))
	if err != nil {
		return nil, err
	}
	// IMPORTANT: if not created like this, empty results will result in a nil value
	//goland:noinspection GoPreferNilSlice
	chapters := []Chapter{}
	if err := cursor.All(context.TODO(), &chapters); err != nil {
		return nil, err
	}
	return chapters, nil
}

func Delete(chapter Chapter) error {
	objectId, err := primitive.ObjectIDFromHex(chapter.Id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", objectId}}
	_, err = modules.GetCollection("chapters").DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
