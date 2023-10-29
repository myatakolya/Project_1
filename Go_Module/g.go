package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// You will be using this Trainer type later in the program
type Day struct {
	Subject string
	Name    string
	Type    string
	Time    string
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://myatakolyajr:0968715133@cluster0.ebfmjow.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	/*err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")*/
	collection := client.Database("proj").Collection("shedule")

	fmt.Println(collection)
}
