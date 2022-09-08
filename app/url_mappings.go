package app

import (
	"github.com/MarcoVitangeli/bookstore_users-api/controlers/ping"
	"github.com/MarcoVitangeli/bookstore_users-api/controlers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
