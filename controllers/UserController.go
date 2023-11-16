package controllers

import (
	"fmt"
	"laras/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := uc.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error hashing password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) Login(c *gin.Context) {

}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(200, users)
}

func (uc *UserController) GetUserById(c *gin.Context) {
	var user models.User

	if err := uc.DB.Preload("Reviews").Preload("Status").Preload("Posts").Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error user not found"})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {

	var user models.User
	if err := uc.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error user not found"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.ID != updatedUser.ID {
		c.JSON(400, gin.H{"error": "Id could not be changed"})
		return
	}

	uc.DB.Save(&updatedUser)
	c.JSON(200, updatedUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {

	var user models.User
	if err := uc.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error user not found"})
		return
	}

	uc.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted"})
}

func (uc *UserController) GetUserSave(c *gin.Context) {
	var userMedia []models.Status
	if err := uc.DB.Where("user_id = ?", c.Param("id")).Find(&userMedia).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error fetching user media"})
		return
	}

	c.JSON(200, userMedia)
}

func (uc *UserController) UpdateUserSave(c *gin.Context) {
	var userMedia models.Status
	if err := c.ShouldBindJSON(&userMedia); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := userMedia.ValidateStatus(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if fmt.Sprint(userMedia.MediaID) != c.Param("id") {
		c.JSON(400, gin.H{"error": "Media id not match"})
	}

	if err := uc.DB.Save(&userMedia).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error updating user media"})
		return
	}

	c.JSON(200, userMedia)
}

func (uc *UserController) DeleteUserSave(c *gin.Context) {
	var userMedia models.Status
	if err := uc.DB.Where("user_id = ? AND id = ?", c.Param("id"), c.Param("id_save")).First(&userMedia).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error fetching user media"})
		return
	}

	uc.DB.Delete(&userMedia)
	c.JSON(200, gin.H{"message": "User media deleted"})
}

func (uc *UserController) CreateStatus(c *gin.Context) {
	var mediaStatus models.Status
	if err := c.ShouldBindJSON(&mediaStatus); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := mediaStatus.ValidateStatus(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if fmt.Sprint(mediaStatus.MediaID) != c.Param("id") && mediaStatus.MediaType != c.Param("content_type") {
		c.JSON(400, gin.H{"error": "Media id not match"})
	}

	if err := uc.DB.Create(&mediaStatus).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating user media"})
		return
	}

	c.JSON(200, mediaStatus)
}
