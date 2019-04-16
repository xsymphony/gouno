package main

import (
	"github.com/gin-gonic/gin"
)

var uno *UnoTaskQueue

func init() {
	initConfig()
	uno = initUnoconv()
}

func main() {
	r := SetupRouter()
	r.Run(":3000")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(MaxSizeAllowed(MAXSIZE))

	r.GET("/", healthCheckHandler)
	r.POST("/unoconv/:toFileType", convertHandler)
	r.POST("/cached/unoconv/:toFileType", convertWithCacheHandler)

	return r
}
