package model

import "gorm.io/gorm"

// Product struct
type Post struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
}
