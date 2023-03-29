package route

import (
	"21_ORM-and-Code-Structure-MVC/praktikum/controller"

	"github.com/labstack/echo"
)

func NewInstance() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	return e
}
