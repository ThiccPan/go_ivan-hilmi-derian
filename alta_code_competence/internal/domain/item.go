package domain

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"categoryID"`
	Category Category `json:"category"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TODO: validate req param
type ItemCreateRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"categoryID"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
}

type ItemGetResponse struct {
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"categoryID"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
}

type ItemEditRequest struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"categoryID"`
	Price uint `json:"price"`
	Stock uint `json:"stock"`
}