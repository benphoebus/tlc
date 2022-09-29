package configs

import(
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create a ConnectDB function that first configures the client to use the correct URI and check for errors. Secondly, we defined a timeout of 10 seconds we wanted to use when trying to connect. Thirdly, check if there is an error while connecting to the database and cancel the connection if the connecting period exceeds 10 seconds. Finally, we pinged the database to test our connection and returned the client instance.

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}


// Client instance 
// Create a DB variable instance of the ConnectDB. This will come in handy when creating collections.
var DB *mongo.Client = ConnectDB()

// getting databse collections 
// Create a GetCollection function to retrieve and create collections on the database.
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}