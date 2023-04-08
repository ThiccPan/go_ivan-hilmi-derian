package controller

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/model"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	path       string
	param      string
	body       string
	expectCode int
	sizeData   int
}

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

func initTestEnv(req *http.Request) (*httptest.ResponseRecorder, echo.Context) {
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

func TestGetUsersController(t *testing.T) {
	testCases := []testCase{
		{
			name:       "(success) get list of user",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec, c := initTestEnv(req)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

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
	var testCases = []testCase{
		{
			name:       "(success) get user with id",
			path:       "/users",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec, c := initTestEnv(req)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.param)
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

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

func TestCreateUserController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "create user",
			path:       "/users",
			body:       `{"name":"Alta","email":"alta@gmail.com","password":"123"}`,
			expectCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {

		// create request, recorder and test env for each testcase because every req body is different
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec, c := initTestEnv(req)

		c.SetPath(testCase.path)
		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()
			t.Log(req.Body.Close())

			var user UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Alta", user.User.Name)
		}
	}
}
