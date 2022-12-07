package service

import (
	api "alfa-indo-soft/service/controllers"

	"github.com/gin-gonic/gin"
)

// ExtRouter Custom Router
func ExtRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.Default()

	// route handling basic
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ahlan wa sahlan alfa-indo-soft",
		})
	})

	router.GET("/article", api.GetAllArticles)
	router.GET("/article_by_title_and_body", api.GetArticleByTitleAndBody)
	router.GET("/article_by_author/:author", api.GetArticleByAuthor)
	router.POST("/article", api.InsertArticle)
	router.PUT("/article/:id", api.UpdateArticle)
	router.DELETE("/article/:id", api.DeleteArticle)

	// route handling for external
	router.GET("/ping", api.Ping)
	router.GET("/blog", api.GetAllBlogs)
	router.GET("/blog/:id", api.GetBlogs)

	// routes only accesiible if logged in
	authonly := router.Group("/")
	authonly.Use(api.LoginMiddleware())
	{
		authonly.POST("/blog", api.InsertBlog)
		authonly.PUT("/blog/:id", api.UpdateBlog)
		authonly.DELETE("/blog/:id", api.DeleteBlog)

		authonly.GET("/category", api.GetAllCategories)
		authonly.GET("/category_by_slug/:id", api.GetCategoryBySlug)
		authonly.POST("/category", api.InsertCategory)
		authonly.PUT("/category/:id", api.UpdateCategory)
		authonly.DELETE("/category/:id", api.DeleteCategory)

		authonly.GET("/post", api.GetAllPosts)
		authonly.GET("/post_by_slug/:id", api.GetPostBySlug)
		authonly.POST("/post", api.InsertPost)
		authonly.PUT("/post/:id", api.UpdatePost)
		authonly.DELETE("/post/:id", api.DeletePost)

		authonly.GET("/tag", api.GetAllTags)
		authonly.GET("/tag_by_slug/:id", api.GetTagBySlug)
		authonly.POST("/tag", api.InsertTag)
		authonly.PUT("/tag/:id", api.UpdateTag)
		authonly.DELETE("/tag/:id", api.DeleteTag)

		authonly.GET("/user", api.GetAllUsers)
	}

	//auth routes
	authGroup := router.Group("/auth")
	authGroup.POST("/signup", api.Signup)
	authGroup.POST("/login", api.Login)
	authGroup.GET("/refresh", api.Refresh)

	return router
}
