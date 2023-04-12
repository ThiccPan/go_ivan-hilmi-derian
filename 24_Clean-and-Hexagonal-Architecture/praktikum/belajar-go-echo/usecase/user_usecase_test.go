package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name      string
		data      dto.CreateUserRequest
		expectErr error
	}{
		{
			name: "success",
			data: dto.CreateUserRequest{
				Email:    "tes@gmail.com",
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

			data, err := service.Create(testCase.data)
			t.Log(data)
			t.Log(err)
			if testCase.expectErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, data.Email, testCase.data.Email)
				assert.Equal(t, data.Password, testCase.data.Password)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		name             string
		expectErr        error
		expectDataLength int
	}{
		{
			name:             "success",
			expectErr:        nil,
			expectDataLength: 1,
		},
	}

	mockData := model.User{
		Email:    "tes@email.com",
		Password: "123",
	}

	for _, testCase := range testCases {
		// init mock list user data to fetch
		mockDataList := []model.User{}
		for i := 0; i < int(testCase.expectDataLength); i++ {
			mockDataList = append(mockDataList, mockData)
		}

		mockUserRepository := repository.NewMockUserRepository()
		mockUserRepository.On("FetchAll").Return(mockDataList, testCase.expectErr)
		service := NewUserUsecase(mockUserRepository)

		data, err := service.GetAll()
		t.Log(data)
		t.Log(err)
		if testCase.expectErr != nil {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, testCase.expectDataLength ,len(mockDataList))
		}
	}
}

func TestLogin(t *testing.T) {

}
