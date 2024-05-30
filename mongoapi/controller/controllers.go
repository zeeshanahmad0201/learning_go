package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/zeeshanahmad0201/learning_go/tree/main/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://zeeshan:123@golearningdb.gpiyigt.mongodb.net/"
const dbName = "netflix"
const collectionName = "watchedList"

var collection *mongo.Collection

func init() {
	// client options
	options := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal("unable to connect to mongodb")
	}

	fmt.Println("mongo db connection success")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection ready to use")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("model has been inserted with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"watched": true}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result)
}

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie deleted with count: ", result)
}

func deleteAllMovies() {
	collection.DeleteMany(context.Background(), bson.D{{}})
}
