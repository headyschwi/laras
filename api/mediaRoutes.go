package api

import (
	"laras/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func mediaRoutes(r *gin.Engine, db *gorm.DB, mc *controllers.MediaController, rc *controllers.ReviewController, uc *controllers.UserController) {

	media := r.Group("/:media_type")
	media.Use(checkMediaType())

	media.POST("", mc.CreateMedia)
	media.GET("", mc.GetMedias)
	media.GET("/:id", mc.GetMedia)
	media.PUT("/:id", mc.UpdateMedia)
	media.DELETE("/:id", mc.DeleteMedia)

	media.POST("/:id/review", rc.CreateReview)
	media.GET("/:id/reviews", rc.GetReviewsByMedia)

	media.POST("/:id/save", uc.CreateStatus)
}

func checkMediaType() gin.HandlerFunc {
	return func(c *gin.Context) {
		Type := c.Param("media_type")
		if Type != "book" && Type != "movie" && Type != "tvshow" {
			c.JSON(400, gin.H{"error": "Invalid media type"})
			c.Abort()
			return
		}
		c.Set("type", Type)
		c.Next()
	}
}
