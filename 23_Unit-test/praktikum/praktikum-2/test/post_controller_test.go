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

type PostResponses struct {
	Message string
	Posts   []model.Post
}

type PostResponse struct {
	Message string
	Post    model.Post
}

func InsertMockPostData() error {
	InsertMockUserData()
	user := model.Post{
		User_id: 1,
		Title:   "Post test",
		Content: "lorem ipsum",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetPostsController(t *testing.T) {

	var testCases = []testCase{
		{
			name:       "(failed) wrong path name",
			path:       "/post",
			sizeData:   0,
			expectCode: http.StatusNotFound,
		},
		{
			name:       "(success) get list of books",
			path:       "/posts",
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
				InsertMockPostData()
			}

			c.SetPath(testCase.path)
			err := controller.GetPostsController(c)
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()
			t.Log(body)

			var posts PostResponses
			err = json.Unmarshal([]byte(body), &posts)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(posts.Posts))
		})
	}
}

func TestGetPostController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) invalid param",
			path:       "/posts/:id",
			param:      "lorem",
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "(success) get user with id",
			path:       "/posts/:id",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec, c := initTestEnv(req)
			InsertMockPostData()

			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues(testCase.param)
			err := controller.GetPostController(c)

			// check handler error
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()

			var post PostResponse
			err = json.Unmarshal([]byte(body), &post)
			t.Log(post.Post)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.param, strconv.Itoa(int(post.Post.ID)))
			}
		})
	}
}

func TestCreatePostController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) invalid req body",
			path:       "/posts",
			body:       `{"lorem":"ipsum"}`,
			expectCode: http.StatusBadRequest,
		},
		{
			name:        "(success) create post",
			path:        "/posts",
			body:        `{"user_id":1,"title":"post test","content":"lorem ipsum"}`,
			expectCode:  http.StatusOK,
			expectTitle: "post test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// create request, recorder and test env for each testcase because every req body is different
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec, c := initTestEnv(req)
			InsertMockUserData()
			InsertMockUserData()
			t.Log(req.Body)

			c.SetPath(testCase.path)
			err := controller.CreatePostController(c)

			// check handler error
			if err != nil {
				t.Log(err.Error())
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}
			body := rec.Body.String()

			var post PostResponse
			err = json.Unmarshal([]byte(body), &post)

			t.Log(post.Message)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.expectTitle, post.Post.Title)
			}
		})

	}
}

func TestUpdatePostController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) invalid param",
			path:       "/posts/:id",
			param:      "lorem",
			expectCode: http.StatusBadRequest,
		},
		{
			name:        "(success) get user with id",
			path:        "/posts/:id",
			param:       "1",
			body:        `{"user_id":1,"title":"post test update","content":"lorem ipsum"}`,
			expectCode:  http.StatusOK,
			expectTitle: "post test update",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// create request, recorder and test env for each testcase because every req body is different
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec, c := initTestEnv(req)

			err := InsertMockPostData()

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

			controller.UpdatePostController(c)
			body := rec.Body.String()

			var post PostResponse
			err = json.Unmarshal([]byte(body), &post)
			t.Log(post)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.expectTitle, post.Post.Title)
			}
		})
	}
}

func TestDeletePostController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) id invalid",
			path:       "/posts/:id",
			param:      "lorem",
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "(success) delete book with id",
			path:       "/posts/:id",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			rec, c := initTestEnv(req)
			InsertMockPostData()

			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues(testCase.param)
			err := controller.DeletePostController(c)

			// check handler error
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()
			t.Log(body)

			var post PostResponse
			err = json.Unmarshal([]byte(body), &post)
			t.Log(post.Post)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.param, strconv.Itoa(int(post.Post.ID)))
			}
		})
	}
}
