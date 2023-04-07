package controller

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type UsersResponse struct {
	Message string
	Users   []model.User
}

type UserResponse struct {
	Message string
	User    model.User
}

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			var user UsersResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(user.Users))
		}
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		param      string
		expectCode int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.param)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			var user UserResponse
			err := json.Unmarshal([]byte(body), &user)
			t.Log(user.User)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.param, strconv.Itoa(int(user.User.ID)))
		}
	}
}
