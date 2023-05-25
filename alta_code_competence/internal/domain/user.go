package domain

type User struct {
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"password"`
}

type UserAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
