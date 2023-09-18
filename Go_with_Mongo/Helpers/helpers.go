package helpers

import (
	"context"
	"fmt"
	"log"

	model "github.com/uniquepg/gowithmongo/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ! Define important variables for db connection
const ConnString = "mongodb://localhost:27017/"

const dbName = "go_netflix" // db name

const colName = "watchlist" // collection name

// * important-> define mongo connection
var collection *mongo.Collection

// ! connect to mongodb ->
// *initilize method -> it runs first time when our application runs(only one time)
func init() {
	// decalair client options
	clientOptions := options.Client().ApplyURI(ConnString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection is successfull...")

	// now stablish database and collectin with collection reference
	collection = client.Database(dbName).Collection(colName)

	// collection reference
	fmt.Println("collection reference is ready")
}

//! Mongodb helpers functions that interact with db

// insert record in db
func InsertOne(movie model.Netflix) {
	// it insert one movie record
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert 1 record in mongodb with id: ", inserted.InsertedID)
}

// * updae colletion values
func UpdateValue(movieId string, movie model.Netflix) {
	//first convert our id to mongo understandable form
	id, _ := primitive.ObjectIDFromHex(movieId)

	//* define filter criteriya
	// bson.M -> it doesnt care about lower case and uper case
	// we can also use bson.D
	filter := bson.M{"_id": id}

	//* now pass updated values
	update := bson.M{"$set": bson.M{"movie": movie.Movie, "director": movie.Director, "year": movie.Year, "watched": movie.Watched}}

	// now call update method
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

// get all collection from db
func GetAllMovies() []primitive.M {
	// pass empty {} sothat it finds all collection (we give no filter)
	//* Finding multiple documents returns a cursor
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	// now make a movies array in which we put our all collections
	var movies []primitive.M

	// Iterating through the cursor allows us to decode documents one at a time
	for cursor.Next(context.Background()) {

		var movie bson.M
		err := cursor.Decode(&movie) // decode our movie

		if err != nil {
			log.Fatal(err)
		}

		// now append each movie into the movies array
		movies = append(movies, movie)

	}

	// now close cursor
	cursor.Close(context.Background())
	return movies
}

// delete one record
func DeleteOneRecord(movieId string) {
	// convert id string into mongo readable idobject
	id, _ := primitive.ObjectIDFromHex(movieId)

	// filter criteria
	filter := bson.M{"_id": id}

	// call delete method
	count, err := collection.DeleteOne(context.Background(), filter) // it gives count of deleted document

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one record with count: ", count)
}

// * delete all records
func DeleteAllDocuments() {

	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted All records with count: ", result.DeletedCount)
}
