package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalf("error connecting to mongo: %v", err)
	}

	collection := client.Database("docker-compose").Collection("hello-world")

	document, err := collection.InsertOne(context.Background(), bson.M{"Hello": "Mongo"})
	if err != nil {
		log.Fatalf("fail to write document to mongo")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := collection.FindOne(context.Background(), bson.M{"_id": document.InsertedID})

		var doc struct {
			Hello string `bson:"Hello"`
		}
		_ = result.Decode(&doc)

		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hello %v", doc.Hello))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
