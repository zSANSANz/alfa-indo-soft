package api

import (
	"context"
	"log"

	"alfa-indo-soft/service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllTags -> gets all tags in db
func GetAllTags(ctx *gin.Context) {
	filter := bson.D{{}}
	// findOptions := options.Find()
	// findOptions.SetLimit(5)

	var results []models.Tag
	cur, err := collectionTags.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Tag
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

func GetTagBySlug(ctx *gin.Context) {
	id := ctx.Param("id")

	filter := bson.A{
		bson.M{
			"$match": bson.M{
				"tag_slug": bson.M{"$regex": id},
			},
		},
	}

	var results []models.Tag
	cur, err := collectionTags.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Tag
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

// GetTags -> gets tags given id
func GetTags(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Tag{}
	filter := bson.M{"_id": docID}
	err := collectionTags.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    result,
		})
	}
}

// InsertTag -> insert one tag
func InsertTag(ctx *gin.Context) {
	var json models.Tag

	ctx.Bind(&json)

	termID := json.TermID
	tagSlug := json.TagSlug
	tagName := json.TagName
	tagDescription := json.TagDescription
	tagCount := json.TagCount

	tagCollection := models.Tag{
		TermID:         termID,
		TagSlug:        tagSlug,
		TagName:        tagName,
		TagDescription: tagDescription,
		TagCount:       tagCount,
	}

	result, err := collectionTags.InsertOne(context.Background(), tagCollection)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"_id":     result,
			"data":    tagCollection,
		})
	}
}

// UpdateTag -> update one tag
func UpdateTag(ctx *gin.Context) {
	id := ctx.Param("id")
	var json models.Tag

	ctx.Bind(&json)

	termID := json.TermID
	tagSlug := json.TagSlug
	tagName := json.TagName
	tagDescription := json.TagDescription
	tagCount := json.TagCount

	// docID, _ := primitive.ObjectIDFromHex(id)
	tagCollection := models.Tag{
		TermID:         termID,
		TagSlug:        tagSlug,
		TagName:        tagName,
		TagDescription: tagDescription,
		TagCount:       tagCount,
	}
	update := bson.M{
		"$set": tagCollection,
	}

	filter := bson.M{"tag_slug": id}
	result := models.Tag{}
	err := collectionTags.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    result,
		})
	}

}

// DeleteTag -> deletes tag based on id
func DeleteTag(ctx *gin.Context) {
	id := ctx.Param("id")
	// docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Tag{}
	filter := bson.M{"tag_slug": id}
	err := collectionTags.FindOneAndDelete(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    result,
		})
	}
}
