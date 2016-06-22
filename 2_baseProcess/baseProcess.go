package main

import (
	"fmt"
)

func ifCon() {
	a, b, c := 13, 14, 8
	if a > b {
		a, b = b, a // a和b交换
	} else if a == b {
		c = (a + b) / 2
	} else {
		fmt.Println("a is less b")
	}
	fmt.Println(a, b, c)
}

func switchCon() {
	x := 21

	switch x {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough // 值为下一个条件
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4, 5, 6")
	default:
		fmt.Println("Default")
	}

	switch { // 条件不是必须的
	case 0 < x && x <= 10: // 不能直接写x
		fmt.Println("0 ~ 10")
	case 10 < x && x <= 20:
		fmt.Println("10 ~ 20")
	case 20 < x && x <= 30:
		fmt.Println("20 ~ 30")
	case 30 < x && x <= 40:
		fmt.Println("30 ~ 40")
	default:
		fmt.Println("unknow x")
	}
}

func forProcess() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	sum = 0
	for {
		sum++
		if sum >= 100 {
			break
		}
	}
	fmt.Println("sum =", sum)
}

func main() {
	fmt.Println("条件判断：")
	ifCon()

	fmt.Println("\nswitch 判断：")
	switchCon()

	fmt.Println("\n for循环：")
	forProcess()
}
