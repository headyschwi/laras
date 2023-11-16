package main

import (
	"laras/api"
	"laras/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := database_connect()
	db.Migrator().DropTable(&models.User{}, &models.Review{}, &models.Status{}, &models.Book{}, &models.Movie{})
	db.AutoMigrate(&models.User{}, &models.Review{}, &models.Status{}, &models.Book{}, &models.Movie{})

	api.Run(2612, db)
}

func database_connect() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/laras?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
