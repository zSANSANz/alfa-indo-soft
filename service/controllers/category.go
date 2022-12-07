package api

import (
	"context"
	"log"

	"alfa-indo-soft/service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllCategories -> gets all categories in db
func GetAllCategories(ctx *gin.Context) {
	filter := bson.D{{}}
	// findOptions := options.Find()
	// findOptions.SetLimit(5)

	var results []models.Category
	cur, err := collectionCategories.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Category
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

func GetCategoryBySlug(ctx *gin.Context) {
	id := ctx.Param("id")

	filter := bson.A{
		bson.M{
			"$match": bson.M{
				"category_nicename": bson.M{"$regex": id},
			},
		},
	}

	var results []models.Category
	cur, err := collectionCategories.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Category
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

// GetCategories -> gets categories given id
func GetCategories(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Category{}
	filter := bson.M{"_id": docID}
	err := collectionCategories.FindOne(context.Background(), filter).Decode(&result)
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

// InsertCategory -> insert one category
func InsertCategory(ctx *gin.Context) {
	var json models.Category

	ctx.Bind(&json)

	termID := json.TermID
	categoryNicename := json.CategoryNicename
	// category := json.Category
	catName := json.CatName
	categoryDescription := json.CategoryDescription

	categoryCollection := models.Category{
		TermID:           termID,
		CategoryNicename: categoryNicename,
		// Category:            category,
		CatName:             catName,
		CategoryDescription: categoryDescription,
	}

	result, err := collectionCategories.InsertOne(context.Background(), categoryCollection)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"_id":     result,
			"data":    categoryCollection,
		})
	}
}

// UpdateCategory -> update one category
func UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	var json models.Category

	ctx.Bind(&json)

	termID := json.TermID
	categoryNicename := json.CategoryNicename
	// category := json.Category
	catName := json.CatName
	categoryDescription := json.CategoryDescription

	// docID, _ := primitive.ObjectIDFromHex(id)
	categoryCollection := models.Category{
		TermID:           termID,
		CategoryNicename: categoryNicename,
		// Category:            category,
		CatName:             catName,
		CategoryDescription: categoryDescription,
	}
	update := bson.M{
		"$set": categoryCollection,
	}

	filter := bson.M{"category_nicename": id}
	result := models.Category{}
	err := collectionCategories.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)

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

// DeleteCategory -> deletes category based on id
func DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	// docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Category{}
	filter := bson.M{"category_nicename": id}
	err := collectionCategories.FindOneAndDelete(context.Background(), filter).Decode(&result)

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
