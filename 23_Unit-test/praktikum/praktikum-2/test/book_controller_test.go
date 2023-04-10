package test

import (
	"22_Middleware/praktikum/model"
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/controller"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type BooksResponse struct {
	Message string
	Books   []model.Book
}

type BookResponse struct {
	Message string
	Book    model.Book
}

func InsertMockBookData() error {
	book := model.Book{
		Title:     "Book test",
		Author:    "Author 1",
		Publisher: "Publisher 1",
	}

	var err error
	if err = config.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBooksController(t *testing.T) {

	var testCases = []testCase{
		{
			name:       "(failed) wrong path name",
			path:       "/books",
			sizeData:   0,
			expectCode: http.StatusNotFound,
		},
		{
			name:       "(success) get list of books",
			path:       "/books",
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
				InsertMockBookData()
			}

			c.SetPath(testCase.path)
			err := controller.GetBooksController(c)
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()
			t.Log(body)

			var book BooksResponse
			err = json.Unmarshal([]byte(body), &book)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(book.Books))
		})
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []testCase{
		{
			name:       "(failed) get book with id",
			path:       "/books/:id",
			param:      "lorem",
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "(success) get book with id",
			path:       "/books/:id",
			param:      "1",
			expectCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec, c := initTestEnv(req)
			InsertMockBookData()

			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues(testCase.param)
			err := controller.GetBookController(c)

			// check handler error
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, testCase.expectCode, he.Code)
				}
			}

			body := rec.Body.String()

			var book BookResponse
			err = json.Unmarshal([]byte(body), &book)
			t.Log(book.Book)

			if err != nil {
				assert.Error(t, err, "error")
			} else {
				assert.Equal(t, testCase.param, strconv.Itoa(int(book.Book.ID)))
			}
		})
	}
}
