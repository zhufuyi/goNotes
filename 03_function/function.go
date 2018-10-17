package main

import (
	"fmt"
)

func mulReturnVal(a, b int) (int, int) { // 支持多返回，(a, b int)表示参数个数和类型，(int, int)表示返回值类型和个数
	return a + b, a * b
}

func singleTypeMulPara(args ...string) { // 单一类型不定参数
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
}

func anyTypeJustPara(args ...interface{}) { //任意类型不定参数
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case float32:
			fmt.Println(arg, "is an float32 value.")
		case float64:
			fmt.Println(arg, "is an float64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

// 匿名函数，把函数当做一个类型
var areaFuc = func(c float32) float32 {
	return 3.14 * c * c
}

func main() {
	fmt.Println("1 多返回值函数：")
	add, mul := mulReturnVal(5, 6)
	fmt.Printf("add = %d, mul = %d\n", add, mul)

	fmt.Println("\n2 传递同一类型的不定参数函数：")
	singleTypeMulPara("刘备", "关羽", "张飞")

	fmt.Println("\n3 传递不同种类型的不定参数函数：")
	tp1 := 11
	tp2 := 11.11
	tp3 := "hello"
	anyTypeJustPara(tp1, tp2, tp3)

	fmt.Println("\n4 匿名函数：")
	fmt.Println("area =", areaFuc(10))
}
