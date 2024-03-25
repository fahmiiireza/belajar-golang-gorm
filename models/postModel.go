package models

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	Title string `json:"title"`
	Post  string `json:"post"`
}
