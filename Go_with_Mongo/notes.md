const connString = "mongodb://localhost:27017/"

// db name
const dbName = "go_netflix"

// collection name
const colName = "watchlist"

// important-> mark mongo connection
var collection *mongo.Collection

// connect to mongodb -> initilize method -> it runs first time when our application runs(only one time)
func init() {
	// decalair client options
	clientOptions := options.Client().ApplyURI(connString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection is successfull")

	// now stablish database and collectin with collection reference
	collection = client.Database(dbName).Collection(colName)

	// collection reference
	fmt.Println("collection reference is ready")
}

// Mongodb helpers

// insert record in db
func insertOne(movie model.Netflix) {
	// it insert one movie record
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert 1 record in mongodb with id: ", inserted.InsertedID)
}

// updae value
func updateValue(movieId string) {
	//first convert our id to mongo understandable form
	id, _ := primitive.ObjectIDFromHex(movieId)

	//* define filter criteriya
	// bson.M -> it doesnt care about lower case and uper case
	// we can also use bson.D
	filter := bson.M{"_id": id}

	//* now pass updated values
	update := bson.M{"$set": bson.M{"watched": true}}

	// now call update method
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

// get all collection from db
func getAllMovies() []primitive.M {
	// pass empty {} sothat it finds all collection (we give no filter)
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	// now make a movies array in which we put our all collections
	var movies []primitive.M

	// now loop on our cursor and put each coll. into the movies array
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