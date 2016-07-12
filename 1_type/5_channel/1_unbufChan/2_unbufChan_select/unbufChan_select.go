// 判断通道发送数据是否超时

package main

import (
	"fmt"
	"time"
)

// 解析命令
func parseCmd(uch <-chan string) { // 隐式转换为只读通道
	t := time.NewTimer(500 * time.Millisecond)
	toFlag := true
	for {
		select { // 可以用select作为接发送或接收选择器
		case cmd, ok := <-uch: // 当关闭通道后，如果有数据也会读出来
			if !ok { // 判断通道是否已经关闭
				t.Stop()
				return
			}
			t.Reset(500 * time.Millisecond)
			if !toFlag { // 忽略超时的数据
				fmt.Printf("cmd = %s\n", cmd)
			}
			toFlag = false

		case <-t.C:
			toFlag = true
			fmt.Println("接收超时")
		}
	}
}

// 发送命令
func sendCmd(uch chan<- string, cmd string, cnt int) {
	time.Sleep(time.Duration(cnt) * time.Millisecond)
	uch <- cmd // 向通道发送数据
}

func main() {
	uch := make(chan string) // 初始化

	go parseCmd(uch) // 接收

	// 超过500时则超时
	sendCmd(uch, "read", 600)
	sendCmd(uch, "write", 400)
	sendCmd(uch, "update", 800)
	sendCmd(uch, "delete", 300)
	sendCmd(uch, "insert", 550)
	sendCmd(uch, "bye", 450)

	close(uch) // 关闭通道
	time.Sleep(time.Millisecond)

	// 结果：
	// 接收超时
	// cmd = write
	// 接收超时
	// cmd = delete
	// 接收超时
	// cmd = bye
}
