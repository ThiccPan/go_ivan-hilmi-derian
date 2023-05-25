package repository

import (
	"github.com/alta_code_competence/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data domain.User) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
}

type userRepository struct {
	DB gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: *db,
	}
}

func (u *userRepository) Create(data domain.User) (domain.User, error) {
	err := u.DB.Create(&data).Error
	if err != nil {
		return domain.User{}, err
	}
	return data, nil
}

func (u *userRepository) GetByEmail(email string) (domain.User, error) {
	data := domain.User{
		Email: email,
	}
	
	err := u.DB.Debug().First(&data).Error
	if err != nil {
		return domain.User{}, err
	}

	return data, nil
}