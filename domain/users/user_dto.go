package users

import (
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"strings"
)

/**
Domain is the layer that represents an entity
from the domain that this API belongs

this is the core of our microservice
*/
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// user should know if it is valid

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
