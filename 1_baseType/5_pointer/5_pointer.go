package main

import (
	"fmt"
)

func main() {
	var i int
	var p *int
	var pp **int

	i = 5
	fmt.Println("i =", i)
	p = &i
	fmt.Println("*p =", *p, "  &p =", &p)
	*p++
	fmt.Println("*p =", *p, "  &p = ", &p, "  i =", i)
	pp = &p
	fmt.Println("*pp =", *pp, "  &pp =", &pp, "  i =", i)
	**pp++ // 等价于(**pp)++,即对该内存的内容加1操作，go语言禁止内存指针自动加减操作,++和--只支持整数类型操作
	fmt.Println("*pp =", *pp, "  &pp =", &pp, "  i =", i)
}
