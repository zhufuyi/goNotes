// 缓冲chan的声明和使用

/*
	缓冲chan和非缓冲chan不同点就是异步和同步传输数据。使用缓冲chan更加高效。
*/

package main

import (
	"fmt"
	"time"
)

// 接收，注意不要在接收端关闭通道，如果要在接收端关闭通道，需要做些辅助手段通知发送端，否则发送端向关闭的通道发送数据会引起恐慌
func receive(bch chan int) {
	fmt.Printf("vl =")
	for {
		vl, ok := <-bch // 当关闭通道后，如果有数据也会读出来
		if !ok {
			fmt.Println("\n通道已经关闭")
			break
		}
		fmt.Printf(" %d,", vl)
	}
}

// 发送
func send(bch chan int, cnt int) {
	for i := 0; i < cnt; i++ {
		func(i int) {
			bch <- i
		}(i)
	}
	time.Sleep(2 * time.Millisecond)
	close(bch) //发送完毕，关闭通道
}

func main() {
	bufChan := make(chan int, 10)

	go send(bufChan, 20)
	go receive(bufChan)

	time.Sleep(2 * time.Second)

	//	close(bufChan) // 注意：关闭 已经关闭过的通道会引起恐慌

	// 结果：
	// vl = 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	// 通道已经关闭
}
