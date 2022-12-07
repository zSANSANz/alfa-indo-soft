package api

import (
	"context"
	"log"

	"alfa-indo-soft/service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllPosts -> gets all Posts in db
func GetAllPosts(ctx *gin.Context) {
	filter := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "author",
				"localField":   "author",
				"foreignField": "author_id",
				"as":           "authors",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "categories",
				"localField":   "categories",
				"foreignField": "term_id",
				"as":           "category",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "tags",
				"localField":   "tags",
				"foreignField": "term_id",
				"as":           "tag",
			},
		},
	}

	var results []models.Post
	cur, err := collectionPosts.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Post
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

// GetPosts -> gets posts given id
func GetPosts(ctx *gin.Context) {
	id := ctx.Param("id")
	docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Post{}
	filter := bson.M{"_id": docID}
	err := collectionPosts.FindOne(context.Background(), filter).Decode(&result)
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

func GetPostBySlug(ctx *gin.Context) {
	id := ctx.Param("id")

	filter := bson.A{
		bson.M{
			"$match": bson.M{
				"slug": bson.M{"$regex": id},
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "author",
				"localField":   "author",
				"foreignField": "author_id",
				"as":           "authors",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "categories",
				"localField":   "categories",
				"foreignField": "term_id",
				"as":           "category",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "tags",
				"localField":   "tags",
				"foreignField": "term_id",
				"as":           "tag",
			},
		},
	}

	var results []models.Post
	cur, err := collectionPosts.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem models.Post
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

// InsertPost -> insert one post
func InsertPost(ctx *gin.Context) {
	var json models.Post

	ctx.Bind(&json)

	title := json.Title
	slug := json.Slug
	date := json.Date
	url := json.Url
	username := json.Username
	description := json.Description
	content := json.Content
	except := json.Except
	commentStatus := json.CommentStatus
	status := json.Status
	ctype := json.Type
	postSurvey := json.PostSurvey
	// postParent := json.PostParent
	postType := json.PostType
	imageId := json.ImageId
	views := json.Views
	author := json.Author
	tags := json.Tags
	categories := json.Categories
	id := json.ID

	post := models.Post{
		ID:            id,
		Title:         title,
		Date:          date,
		Url:           url,
		Username:      username,
		Description:   description,
		Content:       content,
		Except:        except,
		CommentStatus: commentStatus,
		Slug:          slug,
		Status:        status,
		Type:          ctype,
		PostSurvey:    postSurvey,
		// PostParent:    postParent,
		PostType:   postType,
		ImageId:    imageId,
		Views:      views,
		Author:     author,
		Tags:       tags,
		Categories: categories,
	}

	result, err := collectionPosts.InsertOne(context.Background(), post)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"_id":     result,
			"data":    post,
		})
	}
}

// UpdatePost -> update one post
func UpdatePost(ctx *gin.Context) {
	var json models.Post

	ctx.Bind(&json)

	id := ctx.Param("id")
	// fmt.Println(id)
	slug := id
	title := json.Title
	date := json.Date
	url := json.Url
	username := json.Username
	description := json.Description
	content := json.Content
	except := json.Except
	commentStatus := json.CommentStatus
	status := json.Status
	ctype := json.Type
	postSurvey := json.PostSurvey
	// postParent := json.PostParent
	postType := json.PostType
	imageId := json.ImageId
	views := json.Views
	author := json.Author
	tags := json.Tags
	categories := json.Categories

	post := models.Post{
		Slug:          slug,
		Title:         title,
		Date:          date,
		Url:           url,
		Username:      username,
		Description:   description,
		Content:       content,
		Except:        except,
		CommentStatus: commentStatus,
		Status:        status,
		Type:          ctype,
		PostSurvey:    postSurvey,
		// PostParent:    postParent,
		PostType:   postType,
		ImageId:    imageId,
		Views:      views,
		Author:     author,
		Tags:       tags,
		Categories: categories,
	}

	update := bson.M{
		"$set": post,
	}

	filter := bson.M{"slug": id}
	result := models.Post{}
	err := collectionPosts.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)

	// fmt.Println(result)

	if err != nil {
		log.Fatal(err)
	} else {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "success",
			"data":    post,
			"slug":    result.Slug,
		})
	}

}

// DeletePost -> deletes post based on id
func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	// docID, _ := primitive.ObjectIDFromHex(id)
	result := models.Post{}
	filter := bson.M{"slug": id}
	err := collectionPosts.FindOneAndDelete(context.Background(), filter).Decode(&result)

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
