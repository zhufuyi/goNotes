// 缓冲chan的select和for range的使用

package main

import (
	"fmt"
	"sync"
	"time"
)

var exit = make(chan bool) // 退出通知

// select接收
func receiveSelect(bch <-chan int) { // 隐式转换为只读通道
	defer func() { // 发送退出通知
		exit <- true
	}()

	fmt.Printf("rcvData =")
	for {
		select { // 可以用select作为接发送或接收选择器
		case vl, ok := <-bch: // 当关闭通道后，如果有数据也会读出来
			if !ok { // 判断通道是否已经关闭
				fmt.Println("\n通道已经关闭。\n")
				return
			}
			fmt.Printf(" %d,", vl)
		}
	}
}

// 发送，对应select接收
func send1(bch chan<- int, cnt int) {
	wg := new(sync.WaitGroup)
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(i int) { // 并发发送数据
			// 设置发送超时，避免永久阻塞
			select {
			case bch <- i:
			case <-time.After(time.Second):
				fmt.Println("i =", i, "发送超时")
			}
			wg.Done()
		}(i)
	}
	// 当使用并发发送时，不能马上关闭通道，需要等待所有并发goroutine创建完成才能关闭通道，否则会引起恐慌
	wg.Wait()
	close(bch) //发送完毕，关闭通道
}

func selectChan(bch chan int, info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	go receiveSelect(bch)
	go send1(bch, 30)
	<-exit

	// 结果：
	// rcvData = 1, 0, 22, 16, 17, 18, 19, 20, 21, 25, 23, 24, 8, 2, 3, 4, 5, 6, 7, 11, 9, 10, 13, 12, 14, 27, 26, 28, 15, 29,
	// 通道已经关闭。
}

//----------------------------------------------------------------------------------------------

//  forRange 接收
func receiveForRange(bch <-chan int) { // 隐式转换为只读通道
	fmt.Printf("rcvData =")
	for vl := range bch { // 一直循环等待读取数据，直到关闭通道退出循环
		fmt.Printf(" %d,", vl)
	}
	fmt.Println("\n通道已经关闭。\n")
	exit <- true // 发出退出通知
}

// 发送，对应forRange接收
func send2(bch chan int, cnt int) {
	for i := 0; i < cnt; i++ {
		// 设置发送超时
		select {
		case bch <- i:
		case <-time.After(time.Second):
			fmt.Println("i =", i, "发送超时")
		}
	}
	close(bch) //发送完毕，关闭通道
}

func forRangeChan(bch chan int, info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	go receiveForRange(bch)
	go send2(bch, 50)
	<-exit
	// 结果：
	// rcvData = 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	// 通道已经关闭。
}

func main() {
	bch := make(chan int, 5)
	selectChan(bch, "select方式接收数据")

	bch = make(chan int, 10)
	forRangeChan(bch, "for range方式接收数据")
}
