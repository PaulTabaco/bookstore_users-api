package mysql_utils

import (
	"strings"

	"github.com/PaulTabaco/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
	// queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	// queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=? ;"
)

func ParseError(err error) *errors.RestErr {
	if err == nil {
		return errors.NewInternalServerError("error processing request")
	}
	sqlErr, ok := err.(*mysql.MySQLError)
	// if is not mysql db error
	if !ok {
		// ID not exist case:
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no error matching given id")
		}
		// other non mysql errors:
		return errors.NewInternalServerError("error parsing database response")
	}
	// if kind of mysql error
	switch sqlErr.Number {
	// Email already exists
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	// other of mysql errors
	return errors.NewInternalServerError("error processing request")
}
