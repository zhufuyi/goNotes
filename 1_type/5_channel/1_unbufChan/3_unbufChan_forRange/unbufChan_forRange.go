// 非缓冲通道的for_range 用法

package main

import (
	"fmt"
	"time"
)

var exit = make(chan bool) // 退出通知

// 解析命令
func parseCmd(uch <-chan string) { // 隐式转换为只读通道
	for cmd := range uch { // 一直循环等待读取数据，直到关闭通道退出循环
		fmt.Println("cmd =", cmd)
	}
	exit <- true // 发出退出通知
}

// 发送命令
func sendCmd(uch chan<- string, cmd string) {
	uch <- cmd                         // 向通道发送数据
	time.Sleep(500 * time.Millisecond) // 延时0.5s
}

func main() {
	uch := make(chan string) // 初始化

	go parseCmd(uch) // 接收

	sendCmd(uch, "read")
	sendCmd(uch, "write")
	sendCmd(uch, "update")
	sendCmd(uch, "delete")
	sendCmd(uch, "insert")
	sendCmd(uch, "bye")

	close(uch) // 关闭通道
	<-exit     // 等待退出

	// 结果：
	// cmd = read
	// cmd = write
	// cmd = update
	// cmd = delete
	// cmd = insert
	// cmd = bye
}
