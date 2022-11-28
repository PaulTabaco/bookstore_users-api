package users

import (
	"strings"

	"github.com/PaulTabaco/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"` // `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"`  // `json:"last_name" binding:"required"`
	Email       string `json:"email"`      // `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

func (user *User) Validate() *errors.RestErr {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
