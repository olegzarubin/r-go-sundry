package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	// create client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	/*
	   clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	   client, err := mongo.Connect(context.TODO(), clientOptions)
	   if err != nil {
	       log.Fatal(err)
	   }
	*/

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	log.Println("Connected to MongoDB")

	// let’s insert a single document
	collection := client.Database("mydb").Collection("persons")

	ruan := Person{"Ruan", 34, "Cape Town"}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	log.Println("Inserted a Single Document: ", insertResult.InsertedID)

	// writing more than one document
	james := Person{"James", 32, "Nairobi"}
	frankie := Person{"Frankie", 31, "Nairobi"}

	trainers := []interface{}{james, frankie}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	log.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
