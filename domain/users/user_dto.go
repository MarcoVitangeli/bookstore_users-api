package users

import (
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
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
	Status      string `json:"status"`
	Password    string `json:"password"` // this means that we only use this as an internal field
}

type Users []User

// user should know if it is valid

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("invalid empty password")
	}

	return nil
}
