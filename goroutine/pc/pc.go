package main

import (
	"fmt"
	"time"
)

// 协程1 产生的数据写入到通道
func production(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

// 协程2 从通道读取数据来消费
func consumption(ch chan int) {
	for val := range ch {
		fmt.Printf("%d ", val)
		time.Sleep(time.Millisecond * 200)
	}

	//for {
	//	select {
	//	case val, ok := <-ch:
	//		if !ok {
	//			return
	//		}
	//		fmt.Printf("%d ", val)
	//		time.Sleep(time.Millisecond * 200)
	//	}
	//}
}

func main() {
	numCh := make(chan int)

	go production(numCh)
	go consumption(numCh)

	time.Sleep(5 * time.Second)
}
