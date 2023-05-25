package repository

import (
	"github.com/alta_code_competence/internal/domain"
	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAll(queryDTO string) ([]domain.Item, error)
	GetAllByCategory(categoryId int) (domain.Category, error)
	Create(data domain.Item) (domain.Item, error)
	GetById(idDTO string) (domain.Item, error)
	Update(data domain.Item) (domain.Item, error)
	Delete(idDTO string) (domain.Item, error)
}

type psqlItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *psqlItemRepository {
	return &psqlItemRepository{
		DB: db,
	}
}

func (par *psqlItemRepository) GetAll(queryDTO string) ([]domain.Item, error) {
	data := []domain.Item{}
	query := par.DB.Debug().
		Model(&data).
		Preload("Category")
	
	if queryDTO != "" {
		queryParam := "%" + queryDTO + "%"
		query.Where("name LIKE ?", queryParam)
	}

	err := query.Find(&data).Error
	if err != nil {
		return []domain.Item{}, err
	}
	
	return data, nil
}

func (par *psqlItemRepository) GetAllByCategory(categoryId int) (domain.Category, error) {
	data := domain.Category{}
	// TODO: need optimization
	err := par.DB.Debug().
		Where("id = ?", categoryId).
		Preload("Items").
		Find(&data).Error
	if err != nil {
		return domain.Category{}, err
	}
	return data, nil
}

func (par *psqlItemRepository) Create(data domain.Item) (domain.Item, error) {
	err := par.DB.Preload("Category").Create(&data).Error
	if err != nil {
		return domain.Item{}, err
	}
	return data, nil
}

func (par *psqlItemRepository) GetById(idDTO string) (domain.Item, error) {
	data := domain.Item{}
	err := par.DB.Preload("Category").Where("id = ?", idDTO).First(&data).Error
	if err != nil {
		return domain.Item{}, err
	}
	return data, nil
}

func (par *psqlItemRepository) Update(data domain.Item) (domain.Item, error) {
	err := par.DB.Debug().Model(&data).Updates(data).Error
	if err != nil {
		return domain.Item{}, err
	}
	return data, nil
}

func (par *psqlItemRepository) Delete(idDTO string) (domain.Item, error) {
	data := domain.Item{
		Id: idDTO,
	}
	err := par.DB.Delete(&data).Error
	if err != nil {
		return domain.Item{}, err
	}
	return data, nil
}
