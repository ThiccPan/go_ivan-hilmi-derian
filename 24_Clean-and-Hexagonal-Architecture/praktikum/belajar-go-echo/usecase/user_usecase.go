package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	Create(payloads dto.CreateUserRequest) (*model.User, error)
	GetAll() ([]model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) Create(payloads dto.CreateUserRequest) (*model.User, error) {
	user := model.User{
		Email: payloads.Email,
		Password: payloads.Password,
	}

	err := u.userRepository.Create(user)
	return &user, err
}

func (u *userUsecase) GetAll() ([]model.User, error) {
	return u.userRepository.FetchAll()
}
