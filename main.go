package main

import (
	"github.com/gin-gonic/gin"
)

var uno *UnoTaskQueue
var r *gin.Engine

func init() {
	initConfig()
	uno = initUnoconv()
}

func main() {
	r = gin.Default()
	r.Use(MaxSizeAllowed(MAXSIZE))

	r.GET("/", healthCheckHandler)
	r.POST("/unoconv/:toFileType", convertHandler)

	r.Run(":3000")
}
