package controller

import (
	"23_Unit-test/praktikum/praktikum-2/config"
	"23_Unit-test/praktikum/praktikum-2/model"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/labstack/echo"
)

func CreatePostController(c echo.Context) error {
	post := model.Post{}
	user := model.User{}

	c.Bind(&post)

	err := config.DB.First(&user, post.User_id).Where("id = ?", post.User_id).Error

	// check if user with id = user_id exist
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()+", user_id not found in table!")
	}

	// save data to db
	err = config.DB.Create(&post).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add post",
		"post":    post,
	})
}

func GetPostsController(c echo.Context) error {
	posts := []model.Post{}

	err := config.DB.Find(&posts).Error
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all posts",
		"posts":   posts,
	})
}

func GetPostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	post := model.Post{}

	err := config.DB.First(&post, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get post",
		"post":    post,
	})

}

func UpdatePostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	post := &model.Post{}
	postUpdate := &model.Post{}

	err := config.DB.First(post, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(postUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = config.DB.Model(post).Updates(postUpdate).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get post",
		"post":    post,
	})
}

func DeletePostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	post := &model.Post{}
	err := config.DB.First(post, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Delete(post).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete",
		"post":    post,
	})
}
