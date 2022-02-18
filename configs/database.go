package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConfig() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoEnvUri()))

	if err != nil {
		log.Fatal(err)
	}

	// Initiate timeout for 10 seconds
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

// Create instance of DBConfig() for collection creation
var DB *mongo.Client = DBConfig()

func GetCollection(
	client *mongo.Client,
	name string,
) *mongo.Collection {
	collection := client.Database(DatabaseName()).Collection(name)
	return collection
}
