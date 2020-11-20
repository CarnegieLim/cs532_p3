package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cs532/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	database, _ := db.GetDB()
	collection := database.Collection("actors")
	var ctx = context.Background()
	var cur *mongo.Cursor
	var err error

	pipeline := `
	[{
		"$project": { 
			"primaryName": "$primaryName" ,
			"primaryProfession": "$primaryProfession",
			"birthYear": "$birthYear",
			"deathYear": {"$ifNull": ["$deathYear", 2020]}
		}
	}, {
		"$unwind": "$primaryProfession"
	}, {
		"$group": { 
			"_id": {"primaryProfession": "$primaryProfession"},
			"age": { "$avg": {"$subtract": ["$deathYear", "$birthYear"]} } 
		}
	}]`

	fmt.Println(pipeline)
	// return

	opts := options.Aggregate()
	if cur, err = collection.Aggregate(ctx, MongoPipeline(pipeline), opts); err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	total := 0
	var doc bson.M
	for cur.Next(ctx) {
		cur.Decode(&doc)
		fmt.Println(doc)
		total++
	}

	fmt.Println(total)

}

func MongoPipeline(str string) mongo.Pipeline {
	var pipeline = []bson.D{}
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &doc)
		pipeline = append(pipeline, doc)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}
/*
func PreProcess() {
	initDB()
    getFieldData()
    convertToArrays()
    embedRatingInfo()
    embedActorsArray()
    embedAkaArray()
    embedCrewArray()
}
*/