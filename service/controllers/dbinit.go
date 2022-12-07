package api

import (
	"os"
	"strconv"

	"alfa-indo-soft/service/db"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

// initializes single database collection object for all the api requests
var collectionBlogs *mongo.Collection
var collectionUsers *mongo.Collection
var collectionPosts *mongo.Collection
var collectionCategories *mongo.Collection
var collectionTags *mongo.Collection
var jwtKey []byte
var costI int

func init() {
	godotenv.Load()
	jj := os.Getenv("JWT_KEY")
	cost := os.Getenv("COST")

	costI, _ = strconv.Atoi(cost)
	jwtKey = []byte(jj)
	collectionBlogs = db.ConnectClient().Database("chatnews_cron").Collection("blogs")
	collectionUsers = db.ConnectClient().Database("chatnews_cron").Collection("users")
	collectionPosts = db.ConnectClient().Database("chatnews_cron").Collection("posts")
	collectionCategories = db.ConnectClient().Database("chatnews_cron").Collection("categories")
	collectionTags = db.ConnectClient().Database("chatnews_cron").Collection("tags")
}
