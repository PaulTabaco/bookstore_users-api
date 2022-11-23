package users

import (
	"strings"

	"github.com/PaulTabaco/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"  binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	// Just for tutorial flow, but for better way I used - binding: in json tags
	if strings.TrimSpace(strings.ToLower(user.Email)) == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
