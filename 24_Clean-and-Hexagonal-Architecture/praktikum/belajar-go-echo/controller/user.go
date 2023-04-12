package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface{}

type userController struct {
	usercase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	data, err := u.usercase.GetAll()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": data,
	})

}

func (u *userController) CreateUser(c echo.Context) error {
	userPayload := dto.CreateUserRequest{}
	c.Bind(&userPayload)
	data, err := u.usercase.Create(userPayload)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": data,
	})
}

func (u *userController) Login(c echo.Context) error {
	userPayload := dto.LoginUserRequest{}
	c.Bind(&userPayload)
	data, err := u.usercase.Login(userPayload)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": data,
	})
}

func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := make([]model.User, 0)
		err := db.Find(&users).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func CreateUser(db *gorm.DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		err := db.Create(&user).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}
