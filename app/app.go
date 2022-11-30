package app

import (
	"github.com/PaulTabaco/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApp() {
	MapUrls()
	logger.Log.Info("about to start the application")
	router.Run(":8080")
}
