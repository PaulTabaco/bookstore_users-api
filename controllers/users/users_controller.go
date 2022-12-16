package users

import (
	"net/http"
	"strconv"

	"github.com/PaulTabaco/bookstore_oauth/oauth"
	"github.com/PaulTabaco/bookstore_users-api/domain/users"
	"github.com/PaulTabaco/bookstore_users-api/services"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userIdErr := strconv.ParseInt(userIdParam, 10, 64)
	if userIdErr != nil {
		return 0, rest_errors.NewBadRequestError("invalide user id")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	// Prevent access resources with not existed in db token;
	// It's not optimal way for me to look for callerId ==0 in this case, but used in tutorial  callerId 0 returns if token not exists);
	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := rest_errors.NewUnauthorizedError("resource is not available")
		c.JSON(err.Status(), err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	// if callerId == result user.id - Marshal to return internal user
	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	// Othervice marshal to return int or public user depending of X-Public header
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
	// c.JSON(http.StatusOK, user)
}
