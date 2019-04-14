package main

// TIMEOUT定义了附件转化超时时间，单位为秒
// 配置较差的机器下，30M大小的docx文件转化为pdf大概需要10s
const TIMEOUT = 30

// WORKER代表同时启动n个worker接收转化任务，可以同时并发n个转化服务
// 建议设置为5
// 受限于unoconv的转化服务, 这里的worker数量并不是越多越好
const WORKER = 5

// MAXSIZE限制了上传文件的大小，单位为Kb，设置为-1代表不做限制
const MAXSIZE = 30 * 30
