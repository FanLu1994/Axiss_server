package db

import "gorm.io/gorm"

type BenchMark struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
	Tags string `json:"tags"`
}
