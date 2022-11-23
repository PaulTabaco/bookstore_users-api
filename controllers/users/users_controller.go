package users

import (
	"net/http"
	"strconv"

	"github.com/PaulTabaco/bookstore_users-api/domain/users"
	"github.com/PaulTabaco/bookstore_users-api/services"
	"github.com/PaulTabaco/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, &restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userIdErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userIdErr != nil {
		err := errors.NewBadRequestError("invalide user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementing of SearchUser coming soon")
}
