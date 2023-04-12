package repository

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data model.User) error
	FetchAll() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) Create(data model.User) error {
	return u.db.Create(data).Error
}

func (u *userRepository) FetchAll() ([]model.User, error) {
	user := []model.User{}
	err := u.db.Find(&user).Error
	return user, err
}