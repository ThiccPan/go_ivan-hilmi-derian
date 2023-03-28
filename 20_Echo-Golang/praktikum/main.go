package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// find user using linear search
func findUser(users []User, id int) int {
	for idx, user := range users {
		if user.Id == id {
			return idx
		}
	}
	return -1
}

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	// search user with id from req param
	userIdx := findUser(users, id)
	if userIdx == -1 {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	user := users[userIdx]
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user with id " + fmt.Sprint(id),
		"users":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	// search user with id from req param
	userIdx := findUser(users, id)
	if userIdx == -1 {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	tmp := users[userIdx+1:]
	users = users[:userIdx]
	users = append(users, tmp...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success deleting user with id " + fmt.Sprint(id),
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	// search user with id from req param
	userIdx := findUser(users, id)
	if userIdx == -1 {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	user := &users[userIdx]
	newUser := User{}
	c.Bind(&newUser)

	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"user":     user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.POST("/users", CreateUserController)
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
