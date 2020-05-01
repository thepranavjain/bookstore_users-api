package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/thepranavjain/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no resource with matching id found")
		}
		return errors.NewInternalServerError(fmt.Sprintf(
			"error processing the request: %s", err.Error(),
		))
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("resource already exists")
	}
	return errors.NewInternalServerError(fmt.Sprintf(
		"error processing the request: %s", err.Error(),
	))
}
