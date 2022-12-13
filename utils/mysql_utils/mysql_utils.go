package mysql_utils

import (
	"errors"
	"strings"

	"github.com/PaulTabaco/bookstore_utils/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	if err == nil {
		return rest_errors.NewInternalServerError("", errors.New("error processing request"))
	}
	sqlErr, ok := err.(*mysql.MySQLError)
	// if is not mysql db error
	if !ok {
		// ID not exist case:
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("", errors.New("no error matching given id"))
		}
		// other non mysql errors:
		return rest_errors.NewInternalServerError("error parsing database response", err)
	}
	// if kind of mysql error
	switch sqlErr.Number {
	// Email already exists
	case 1062:
		return rest_errors.NewBadRequestError("", errors.New("invalid data"))
	}
	// other of mysql errors
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
