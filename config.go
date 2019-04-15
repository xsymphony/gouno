package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var TIMEOUT time.Duration
var WORKER int
var MAXSIZE int64

func initConfig() {
	var err error
	// TIMEOUT定义了附件转化超时时间，单位为秒
	// 配置较差的机器下，30M大小的docx文件转化为pdf大概需要10s
	TIMEOUT, err = Atodu(os.Getenv("TIMEOUT"))
	if err != nil {
		TIMEOUT = 30
		log.Printf("[Init config] Read TIMEOUT failed, use default value with %d", TIMEOUT)
	}
	// WORKER代表同时启动n个worker接收转化任务，可以同时并发n个转化服务
	// 建议设置为5
	// 受限于unoconv的转化服务, 这里的worker数量并不是越多越好
	WORKER, err = strconv.Atoi(os.Getenv("WORKER"))
	if err != nil {
		WORKER = 5
		log.Printf("[Init config] Read WORKER failed, use default value with %d", WORKER)
	}
	// MAXSIZE限制了上传文件的大小，单位为B，设置为-1代表不做限制，当前设置为50M
	MAXSIZE, err = strconv.ParseInt(os.Getenv("MAXSIZE"), 10, 0)
	if err != nil {
		MAXSIZE = 50 * 1024 * 1024
		log.Printf("[Init config] Read MAXSIZE failed, use default value with %d", MAXSIZE)
	}
}

func Atodu(s string) (time.Duration, error) {
	t, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return time.Duration(t), nil
}
