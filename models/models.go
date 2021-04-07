package models

import (
	"fmt"
	"os"

	"github.com/kidussolo/gocrud/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize the db
func InitDb() {
	utils.Loadenv()
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		port     = 5432
		dbname   = os.Getenv("DB_NAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db")
	DB = db
}
