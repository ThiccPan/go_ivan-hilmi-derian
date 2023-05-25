package usecase

import (
	"github.com/alta_code_competence/internal/domain"
	"github.com/alta_code_competence/helper"
	"github.com/alta_code_competence/internal/item/repository"
)

type ItemUsecase interface {
	GetAll(queryDTO string) ([]domain.Item, error)
	GetAllByCategory(categoryId int) (domain.Category, error)
	Create(requestDTO domain.ItemCreateRequest) (domain.Item, error)
	GetById(idDTO string) (domain.Item, error)
	Update(reqDTO domain.ItemEditRequest) (domain.Item, error)
	Delete(idDTO string) (domain.Item, error)
}

type itemUsecase struct {
	rp repository.ItemRepository
	uuidGenerator helper.UuidGenerator
}

func NewItemUsecase(repo repository.ItemRepository, ug helper.UuidGenerator) itemUsecase {
	return itemUsecase{
		rp: repo,
		uuidGenerator: ug,
	}
}

func (iu *itemUsecase) GetAll(queryDTO string) ([]domain.Item, error) {
	return iu.rp.GetAll(queryDTO)
}

func (iu *itemUsecase) GetAllByCategory(categoryId int) (domain.Category, error) {
	return iu.rp.GetAllByCategory(categoryId)
}

func (iu *itemUsecase) Create(requestDTO domain.ItemCreateRequest) (domain.Item, error) {
	generatedId, err := iu.uuidGenerator.GenerateUUID()
	if err != nil {
		return domain.Item{}, err
	}

	itemData := domain.Item{
		Id: generatedId,
		Name: requestDTO.Name,
		Description: requestDTO.Description,
		CategoryID: requestDTO.CategoryID,
		Price: requestDTO.Price,
		Stock: requestDTO.Stock,
	}
	
	data, err := iu.rp.Create(itemData)
	if err != nil {
		return domain.Item{}, err
	}
	
	return data, nil
}

func (iu *itemUsecase) GetById(idDTO string) (domain.Item, error) {
	return iu.rp.GetById(idDTO)
}

func (iu *itemUsecase) Update(reqDTO domain.ItemEditRequest) (domain.Item, error) {
	itemData := domain.Item{
		Id: reqDTO.Id,
		Name: reqDTO.Name,
		Description: reqDTO.Description,
		CategoryID: reqDTO.CategoryID,
		Price: reqDTO.Price,
		Stock: reqDTO.Stock,
	}
	return iu.rp.Update(itemData)
}

func (iu *itemUsecase) Delete(idDTO string) (domain.Item, error) {
	return iu.rp.Delete(idDTO)
}