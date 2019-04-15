package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func convertHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	ifExistErrorAbort400(c, err)
	defer file.Close()

	// isCompress 是否降低转化后的文件质量
	isCompress, err := strconv.ParseBool(c.Request.PostFormValue("compress"))
	if err != nil {
		isCompress = false
	}

	toFileType := c.Param("toFileType")
	// 发送转化任务
	err = uno.Send(toFileType, file, c.Writer, isCompress)
	ifExistErrorAbort500(c, err)
}

func genericAbortCode(c *gin.Context, err error, code int) {
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
		return
	}
}

func ifExistErrorAbort400(c *gin.Context, err error) {
	genericAbortCode(c, err, http.StatusBadRequest)
}

func ifExistErrorAbort500(c *gin.Context, err error) {
	genericAbortCode(c, err, http.StatusInternalServerError)
}
