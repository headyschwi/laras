package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Sinopsis    string `json:"sinopsis" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Rating      int    `json:"rating"`
	Pages       int    `json:"pages" binding:"required"`
	ReleaseDate string `json:"release_date" binding:"required"`

	Reviews []Review `gorm:"polymorphic:Media;" json:"reviews"`
}

type Movie struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Sinopsis    string `json:"sinopsis" binding:"required"`
	Director    string `json:"director" binding:"required"`
	Rating      int    `json:"rating"`
	ReleaseDate string `json:"release_date" binding:"required"`

	Reviews []Review `gorm:"polymorphic:Media;"`
}

type TVShow struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Sinopsis    string `json:"sinopsis" binding:"required"`
	Director    string `json:"director" binding:"required"`
	Rating      int    `json:"rating"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Episodes    int    `json:"episodes" binding:"required"`

	Reviews []Review `gorm:"polymorphic:Media;"`
}

type Manga struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Sinopsis    string `json:"sinopsis" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Rating      int    `json:"rating"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Chapters    int    `json:"chapters" binding:"required"`

	Reviews []Review `gorm:"polymorphic:Media;"`
}

type Anime struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Sinopsis    string `json:"sinopsis" binding:"required"`
	Director    string `json:"director" binding:"required"`
	Rating      int    `json:"rating"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Episodes    int    `json:"episodes" binding:"required"`

	Reviews []Review `gorm:"polymorphic:Media;"`
}

type Games struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Sinopsis    string `json:"sinopsis" binding:"required"`
	Developer   string `json:"developer" binding:"required"`
	Rating      int    `json:"rating"`
	ReleaseDate string `json:"release_date" binding:"required"`

	Reviews []Review `gorm:"polymorphic:Media;"`
}
