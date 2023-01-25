package modules

import (
	"context"
	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func InitMongo() {
	golog.Info("loading mongodb connection")
	uri := EnsureEnv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetAppName("Helena"))
	if err != nil {
		golog.Fatal("failed to connect to mongo : ", err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		golog.Fatal("failed to receive heartbeat from mongo : ", err)
	}
	MongoClient = client
}

func GetCollection(name string) *mongo.Collection {
	return MongoClient.Database("moonlight").Collection(name)
}
