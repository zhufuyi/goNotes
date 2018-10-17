// 无缓冲通道的初始化和使用
/*
无缓冲通道，即用make初始化时设置第二个参数为0或者省略第二个参数

发送和接收直接同步传递，发送和接收那个动作先进行不重要，重要的是发送和接收要有配对的机会。

实现多个goroutine之间同步

*/
package main

import (
	"fmt"
	"time"
)

// 解析命令
func parseCmd(uch chan string) {
	for {
		cmd, ok := <-uch // 没有数据时，一直阻塞在这里，当检测到通道关闭时，会跳出阻塞
		if !ok {         // 判断通道是否已经关闭，注意：向关闭通道发送会引起恐慌
			break
		}
		fmt.Println("cmd =", cmd)
	}
}

// 发送命令
func sendCmd(uch chan string, cmd string) {
	uch <- cmd                         // 向通道发送数据
	time.Sleep(500 * time.Millisecond) // 等待0.5s
}

func main() {
	uch := make(chan string) // 初始化

	go parseCmd(uch) // 接收

	sendCmd(uch, "read")
	sendCmd(uch, "write")
	sendCmd(uch, "update")
	sendCmd(uch, "delete")
	sendCmd(uch, "bye")

	close(uch) // 关闭通道
	time.Sleep(time.Millisecond)

	// 结果：
	// cmd = read
	// cmd = write
	// cmd = update
	// cmd = delete
	// cmd = bye
}
