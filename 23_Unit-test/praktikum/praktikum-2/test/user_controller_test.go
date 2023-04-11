package test

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/controller"
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

type UsersResponse struct {
	Message string
	Users   []model.User
}

type UserResponse struct {
	Message string
	User    model.User
}

func InsertMockUserData() error {
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

	var testCases = []testCase{
		{
			name:       "(failed) wrong path name",
			path:       "/user",
			sizeData:   0,
			expectCode: http.StatusNotFound,
		},
		{
			name:       "(success) get list of books",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec, c := initTestEnv(req)
			t.Log(rec, c)

			for i := 0; i < testCase.sizeData; i++ {
				InsertMockUserData()
			}

			c.SetPath(testCase.path)
			err := controller.GetUsersController(c)
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()
			t.Log(body)

			var users UsersResponse
			err = json.Unmarshal([]byte(body), &users)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(users.Users))
		})
	}
}

func TestGetUserController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) get user with id",
			path:       "/users/:id",
			param:      "lorem",
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "(success) get user with id",
			path:       "/users/:id",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec, c := initTestEnv(req)
			InsertMockUserData()

			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues(testCase.param)
			err := controller.GetUserController(c)

			// check handler error
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()

			var user UserResponse
			err = json.Unmarshal([]byte(body), &user)
			t.Log(user.User)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.param, strconv.Itoa(int(user.User.ID)))
			}
		})
	}
}

func TestCreateUserController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) invalid req body",
			path:       "/users",
			body:       `{"lorem":"ipsum"}`,
			expectCode: http.StatusBadRequest,
		},
		{
			name:        "(success) create user",
			path:        "/users",
			body:        `{"name":"Alta","email":"alta@gmail.com","password":"123"}`,
			expectCode:  http.StatusOK,
			expectTitle: "Alta",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// create request, recorder and test env for each testcase because every req body is different
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec, c := initTestEnv(req)

			c.SetPath(testCase.path)
			err := controller.CreateUserController(c)

			// check handler error
			if err != nil {
				t.Log(err.Error())
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}
			body := rec.Body.String()

			var user UserResponse
			err = json.Unmarshal([]byte(body), &user)

			t.Log(user.Message)

			if err != nil {
				assert.Error(t, err, "error")
			} else {

				assert.Equal(t, testCase.expectTitle, user.User.Name)
			}
		})

	}
}

func TestUpdateUserController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) id not found",
			path:       "/users/:id",
			param:      "nil",
			body:       `{"name":"Alta Update","email":"altaupdate@gmail.com","password":"123"}`,
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "(success) update user",
			path:       "/users/:id",
			param:      "1",
			body:       `{"name":"Alta Update","email":"altaupdate@gmail.com","password":"123"}`,
			expectCode: http.StatusOK,
			expectTitle: "Alta Update",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// create request, recorder and test env for each testcase because every req body is different
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec, c := initTestEnv(req)

			err := InsertMockUserData()

			// check handler error
			if err != nil {
				t.Log(err.Error())
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues(testCase.param)

			controller.UpdateUserController(c)
			body := rec.Body.String()

			var user UserResponse
			err = json.Unmarshal([]byte(body), &user)
			t.Log(user)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.expectTitle, user.User.Name)
			}
		})
	}
}

func TestDeleteUserController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) id invalid",
			path:       "/users/:id",
			param:      "lorem",
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "(success) delete book with id",
			path:       "/users/:id",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			rec, c := initTestEnv(req)
			InsertMockUserData()

			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues(testCase.param)
			err := controller.DeleteUserController(c)

			// check handler error
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()

			var user UserResponse
			err = json.Unmarshal([]byte(body), &user)
			t.Log(user.User)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.param, strconv.Itoa(int(user.User.ID)))
			}
		})
	}
}
