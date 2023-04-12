package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/repository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testCases := []struct{
		name string
		data dto.CreateUserRequest
		expectErr error
	}{
		{
			name: "success",
			data: dto.CreateUserRequest{
				Email: "tes@gmail.com",
				Password: "123",
			},
			expectErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockUserRepository := repository.NewMockUserRepository()
			mockUserRepository.On("Create", mock.Anything).Return(testCase.expectErr)
			service := NewUserUsecase(mockUserRepository)
			// t.Log(service)
	
			user, err := service.Create(testCase.data);
			t.Log(user)
			t.Log(err)
			if testCase.expectErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, user.Email, testCase.data.Email)
				assert.Equal(t, user.Password, testCase.data.Password)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	
}

func TestLogin(t *testing.T) {
	
}