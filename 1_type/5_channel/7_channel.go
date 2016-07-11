package main

import (
	"fmt"
	"time"
)

var sum = 0

func count(mt chan int, n int) {
	if n != 0 {
		time.Sleep(100 * time.Duration(n) * time.Millisecond)
	}

	sum++
	mt <- sum
}

func synchronization() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		go count(ch, i)
	}

	flag := 0
	for {
		select {
		case s := <-ch: // 通道只有一个，写满后，必须读取才会执行下一次写入
			fmt.Println("read ch, sum =", s)
			flag++
		case <-time.After(1 * time.Second):
			flag = 10
			fmt.Println("time out")
		}
		if flag >= 10 {
			break
		}
	}
	fmt.Println("    synchronization sum =", sum)
}

func asynchronous() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		go count(ch, 0)
	}

	flag := 0
	for {
		select {
		case s := <-ch: // 通道只有一个，写满后，必须读取才会执行下一次写入
			fmt.Println("read ch, sum =", s)
			flag++
		}
		if flag >= 10 {
			break
		}
	}
	fmt.Println("    asynchronous sum =", sum)
}

func main() {
	synchronization() // 不带缓冲区的channel
	asynchronous()    // 带缓冲区的channel
}
