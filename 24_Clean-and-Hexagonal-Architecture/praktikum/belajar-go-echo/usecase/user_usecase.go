package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"belajar-go-echo/middleware"
)

type UserUsecase interface {
	Create(payloads dto.CreateUserRequest) (*model.User, error)
	GetAll() ([]model.User, error)
	Login(payload dto.LoginUserRequest) (string, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	authMethod middleware.AuthJWT
}

func NewUserUsecase(userRepo repository.UserRepository, authMeth middleware.AuthJWT) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
		authMethod: authMeth,
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

func (u *userUsecase) Login(payload dto.LoginUserRequest) (string, error) {
	user, err := u.userRepository.FetchByEmail(payload.Email)
	if err != nil {
		return "", err
	}

	if payload.Password != user.Password {
		return "", err
	}
	
	t, mErr := u.authMethod.GenerateToken(user.Email)
	if mErr != nil {
		return "", err
	}

	return t, err
}
