package test

import (
	"22_Middleware/praktikum/model"
	"23_Unit-test/praktikum/praktikum-2/config"
)

type BooksResponse struct {
	Message string
	Users   []model.Book
}

type BookResponse struct {
	Message string
	User    model.Book
}

func InsertMockBookData() error {
	user := model.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}