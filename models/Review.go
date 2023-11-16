package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID    uint   `json:"user_id" binding:"required"`
	MediaID   uint   `json:"media_id" binding:"required"`
	MediaType string `json:"media_type" binding:"required"`
	Content   string `json:"content"`
	Rating    int    `json:"rating" binding:"required"`

	UpVotes []UpVote `gorm:"foreignKey:ReviewID"`
}

type UpVote struct {
	gorm.Model
	UserID   uint `json:"user_id" binding:"required"`
	ReviewID uint `json:"review_id" binding:"required"`
}
