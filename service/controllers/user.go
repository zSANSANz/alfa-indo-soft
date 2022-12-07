package api

import (
	"context"
	"log"

	"alfa-indo-soft/service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers(ctx *gin.Context) {
	filter := bson.A{}

	var results []models.User
	cur, err := collectionUsers.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"result":  results,
		})
	}
}
