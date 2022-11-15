package users

import (
	"github.com/MarcoVitangeli/bookstore_users-api/domain/users"
	"github.com/MarcoVitangeli/bookstore_users-api/services"
	"github.com/MarcoVitangeli/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
No business logic should be in the controller
*/

func getUserId(userIdParam string) (int64, *errors.RestErr) {
    userId, err := strconv.ParseInt(userIdParam, 10, 64)
    if err != nil {
        return 0, errors.NewBadRequestError("user id should be a number")
    }

    return userId, nil
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
    
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
        return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Create(c *gin.Context) {
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

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
        return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	res, err := services.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, res)
}

func Delete(c *gin.Context) {
    userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
        return
	}

    if err := services.DeleteUser(userId); err != nil {
        c.JSON(err.Status, err)
        return
    }

    c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
