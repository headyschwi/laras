package controllers

import (
	"laras/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if post.ReviewID != 0 {
		var existingReview models.Review
		if err := pc.DB.Where("id = ?", post.ReviewID).First(&existingReview).Error; err != nil {
			c.JSON(400, gin.H{"error": "Review does not exist"})
			return
		}
	}

	if post.StatusID != 0 {
		var existingStatus models.Status
		if err := pc.DB.Where("id = ?", post.StatusID).First(&existingStatus).Error; err != nil {
			c.JSON(400, gin.H{"error": "Update does not exist"})
			return
		}
	}

	if post.ParentID != 0 {
		var existingPost models.Post
		if err := pc.DB.Where("id = ?", post.ParentID).First(&existingPost).Error; err != nil {
			c.JSON(400, gin.H{"error": "Parent post does not exist"})
			return
		}
	}

	if post.RepostID != 0 {
		var existingPost models.Post
		if err := pc.DB.Where("id = ?", post.RepostID).First(&existingPost).Error; err != nil {
			c.JSON(400, gin.H{"error": "Post to repost does not exist"})
			return
		}
	}

	if err := pc.DB.Create(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating post"})
		return
	}

	c.JSON(200, post)
}

func (pc *PostController) GetPosts(c *gin.Context) {

	var posts []models.Post
	if err := pc.DB.Find(&posts).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error fetching posts"})
		return
	}

	c.JSON(200, posts)
}

func (pc *PostController) GetPostByID(c *gin.Context) {

	var post models.Post

	if err := pc.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error post not found"})
		return
	}

	c.JSON(200, post)
}

func (pc *PostController) UpdatePost(c *gin.Context) {

	var inputPost struct {
		Content string `json:"content"`
		PostID  uint   `json:"post_id"`
	}

	if err := c.ShouldBindJSON(&inputPost); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := pc.DB.Where("id = ?", inputPost.PostID).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error post not found"})
		return
	}

	post.Content = inputPost.Content

	pc.DB.Save(&post)
	c.JSON(200, post)
}

func (pc *PostController) DeletePost(c *gin.Context) {

	var post models.Post
	if err := pc.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error post not found"})
		return
	}

	pc.DB.Delete(&post)
	c.JSON(200, post)
}

func (pc *PostController) LikePost(c *gin.Context) {

	var like models.PostLike
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := pc.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error post not found"})
		return
	}

	if like.PostID != post.ID {
		c.JSON(400, gin.H{"error": "Post id does not match"})
		return
	}

	if err := pc.DB.Create(&like).Error; err != nil {
		c.JSON(400, gin.H{"error": "Error creating like"})
		return
	}

	c.JSON(200, like)
}
