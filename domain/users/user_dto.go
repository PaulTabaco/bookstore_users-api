package users

import (
	"errors"
	"strings"

	"github.com/PaulTabaco/bookstore_utils/rest_errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"` // `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"`  // `json:"last_name" binding:"required"`
	Email       string `json:"email"`      // `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` //`json:"-"`
}

type Users []User

func (user *User) Validate() *rest_errors.RestErr {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("", errors.New("invalid email address"))
	}

	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return rest_errors.NewBadRequestError("", errors.New("invalid password"))
	}
	return nil
}
