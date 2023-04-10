package route

import (
	"23_Unit-test/praktikum/praktikum-2/constants"
	"23_Unit-test/praktikum/praktikum-2/controller"
	"23_Unit-test/praktikum/praktikum-2/util"
	m "23_Unit-test/praktikum/praktikum-2/middleware"

	eMid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func NewInstance() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	m.Logger(e)
	util.InitValidateConf(e)

	eJWT := e.Group("")
	eJWT.Use(eMid.JWT([]byte(constants.JWT_SECRET)))
	// Route / to handler function
	// user route list
	eJWT.GET("/users", controller.GetUsersController)
	eJWT.GET("/users/:id", controller.GetUserController)
	eJWT.PUT("/users/:id", controller.UpdateUserController)
	eJWT.DELETE("/users/:id", controller.DeleteUserController)
	e.POST("/users", controller.CreateUserController)
	e.POST("/users/login", controller.LoginUserFunc)

	// books route list
	eJWT.GET("/books", controller.GetBooksController)
	eJWT.GET("/books/:id", controller.GetBookController)
	eJWT.POST("/books", controller.CreateBooksController)
	eJWT.PUT("/books/:id", controller.UpdateBookController)
	eJWT.DELETE("/books/:id", controller.DeleteBookController)

	// posts route list
	e.POST("/posts", controller.CreatePostController)
	e.GET("/posts", controller.GetPostsController)
	e.GET("/posts/:id", controller.GetPostController)
	e.PUT("/posts/:id", controller.UpdatePostController)
	e.DELETE("/posts/:id", controller.DeletePostController)

	return e

}
