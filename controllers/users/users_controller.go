package users

import (
	"fmt"
	"net/http"

	"github.com/PaulTabaco/bookstore_users-api/domain/users"
	"github.com/PaulTabaco/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println("Read error", err)
	// 	//TODO: - Handle error
	// 	return
	// }

	// if err = json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println("Unmarshal err", err, "Bytes - ", string(bytes))
	// 	//TODO: - Handle error
	// 	return
	// }

	/// !! next ..ShouldBindJSON.. do same as commented code above
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("ShouldBindJSON", err)
		//TODO: - Handle json error
		return
	}

	result, err := services.CreateUser(user)

	if err != nil {
		//TODO: - Handle error
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
