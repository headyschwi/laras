package api

import (
	"fmt"
	"laras/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(port int, db *gorm.DB) {
	p := fmt.Sprintf(":%d", port)

	uc := controllers.NewUserController(db)
	mc := controllers.NewMediaController(db)
	rc := controllers.NewReviewController(db)
	pc := controllers.NewPostController(db)

	r := gin.Default()
	userRoutes(r, db, uc)
	mediaRoutes(r, db, mc, rc, uc)
	reviewRoutes(r, db, rc)
	postRoutes(r, db, pc)

	r.Run(p)
}
