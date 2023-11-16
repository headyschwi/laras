package controllers

import (
	"laras/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MediaController struct {
	DB *gorm.DB
}

func NewMediaController(db *gorm.DB) *MediaController {
	return &MediaController{DB: db}
}

func (mc *MediaController) CreateMedia(c *gin.Context) {

	if c.GetString("type") == "book" {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingBook models.Book
		if err := mc.DB.Where("isbn = ?", book.ISBN).First(&existingBook).Error; err == nil {
			c.JSON(400, gin.H{"error": "Book already exists"})
			return
		}

		mc.DB.Create(&book)
		c.JSON(200, book)
		return
	} else if c.GetString("type") == "movie" {
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingMovie models.Movie
		if err := mc.DB.Where("title = ?", movie.Title).First(&existingMovie).Error; err == nil {
			c.JSON(400, gin.H{"error": "Movie already exists"})
			return
		}

		mc.DB.Create(&movie)
		c.JSON(200, movie)
		return
	} else if c.GetString("type") == "tvshow" {
		var tvshow models.TVShow
		if err := c.ShouldBindJSON(&tvshow); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingTVShow models.TVShow
		if err := mc.DB.Where("title = ?", tvshow.Title).First(&existingTVShow).Error; err == nil {
			c.JSON(400, gin.H{"error": "TV Show already exists"})
			return
		}

		mc.DB.Create(&tvshow)
		c.JSON(200, tvshow)
		return
	} else if c.GetString("type") == "manga" {
		var manga models.Manga
		if err := c.ShouldBindJSON(&manga); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingManga models.Manga
		if err := mc.DB.Where("title = ?", manga.Title).First(&existingManga).Error; err == nil {
			c.JSON(400, gin.H{"error": "Manga already exists"})
			return
		}

		mc.DB.Create(&manga)
		c.JSON(200, manga)
		return
	} else if c.GetString("type") == "anime" {
		var anime models.Anime
		if err := c.ShouldBindJSON(&anime); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingAnime models.Anime
		if err := mc.DB.Where("title = ?", anime.Title).First(&existingAnime).Error; err == nil {
			c.JSON(400, gin.H{"error": "Anime already exists"})
			return
		}

		mc.DB.Create(&anime)
		c.JSON(200, anime)
		return
	}

}

func (mc *MediaController) GetMedia(c *gin.Context) {
	if c.GetString("type") == "book" {
		var book models.Book
		if err := mc.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
			c.JSON(400, gin.H{"error": "Book does not exist"})
			return
		}
		c.JSON(200, book)
		return
	} else if c.GetString("type") == "movie" {
		var movie models.Movie
		if err := mc.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
			c.JSON(400, gin.H{"error": "Movie does not exist"})
			return
		}
		c.JSON(200, movie)
		return
	} else if c.GetString("type") == "tvshow" {
		var tvshow models.TVShow
		if err := mc.DB.Where("id = ?", c.Param("id")).First(&tvshow).Error; err != nil {
			c.JSON(400, gin.H{"error": "TV Show does not exist"})
			return
		}
		c.JSON(200, tvshow)
		return
	}
}

func (mc *MediaController) GetMedias(c *gin.Context) {

	if c.GetString("type") == "book" {
		var books []models.Book
		mc.DB.Find(&books)
		c.JSON(200, books)
		return
	} else if c.GetString("type") == "movie" {
		var movies []models.Movie
		mc.DB.Find(&movies)
		c.JSON(200, movies)
		return
	} else if c.GetString("type") == "tvshow" {
		var tvshows []models.TVShow
		mc.DB.Find(&tvshows)
		c.JSON(200, tvshows)
		return
	}
}

func (mc *MediaController) UpdateMedia(c *gin.Context) {

	if c.GetString("type") == "book" {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingBook models.Book
		if err := mc.DB.Where("id = ?", c.Param("id")).First(&existingBook).Error; err != nil {
			c.JSON(400, gin.H{"error": "Book does not exist"})
			return
		}

		if book.ISBN != existingBook.ISBN {
			var existingBook models.Book
			if err := mc.DB.Where("isbn = ?", book.ISBN).First(&existingBook).Error; err == nil {
				c.JSON(400, gin.H{"error": "There is already a book with this ISBN and it's not the one you're trying to update."})
			}
		}

		mc.DB.Model(&existingBook).Updates(book)
		c.JSON(200, existingBook)
		return
	} else if c.GetString("type") == "movie" {
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingMovie models.Movie
		if err := mc.DB.Where("id = ?", c.Param("id")).First(&existingMovie).Error; err != nil {
			c.JSON(400, gin.H{"error": "Movie does not exist"})
			return
		}

		if movie.Title != existingMovie.Title {
			var existingMovie models.Movie
			if err := mc.DB.Where("title = ?", movie.Title).First(&existingMovie).Error; err == nil {
				c.JSON(400, gin.H{"error": "There is already a movie with this title and it's not the one you're trying to update."})
			}
		}

		mc.DB.Model(&existingMovie).Updates(movie)
		c.JSON(200, existingMovie)
		return
	} else if c.GetString("type") == "tvshow" {
		var tvshow models.TVShow
		if err := c.ShouldBindJSON(&tvshow); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var existingTVShow models.TVShow
		if err := mc.DB.Where("id = ?", c.Param("id")).First(&existingTVShow).Error; err != nil {
			c.JSON(400, gin.H{"error": "TV Show does not exist"})
			return
		}

		if tvshow.Title != existingTVShow.Title {
			var existingTVShow models.TVShow
			if err := mc.DB.Where("title = ?", tvshow.Title).First(&existingTVShow).Error; err == nil {
				c.JSON(400, gin.H{"error": "There is already a TV Show with this title and it's not the one you're trying to update."})
			}
		}

		mc.DB.Model(&existingTVShow).Updates(tvshow)
		c.JSON(200, existingTVShow)
		return
	}
}

func (mc *MediaController) DeleteMedia(c *gin.Context) {
	if c.GetString("type") == "book" {

		var existingBook models.Book
		if err := mc.DB.Where("id ?", c.Param("id")).First(&existingBook).Error; err != nil {
			c.JSON(400, gin.H{"error": "Book does not exist"})
			return
		}

		mc.DB.Delete(&existingBook)
		c.JSON(200, "Book deleted")
	} else if c.GetString("type") == "movie" {
		var movie models.Movie

		if err := mc.DB.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
			c.JSON(400, gin.H{"error": "Movie does not exist"})
			return
		}

		mc.DB.Delete(&movie)
		c.JSON(200, "Movie deleted")
	} else if c.GetString("type") == "tvshow" {
		var tvshow models.TVShow

		if err := mc.DB.Where("id = ?", c.Param("id")).First(&tvshow).Error; err != nil {
			c.JSON(400, gin.H{"error": "TV Show does not exist"})
			return
		}

		mc.DB.Delete(&tvshow)
		c.JSON(200, "TV Show deleted")
	}
}
