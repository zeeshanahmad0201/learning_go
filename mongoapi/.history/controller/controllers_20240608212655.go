package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
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

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		cursor.Decode(&movie)

		movies = append(movies, movie)
	}

	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www.form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted!")
}
