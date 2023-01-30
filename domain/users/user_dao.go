package users

import (
	"fmt"
	"github.com/MarcoVitangeli/bookstore_users-api/datasources/mysql/users_db"
	"github.com/MarcoVitangeli/bookstore_users-api/logger"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	errorNoRows           = "no rows in result set"
	MySqlDuplicateKeyErr  = 1062
)

// only entry point for our database

// Save saves the user if the ID is not present
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error trying to prepare save user query", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error trying to save user", err)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error trying to get last inserted id", err)
		return errors.NewInternalServerError("database error")
	}

	user.Id = userId
	return nil
}

// Get gets the user by ID (primary key)
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error trying to prepare get user by id query", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow(user.Id)

	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		logger.Error("error trying to execute update user statement", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("error trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	if err != nil {
		logger.Error("error trying to execute delete user statement", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)

	if err != nil {
		logger.Error("error preparing find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error querying data for finding user by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	usersArr := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.Id,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.DateCreated,
			&u.Status,
		); err != nil {
			logger.Error("error reading data from find by status query", err)
			return nil, errors.NewInternalServerError("database error")
		}
		usersArr = append(usersArr, u)
	}

	if err := rows.Err(); err != nil {
		logger.Error("rows error", err)
		return nil, errors.NewInternalServerError("database error")
	}

	if len(usersArr) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return usersArr, nil
}
