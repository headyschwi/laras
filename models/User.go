package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`

	Posts   []Post   `gorm:"foreignKey:UserID"`
	Reviews []Review `gorm:"foreignKey:UserID"`
	Status  []Status `gorm:"foreignKey:UserID"`
}

type UserMediaStatus string

const (
	StatusInProgress UserMediaStatus = "in_progress"
	StatusCompleted  UserMediaStatus = "completed"
	StatusPaused     UserMediaStatus = "paused"
	StatusDropped    UserMediaStatus = "dropped"
	StatusPlansTo    UserMediaStatus = "plans_to"
)

type Status struct {
	gorm.Model
	UserID    uint            `json:"user_id" binding:"required"`
	MediaID   uint            `json:"media_id" binding:"required"`
	MediaType string          `json:"media_type" binding:"required"`
	Status    UserMediaStatus `json:"status" binding:"required"`
	Progress  float64         `json:"progress"`
}

func (um *Status) ValidateStatus() error {
	switch um.Status {
	case StatusInProgress, StatusCompleted, StatusPaused:
		return nil
	default:
		return errors.New("Invalid status")
	}
}
