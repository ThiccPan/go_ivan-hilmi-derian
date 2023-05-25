package repository

import (
	"github.com/alta_code_competence/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]domain.Category, error)
}

type psqlCategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *psqlCategoryRepository {
	return &psqlCategoryRepository{
		DB: db,
	}
}

func (par *psqlCategoryRepository) GetAll() ([]domain.Category, error) {
	data := []domain.Category{}
	err := par.DB.Find(&data).Error
	if err != nil {
		return []domain.Category{}, err
	}
	return data, nil
}

