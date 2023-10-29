package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("users")

	//Create (Insert Documents)
	user := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	}

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

	//Read (Find Documents)
	// Find a single document
	var result map[string]interface{}
	err = collection.FindOne(context.TODO(), map[string]interface{}{"name": "John Doe"}).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found a single document: ", result)

	//Update Documents
	filter := map[string]interface{}{"name": "John Doe"}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"age": 31,
		},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//Delete document
	deleteResult, err := collection.DeleteOne(context.TODO(), map[string]interface{}{"name": "John Doe"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//Disconnect from MongoDb
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
