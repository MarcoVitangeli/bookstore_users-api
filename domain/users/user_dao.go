package users

import (
	"github.com/MarcoVitangeli/bookstore_users-api/datasources/mysql/users_db"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/date"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/mysql"
)

const (
	queryInsertUser      = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser         = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser      = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	errorNoRows          = "no rows in result set"
	MySqlDuplicateKeyErr = 1062
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysql.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql.ParseError(saveErr)
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

	res := stmt.QueryRow(user.Id)

	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		mysql.ParseError(err)
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql.ParseError(err)
	}

	return nil
}
