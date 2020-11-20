package main

import (
	"strconv"
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"cs532/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/topK", func(c *gin.Context) {
		genresA := c.DefaultQuery("genresA", "")
		genresB := c.DefaultQuery("genresB", "")
		numVote := c.DefaultQuery("numVote", "")
		k := c.DefaultQuery("k", "")
		n, _ := strconv.Atoi(numVote)
		kValue, _ := strconv.Atoi(k)

		database, _ := db.GetDB()
		collection := database.Collection("movies")
		options := options.Find()
		options.SetSort(bson.D{{"ratings.averageRating", -1}})
		options.SetLimit(int64(kValue))

		cur, _ := collection.Find(context.TODO(), bson.M{"ratings.numVotes": bson.D{{"$gte", n}}, "genres": bson.D{{"$all", bson.A{genresA, genresB}}}}, options)

		type Result struct {
			PrimaryTitle string `json:"primaryResult"`
		}

		var results []Result

		for cur.Next(context.TODO()) {
			var result Result
			err := cur.Decode(&result)
			results = append(results, result)
			// fmt.Println(result.PrimaryTitle)
			if err != nil {
				log.Fatal(err)
			}

		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.TODO())

		c.JSON(200, gin.H{
			"results": results,
		})
	})

	router.GET("/avgAge", func(c *gin.Context) {
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

		// fmt.Println(pipeline)
		// return

		opts := options.Aggregate()
		if cur, err = collection.Aggregate(ctx, MongoPipeline(pipeline), opts); err != nil {
			log.Fatal(err)
		}
		var results []bson.M
		var doc bson.M
		for cur.Next(ctx) {
			cur.Decode(&doc)
			results = append(results, doc)
			// fmt.Println(doc)
			// total++
		}

		defer cur.Close(ctx)

		c.JSON(200, gin.H{
			"results": results,
		})
	})



	router.Run()
}

func scoreToSentiment(score uint8) string {
	if score == 0 {
		return "negative"
	}
	return "positive"
}
