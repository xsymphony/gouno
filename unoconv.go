package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// Task代表一个转化任务，其中包括了
// fromFilePath:   文件完整路径名
// toFileType:     要转化为的文件类型，包括doc xls html pdf等
// isCompress:     是否压缩文件（例如降低图像质量)
// w:              输出重定向的位置（转化后的文件）
// errMsgCh:       存储转化错误信息
type Task struct {
	w          io.Writer
	r          io.Reader
	toFileType string
	isCompress bool
	errMsgCh   chan error
}

func (t *Task) Run() {
	var cmd *exec.Cmd
	if t.isCompress {
		// 压缩文件内容
		cmd = exec.Command("/bin/unoconv",
			"-f", t.toFileType,
			"-e", "UseLosslessCompression=false",
			"-e", "ReduceImageResolution=false",
			"--stdout",
			"--stdin",
		)
	} else {
		// 不压缩文件
		cmd = exec.Command("/bin/unoconv",
			"-f", t.toFileType,
			"--stdout",
			"--stdin",
		)
	}
	// 输出重定向到转化任务的writer
	cmd.Stdin = t.r
	cmd.Stdout = t.w
	stderr := bytes.NewBufferString("Unoconv execute failed:\r\n")
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		t.errMsgCh <- err
	}
	done := make(chan error)

	go func() { done <- cmd.Wait() }()

	// 如果转化超时则抛出异常
	select {
	case err := <-done:
		if err != nil {
			fmt.Println(stderr.String())
		}
		t.errMsgCh <- err
	case <-time.After(TIMEOUT * time.Second):
		t.errMsgCh <- errors.New(
			fmt.Sprintf("Timeout: task execute more than %d seconds.", TIMEOUT),
		)
	}
}

// UnoTaskQueue是转化任务队列
type UnoTaskQueue struct {
	ch chan *Task
}

// Send向UnoTaskQueue任务队列里添加新的任务并等待任务执行结果
func (u *UnoTaskQueue) Send(toFileType string, r io.Reader, w io.Writer, isCompress bool) error {
	// 定义error chan，实现跨协程记录错误信息
	msg := make(chan error)
	t := &Task{
		toFileType: toFileType,
		w:          w,
		r:          r,
		errMsgCh:   msg,
		isCompress: isCompress,
	}
	u.ch <- t
	return <-msg
}

// Consume运行一个转化任务
func (u *UnoTaskQueue) Consume() {
	for {
		select {
		case t := <-u.ch:
			t.Run()
		}
	}
}

// initUnconv初始化UnoTaskQueue，并开始n个worker持续执行转化任务
func initUnoconv() *UnoTaskQueue {
	uno := &UnoTaskQueue{ch: make(chan *Task, WORKER)}

	startRunWorkers(WORKER, uno)

	return uno
}

// startRunWorkers启动n个转化任务的worker
func startRunWorkers(n int, u *UnoTaskQueue) {
	for i := 0; i < n; i++ {
		go u.Consume()
	}
}
