package app

import (
	"github.com/MarcoVitangeli/bookstore_users-api/controlers/ping"
	"github.com/MarcoVitangeli/bookstore_users-api/controlers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
    router.DELETE("/users/:user_id", users.Delete)
}
