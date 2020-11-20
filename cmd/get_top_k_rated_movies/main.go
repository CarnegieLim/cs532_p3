package main


import (
	"context"
	"fmt"
	"log"

	"cs532/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	
	n := 270
	k := int64(5)
	genresA := "Documentary"
	genresB := "Short"

	database, _ := db.GetDB()
	collection := database.Collection("movies")
	options := options.Find()
	options.SetSort(bson.D{{"ratings.averageRating", -1}})
	options.SetLimit(k)
	
	cur, _ := collection.Find(context.TODO(), bson.M{"ratings.numVotes": bson.D{{ "$gte", n}},"genres": bson.D{{"$all",bson.A{genresA, genresB}}},}, options)

	type Result struct {
		PrimaryTitle string `json:"primaryTitle"`
	}

	for cur.Next(context.TODO()) {
		var result Result
		err := cur.Decode(&result)
		fmt.Println(result.PrimaryTitle)
		if err != nil {
			log.Fatal(err)
		}

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
}