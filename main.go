package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongodb "GinFrameWork/MongoDB" // Adjust the import path as necessary
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

	// Create a context
	ctx := context.TODO()

	// Define the filter and option
	filter := bson.D{
		{"maths", bson.D{{"$gt", 70}}},
	}
	option := bson.D{{"_id", 0}}

	// Create an instance of MongoDBOperations
	var dbOps mongodb.DatabaseOperations = mongodb.MongoDBOperations{}

	if result, err := dbOps.Insert(client, ctx, "gfg", "marks", bson.D{{"maths", 80}, {"science", 90}}); err != nil {
		panic(err)
	} else {
		fmt.Println("Inserted document with ID:", result.InsertedID)
	}

	// Call the Query method
	cursor, err := dbOps.Query(client, ctx, "gfg", "marks", filter, option)
	if err != nil {
		panic(err)
	}

	var results []bson.D

	// Get BSON objects from cursor
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	if len(results) == 0 {
		fmt.Println("No documents found")
		return
	}
	// Print the results
	for _, result := range results {
		fmt.Println(result)
		resultMap := result.Map()
		fmt.Println(resultMap["maths"])
	}

	// Define the filter and update
	filterForUpdate := bson.D{{"maths", 80}}
	update := bson.D{{"$set", bson.D{{"maths", 85}, {"science", 90}, {"history", 90}, {"geo", 90}}}}

	if updateresult, err := dbOps.Update(client, ctx, "gfg", "marks", filterForUpdate, update, false); err != nil {

		panic(err)
	} else {
		fmt.Printf("Matched %v documents and updated %v documents.\n", updateresult.MatchedCount, updateresult.ModifiedCount)
	}

	// Define the filter for deletion
	filterForDeletion := bson.D{{"maths", 85}}
	if deleteResult, err := dbOps.Delete(client, ctx, "gfg", "marks", filterForDeletion, true); err != nil {
		panic(err)
	} else {
		fmt.Printf("Deleted %v documents.\n", deleteResult.DeletedCount)
	}
}
