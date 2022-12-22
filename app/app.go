package app

import (
	"github.com/PaulTabaco/bookstore_utils/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApp() {
	MapUrls()
	// logger.Log.Info("about to start the application") // Info in logger.Log
	logger.Info("about to start the application") // Customized by me in logger with Sync()
	router.Run(":8081")
}
