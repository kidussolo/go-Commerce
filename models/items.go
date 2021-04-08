package models

import "gorm.io/gorm"

// Item model
type Item struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
