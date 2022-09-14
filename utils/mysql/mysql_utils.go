package mysql

import (
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errNoRows       = "no rows in result set"
	DuplicateKeyErr = 1062
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		/**
		we don't validate this error as mysql because it is used by
		the packaged sql as `var ErrNoRows = error.New(// err string)`
		and it is hardcoded
		*/
		if strings.Contains(err.Error(), errNoRows) {
			return errors.NewBadRequestError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case DuplicateKeyErr:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
