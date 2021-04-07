package models

import (
	"gorm.io/gorm"
)

// Item model
type Item struct {
	gorm.Model
	Name  string
	Price float64
}
