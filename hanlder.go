package main

import (
	"github.com/gin-gonic/gin"
	"log"
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

	//tempfile, err := ioutil.TempFile(os.TempDir(), "gouno")
	//ifExistErrorAbort400(c, err)
	//defer tempfile.Close()
	//
	//_, err = io.Copy(tempfile, file)
	//ifExistErrorAbort500(c, err)
	//
	//filename := tempfile.Name() + filepath.Ext(header.Filename)
	//os.Rename(tempfile.Name(), filename)
	//defer os.Remove(filename)

	toFileType := c.Param("toFileType")
	// 发送转化任务
	err = uno.Send(toFileType, file, c.Writer, isCompress)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
			"code": http.StatusRequestEntityTooLarge,
			"msg":  err.Error(),
		})
		return
	}
}

func ifExistErrorAbort400(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
}

func ifExistErrorAbort500(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
}
