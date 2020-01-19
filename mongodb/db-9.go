package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

	persons := []interface{}{james, frankie}

	insertManyResult, err := collection.InsertMany(context.TODO(), persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	log.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// updating Frankie’s age
	filter := bson.D{{"name", "Frankie"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// reading the data
	// create a value into which the result can be decoded
	var result Person

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	log.Printf("Found a single document: %+v\n", result)

	// finding multiple documents and returning the cursor
	// pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)
	filter = bson.D{{}}

	// here's an array in which you can store the decoded documents
	var results []*Person

	// passing nil as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// finding multiple documents returns a cursor
	// iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Person
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	log.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// delete documents
	filter = bson.D{{}}

	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the persons collection\n", deleteResult.DeletedCount)
	log.Printf("Deleted %v documents in the persons collection\n", deleteResult.DeletedCount)

	// closing the connection
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
		log.Println("Connection to MongoDB closed.")
	}

}
