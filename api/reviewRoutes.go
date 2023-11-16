package api

import (
	"laras/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func reviewRoutes(r *gin.Engine, db *gorm.DB, rc *controllers.ReviewController) {

	reviews := r.Group("/reviews")
	reviews.GET("", rc.GetReviews)
	reviews.GET("/:id", rc.GetReview)

}
