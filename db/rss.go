package db

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
