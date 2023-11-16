package api

import (
	"laras/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func postRoutes(r *gin.Engine, db *gorm.DB, pc *controllers.PostController) {

	post := r.Group("/post")

	post.POST("/", pc.CreatePost)
	post.GET("/", pc.GetPosts)
	post.GET("/:id", pc.GetPostByID)
	post.PUT("/:id", pc.UpdatePost)
	post.DELETE("/:id", pc.DeletePost)

	post.POST("/:id/like", pc.LikePost)
}
