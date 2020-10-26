package notesdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Note struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

var client *mongo.Client
var err error

func CreateConnection(uri string) {
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return
}

func InsertNote(title string, content string) {
	post := Note{title, content}

	collection := client.Database("notesdb").Collection("notes")

	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted note with ID:", insertResult.InsertedID)
}

func GetNotes() []Note {
	collection := client.Database("notesdb").Collection("notes")
	var results []Note

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Note
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	return results
}
