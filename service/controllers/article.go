package api

import (
	"context"
	"log"

	"alfa-indo-soft/service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllArticles -> gets all articles in db
func GetAllArticles(ctx *gin.Context) {
	filter := bson.D{{}}
	// findOptions := options.Find()
	// findOptions.SetLimit(5)

	var results []models.Article
	cur, err := collectionArticles.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Article
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
			"result": results,
		})
	}
}

func GetArticleByTitleAndBody(ctx *gin.Context) {
	title := ctx.Param("title")

	filter := bson.A{
		bson.M{
			"$match": bson.M{
				"title": bson.M{"$regex": title},
			},
		},
	}

	var results []models.Article
	cur, err := collectionArticles.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Article
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
			"data":    results,
		})
	}
}

// GetArticles -> gets articles given id
func GetArticles(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Article{}
	filter := bson.M{"_id": docID}
	err := collectionArticles.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"result": result,
		})
	}
}

// InsertArticle -> insert one article
func InsertArticle(ctx *gin.Context) {
	var json models.Article

	ctx.Bind(&json)

	title := json.Title
	author := json.Author
	body := json.Body

	article := models.Article{
		Title:  title,
		Author: author,
		Body:   body,
	}

	result, err := collectionArticles.InsertOne(context.Background(), article)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"inserted_id": result.InsertedID,
		})
	}
}

// UpdateArticle -> update one article
func UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")
	body := ctx.PostForm("body")

	docID, _ := primitive.ObjectIDFromHex(id)
	article := models.Article{
		Title:  title,
		Author: author,
		Body:   body,
	}
	update := bson.M{
		"$set": article,
	}

	filter := bson.M{"_id": docID}
	result := models.Article{}
	err := collectionArticles.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"updated_id": result.ID,
		})
	}

}

// DeleteArticle -> deletes article based on id
func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Article{}
	filter := bson.M{"_id": docID}
	err := collectionArticles.FindOneAndDelete(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"deleted_id": result.ID,
		})
	}
}
