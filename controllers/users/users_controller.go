package users

import (
	"net/http"

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
		//restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementing of GetUser coming soon")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementing of SearchUser coming soon")
}
