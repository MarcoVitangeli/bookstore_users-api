package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

/**
All go sql drivers (mysql, postgres, etc)
implement the database/sql interface
*/

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB
)

// init is a function called once when first imported this package
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var (
		username = os.Getenv(mysqlUsersUsername)
		password = os.Getenv(mysqlUsersPassword)
		host     = os.Getenv(mysqlUsersHost)
		schema   = os.Getenv(mysqlUsersSchema)
	)

	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
