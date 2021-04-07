package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Load environment variables
func Loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error Loading .env file")
	}
}
