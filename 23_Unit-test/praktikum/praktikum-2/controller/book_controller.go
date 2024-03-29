package controller

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	books := []model.Book{}

	err := config.DB.Find(&books).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(books) < 1 {
		return echo.NewHTTPError(http.StatusNotFound, "book list is empty!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id!")
	}
	var book model.Book

	err = config.DB.First(&book, id).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book",
		"book":    book,
	})
}

func CreateBooksController(c echo.Context) error {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Create(&book).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// find book with same id
	book := model.Book{}
	err := config.DB.First(&book, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// bind req to book
	updatedBook := model.Book{}
	err = c.Bind(&updatedBook)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// update the book
	err = config.DB.Model(&book).Updates(updatedBook).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := &model.Book{}
	err := config.DB.First(book, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Delete(book).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete",
		"book":    book,
	})
}
