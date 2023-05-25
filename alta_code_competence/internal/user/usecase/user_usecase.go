package usecase

import (
	"errors"

	"github.com/alta_code_competence/internal/domain"
	"github.com/alta_code_competence/internal/user/repository"
)

type UserUsecase interface {
	Register(reqDTO domain.UserAuthRequest) (domain.User, error)
	Login(reqDTO domain.UserAuthRequest) (domain.User, error)
}

type userUsecase struct {
	rp repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) userUsecase {
	return userUsecase{
		rp: repo,
	}
}

func (u *userUsecase) Register(reqDTO domain.UserAuthRequest) (domain.User, error) {
	data := domain.User{
		Email:    reqDTO.Email,
		Password: reqDTO.Password,
	}
	return u.rp.Create(data)
}

func (u *userUsecase) Login(reqDTO domain.UserAuthRequest) (domain.User, error) {
	loginData := domain.User{
		Email:    reqDTO.Email,
		Password: reqDTO.Password,
	}

	data, err := u.rp.GetByEmail(loginData.Email)
	if err != nil {
		return domain.User{}, err
	}

	if loginData.Password != data.Password {
		return domain.User{}, errors.New("invalid email or password")
	}

	return data, nil
}