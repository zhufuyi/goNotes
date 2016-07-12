// 非缓冲chan的select和for range的使用

package main

import (
	"fmt"
	"time"
)

var exit = make(chan bool) // 退出通知

// 解析命令
func parseCmdSelect(uch <-chan string) { // 隐式转换为只读通道
	defer func() {
		exit <- true
	}()

	for {
		select { // 可以用select作为接发送或接收选择器
		case cmd, ok := <-uch: // 当关闭通道后，如果有数据也会读出来
			if !ok { // 判断通道是否已经关闭
				return
			}
			fmt.Printf("cmd = %s\n", cmd)
		}
	}

}

// 发送命令，对应parseCmdSelect函数
func sendCmd1(uch chan<- string, cmd string, cnt int) {
	time.Sleep(time.Duration(cnt) * time.Millisecond) // 发送时间间隔
	select {
	case uch <- cmd: // 向通道发送数据
	case <-time.After(time.Second): // 设置发送超时1s
	}
}

func selectChan(uch chan string, info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	go parseCmdSelect(uch) // 接收

	sendCmd1(uch, "read", 10)
	sendCmd1(uch, "write", 200)
	sendCmd1(uch, "update", 600)
	sendCmd1(uch, "delete", 800)
	sendCmd1(uch, "insert", 250)
	sendCmd1(uch, "bye", 10)
	close(uch) // 关闭通道
	<-exit

	// 结果：
	// cmd = read
	// cmd = write
	// cmd = update
	// cmd = delete
	// cmd = insert
	// cmd = bye
}

//----------------------------------------------------------------------------------------------

// 解析命令
func parseCmdForRange(uch <-chan string) { // 隐式转换为只读通道
	for cmd := range uch { // 一直循环等待读取数据，直到关闭通道退出循环
		fmt.Println("cmd =", cmd)
	}
	exit <- true // 发出退出通知
}

// 发送命令
func sendCmd2(uch chan<- string, cmd string) {
	select {
	case uch <- cmd: // 向通道发送数据
	case <-time.After(time.Second): // 设置发送超时1s
	}
}

func forRangeChan(uch chan string, info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	go parseCmdForRange(uch)

	sendCmd2(uch, "cp")
	sendCmd2(uch, "touch")
	sendCmd2(uch, "cd")
	sendCmd2(uch, "rm")
	sendCmd2(uch, "mv")
	sendCmd2(uch, "bye")

	close(uch) // 关闭通道
	<-exit     // 等待退出

	// 结果：
	// cmd = cp
	// cmd = touch
	// cmd = cd
	// cmd = rm
	// cmd = mv
	// cmd = bye
}

func main() {
	uch := make(chan string) // 初始化
	selectChan(uch, "select方式接收数据")

	uch = make(chan string) // 初始化
	forRangeChan(uch, "for range方式接收数据")
}
