package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func MaxSizeAllowed(n int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if n == -1 {
			c.Next()
			return
		}
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, n)

		buff, errRead := c.GetRawData()

		if errRead != nil {
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
				"code": http.StatusRequestEntityTooLarge,
				"msg":  fmt.Sprintf("file is too large, more than %d B", n),
			})
			return
		}

		buf := bytes.NewBuffer(buff)
		c.Request.Body = ioutil.NopCloser(buf)
		c.Next()
	}
}
