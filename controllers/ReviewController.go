package controllers

import (
	"laras/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReviewController struct {
	DB *gorm.DB
}

func NewReviewController(db *gorm.DB) *ReviewController {
	return &ReviewController{DB: db}
}

func (rc *ReviewController) CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingReview models.Review
	if err := rc.DB.Where("user_id = ? AND media_id = ? AND media_type = ?", review.UserID, review.MediaID, review.MediaType).First(&existingReview).Error; err == nil {
		c.JSON(400, gin.H{"error": "Review already exists"})
		return
	}

	if err := rc.DB.Create(&review).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating review"})
		return
	}

	c.JSON(200, review)
}

func (rc *ReviewController) GetReviews(c *gin.Context) {
	var reviews []models.Review

	if err := rc.DB.Find(&reviews).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting reviews"})
		return
	}

	c.JSON(200, reviews)
}

func (rc *ReviewController) GetReview(c *gin.Context) {
	var review models.Review

	if err := rc.DB.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting review"})
		return
	}

	c.JSON(200, review)
}

func (rc *ReviewController) UpdateReview(c *gin.Context) {
	var review models.Review

	if err := rc.DB.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting review"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Save(&review).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error updating review"})
		return
	}

	c.JSON(200, review)
}

func (rc *ReviewController) DeleteReview(c *gin.Context) {
	var review models.Review

	if err := rc.DB.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting review"})
		return
	}

	if err := rc.DB.Delete(&review).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error deleting review"})
		return
	}

	c.JSON(200, gin.H{"message": "Review deleted"})
}

func (rc *ReviewController) UpVote(c *gin.Context) {
	var upVote models.UpVote

	if err := c.ShouldBindJSON(&upVote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUpVote models.UpVote
	if err := rc.DB.Where("user_id = ? AND review_id = ?", upVote.UserID, upVote.ReviewID).First(&existingUpVote).Error; err == nil {
		c.JSON(400, gin.H{"error": "Upvote already exists"})
		return
	}

	if err := rc.DB.Create(&upVote).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating upvote"})
		return
	}

	c.JSON(200, upVote)
}

func (rc *ReviewController) RemoveUpVote(c *gin.Context) {

	var upVote models.UpVote

	if err := rc.DB.Where("id = ?", c.Param("id")).First(&upVote).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting upvote"})
		return
	}

	if err := rc.DB.Delete(&upVote).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error deleting upvote"})
		return
	}

	c.JSON(200, gin.H{"message": "Upvote removed"})
}

func (rc *ReviewController) GetReviewsByUser(c *gin.Context) {
	var reviews []models.Review

	if err := rc.DB.Where("user_id = ?", c.Param("id")).Find(&reviews).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting reviews"})
		return
	}

	c.JSON(200, reviews)
}

func (rc *ReviewController) GetReviewsByMedia(c *gin.Context) {
	var reviews []models.Review

	if err := rc.DB.Where("media_id = ? AND media_type = ?", c.Param("id"), c.GetString("type")).Find(&reviews).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error getting reviews"})
		return
	}

	c.JSON(200, reviews)
}
