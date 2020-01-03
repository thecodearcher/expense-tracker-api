package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"math"
)

//User model
type User struct {
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT" `
	Name      string     `json:"name" `
	Email     string     `json:"email" `
	Password  string     `json:"-" `
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

//LoginUser struct for logging in
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Validate user model
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, math.MaxInt64)),
	)
}

//Validate login data
func (u LoginUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required),
	)
}
