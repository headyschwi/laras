package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint   `json:"user_id" binding:"required"`
	Content string `json:"content"`

	ReviewID uint `json:"review_id" default:"0"`
	StatusID uint `json:"status_id" default:"0"`
	ParentID uint `json:"parent_id" default:"0"`
	RepostID uint `json:"repost_id" default:"0"`

	Likes []PostLike `gorm:"foreignKey:PostID"`
}

type PostLike struct {
	gorm.Model
	UserID uint `json:"user_id" binding:"required"`
	PostID uint `json:"post_id" binding:"required"`
}
