package main

import (
	"github.com/kidussolo/gocrud/app"
	"github.com/kidussolo/gocrud/models"
)

func main() {
	// Initialize db
	models.InitDb()

	// Make migrations
	models.DB.AutoMigrate(&models.Item{})

	// Add data to db
	models.DB.Create(&models.Item{Name: "iphone", Price: 1000.00})

	app.Start()
}
