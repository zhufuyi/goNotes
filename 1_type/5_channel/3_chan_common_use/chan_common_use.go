// chan的常用场景和注意事项

/*
总结：
发送：
• 向 nil channel 发送数据，阻塞。
• 向 closed channel 发送数据，出错。
• 同步发送: 如有接收者，交换数据。否则排队、阻塞。
• 异步发送: 如有空槽，拷贝数据到缓冲区。否则排队、阻塞。
接收：
• 从 nil channel 接收数据，阻塞。
• 从 closed channel 接收数据，返回已有数据或零值。
• 同步接收: 如有发送者，交换数据。否则排队、阻塞。
• 异步接收: 如有缓冲项，拷贝数据。否则排队、阻塞。

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// chan的注意事项和使用小技巧
func notice(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)
	fmt.Println("\n    #1 关闭已经关闭的通道会引起恐慌\n")
	ch1 := make(chan bool)
	close(ch1)
	//	close(ch)  // 这样会panic的，channel不能close两次

	fmt.Println("\n    #2 读取的channel提前关闭了，不会引起恐慌 \n")
	ch2 := make(chan string)
	close(ch2)
	str := <-ch2 // 不会panic, i读取到的值是空 "",  如果channel是bool的，那么读取到的是false
	fmt.Println("str =", str)

	// #3 向已经关闭的通道会引起恐慌
	ch3 := make(chan string)
	close(ch3)
	//	ch3 <- "hello" // 会panic的

	fmt.Println("\n    #4 判断channel是否close\n")
	ch4 := make(chan string)
	close(ch4)
	i, ok := <-ch4
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("通道已经关闭。")
	}

	fmt.Println("\n    #5 for range循环读取channel\n")
	ch5 := make(chan int, 10)
	go func() {
		for i := range ch5 { // ch关闭时，for循环会自动结束
			fmt.Println(i)

		}
		fmt.Println()
	}()
	ch5 <- 1
	ch5 <- 2
	ch5 <- 4
	ch5 <- 8
	close(ch5)

	time.Sleep(time.Millisecond)
	fmt.Println("\n    #6 防止读取超时\n")
	ch6 := make(chan string)
	go func() {
		select {
		case <-time.After(time.Second):
			fmt.Println("读取通道超时。")
		case str := <-ch6:
			fmt.Println(str)
		}
	}()
	time.Sleep(1111 * time.Millisecond)
	//	ch6<-"hello"

	fmt.Println("\n    #7 防止写入超时\n")
	ch7 := make(chan string)
	go func() {
		select {
		case <-time.After(time.Second):
			fmt.Println("写入通道超时。")
		case ch7 <- "hello":
			fmt.Println(str)
		}
	}()
	time.Sleep(1111 * time.Millisecond)
	//	<-ch7
}

// --------------------------------------------------------------------------------------------

// 用channel实现打包并发任务
func NewCh() chan int {
	c := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	go func() {
		time.Sleep(time.Second)
		c <- r.Intn(100)
	}()
	return c
}

func runCh(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)
	c := NewCh()
	fmt.Println("c =", <-c)
}

// --------------------------------------------------------------------------------------------

func handleTask(i int) {
	fmt.Println("正在处理任务", i)
	time.Sleep(300 * time.Millisecond)
}

// 用channel 实现信号量 (semaphore)
func semaphore(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	wg := sync.WaitGroup{}
	sem := make(chan int, 1) // 只缓冲一个数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- 1 // 向 sem 发送数据，阻塞或者成功。
			handleTask(id)
			<-sem // 接收数据，使得其他阻塞 goroutine 可以发送数据。
		}(i)
	}
	wg.Wait()
}

// --------------------------------------------------------------------------------------------

// 用closed channel 发出退出通知
func quitCh(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	var wg sync.WaitGroup
	quit := make(chan bool)
	go func() {
		for i := 1; i < 4; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				for { // 循环执行任务
					select {
					case <-quit: // closed channel不会阻塞，因此可用作退出通知。
						return
					default: // 执行正常任务。
						handleTask(id)
					}
				}
			}(i)
		}
	}()

	// 创造退出执行任务的条件
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		time.Sleep(100 * time.Millisecond)
		if r.Intn(10) == 3 { // 当出现随机数等于3时推出执行任务
			fmt.Println("\n退出执行任务。\n")
			close(quit) // 发出退出通知。
			return
		}
	}

	time.Sleep(time.Second)
	wg.Wait() // 等待任务结束
}

func main() {
	notice("chan的注意事项和使用小技巧")
	runCh("用channel实现打包并发任务")
	semaphore("用channel 实现信号量")
	quitCh("用closed channel 发出退出通知")
}
