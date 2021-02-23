package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Job(i int) {
	time.Sleep(time.Duration(100+rand.Intn(700)) * time.Millisecond) // 随机延时，模拟执行任务时间
	fmt.Printf("job number is %d\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go Job(i) // 并发执行任务
	}

	time.Sleep(time.Second)
}
