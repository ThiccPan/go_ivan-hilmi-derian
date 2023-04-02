package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	User_id uint
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}


