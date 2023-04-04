package controller

import (
	"22_Middleware/praktikum/config"
	m "22_Middleware/praktikum/middleware"
	"22_Middleware/praktikum/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// // get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := &model.User{}

	// search user with id from req param
	err := config.DB.First(user, id).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user with id " + fmt.Sprint(id),
		"users":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Create(&user).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := &model.User{}

	err := config.DB.First(user, id).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = config.DB.Delete(user, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := &model.User{}
	userUpdate := &model.User{}

	err := config.DB.First(user, id).Where("id = ?", id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(userUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Update attributes with `struct`, will only update non-zero fields
	config.DB.Model(user).Updates(userUpdate)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update",
		"user":    user,
	})

}

// login
func LoginUserFunc(c echo.Context) error {
	loginReq := new(model.User)
	c.Bind(loginReq)

	email := loginReq.Email
	password := loginReq.Password

	userData := new(model.User)

	err := config.DB.First(&userData, "email = ? AND password = ?", email, password).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user not found")
	}

	t, err := m.GenerateToken(email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"token":   t,
	})

}
