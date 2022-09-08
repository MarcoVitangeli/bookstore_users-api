package users

import (
	"github.com/MarcoVitangeli/bookstore_users-api/domain/users"
	"github.com/MarcoVitangeli/bookstore_users-api/services"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
No business logic should be in the controller
*/

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
