package users

import (
	"fmt"
	"github.com/MarcoVitangeli/bookstore_users-api/datasources/mysql/users_db"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/date"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	errorNoRows     = "no rows in result set"
)

var (
	usersDB = make(map[int64]*User)
)

// only entry point for our database

// Save saves the user if the ID is not present
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error trying to save user: %s", err.Error()))
	}

	user.Id = userId
	return nil
}

// Get gets the user by ID (primary key)
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow()

	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get user %d", user.Id))
	}
	return nil
}
