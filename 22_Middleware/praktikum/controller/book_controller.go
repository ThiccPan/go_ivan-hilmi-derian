package controller

import (
	"22_Middleware/praktikum/config"
	"22_Middleware/praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	err := config.DB.Find(&books).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book

	err := config.DB.First(&book, id).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success get book",
		"book":   book,
	})
}

func CreateBooksController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	err := config.DB.Create(&book).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add book",
		"book":   book,
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
