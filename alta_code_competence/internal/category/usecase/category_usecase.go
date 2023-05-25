package usecase

import (
	"github.com/alta_code_competence/internal/domain"
	"github.com/alta_code_competence/internal/category/repository"
)

type CategoryUsecase interface {
	GetAll() ([]domain.Category, error)
}

type categoryUsecase struct {
	rp repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) categoryUsecase {
	return categoryUsecase{
		rp: repo,
	}
}

func (iu *categoryUsecase) GetAll() ([]domain.Category, error) {
	return iu.rp.GetAll()
}