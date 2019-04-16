package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiongsyao/gouno/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func convertHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err})
		return
	}
	defer file.Close()

	// isCompress 是否降低转化后的文件质量
	isCompress, err := strconv.ParseBool(c.Request.PostFormValue("compress"))
	if err != nil {
		isCompress = false
	}
	// 将要转化为的文件类型
	toFileType := c.Param("toFileType")
	// 发送转化任务
	err = uno.Send(toFileType, file, c.Writer, isCompress)
	if err != nil {
		InternalServerErrorResponse(c, err, "执行转化任务时失败!")
		return
	}
}

func convertWithCacheHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err})
		return
	}
	defer file.Close()

	// 读取上传的文件到临时文件里
	tempFile, err := ioutil.TempFile("", header.Filename)
	if err != nil {
		InternalServerErrorResponse(c, err, "创建临时文件时失败!")
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
	err = utils.CopyAndSeek0(tempFile, file)
	if err != nil {
		InternalServerErrorResponse(c, err, "写入到临时文件时失败!")
		return
	}

	// 计算文件的hash值
	hash, err := utils.CalculateHash(tempFile)
	if err != nil {
		InternalServerErrorResponse(c, err, "计算文件hash值时失败!")
		return
	}
	toFileType := c.Param("toFileType")

	// 获取缓存文件路径
	toFilePath, err := utils.JoinCacheFile(CACHE_FILE_DIR, hash+"."+toFileType)
	if err != nil {
		InternalServerErrorResponse(c, err, "获取缓存文件路径时失败!")
		return
	}

	if !utils.IsFileExists(toFilePath) {
		f, err := os.Create(toFilePath)
		if err != nil {
			InternalServerErrorResponse(c, err, "创建转化后的文件时失败!")
			return
		}
		defer f.Close()

		// 转化结果写入到缓存文件中
		tempFile.Seek(0, 0)
		err = uno.Send(toFileType, tempFile, f, false)
		if err != nil {
			InternalServerErrorResponse(c, err, "执行转化任务时失败!")
			return
		}
	}
	c.File(toFilePath)
}

func InternalServerErrorResponse(c *gin.Context, err error, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":  http.StatusInternalServerError,
		"error": err.Error(),
		"msg":   msg,
	})
}
