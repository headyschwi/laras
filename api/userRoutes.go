package api

import (
	"laras/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func userRoutes(r *gin.Engine, db *gorm.DB, uc *controllers.UserController) {
	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)

	user := r.Group("/users")
	user.GET("", uc.GetUsers)
	user.GET("/:id", uc.GetUserById)
	user.PUT("/:id", uc.UpdateUser)
	user.DELETE("/:id", uc.DeleteUser)

	user.GET("/:id/saves", uc.GetUserSave)
	user.PUT("/:id/:id_save/update", uc.UpdateUserSave)
	user.DELETE("/:id/:id_save/remove", uc.DeleteUserSave)
}
