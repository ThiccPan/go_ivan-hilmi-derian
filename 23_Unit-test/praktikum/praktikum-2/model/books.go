package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title" validate:"required"`
	Author    string `json:"author" form:"author" validate:"required"`
	Publisher string `json:"publisher" form:"publisher" validate:"required"`
}
