package route

import (
	"21_ORM-and-Code-Structure-MVC/praktikum/controller"

	"github.com/labstack/echo"
)

func NewInstance() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	// user route list
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)

	// books route list
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBooksController)
	e.PUT("/books/:id", controller.UpdateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)

	// posts route list
	e.POST("/posts", controller.CreatePostController)
	e.GET("/posts", controller.GetPostsController)
	e.GET("/posts/:id", controller.GetPostController)
	e.PUT("/posts/:id", controller.UpdatePostController)
	e.DELETE("/posts/:id", controller.DeletePostController)

	return e

}
