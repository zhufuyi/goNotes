package main

import "fmt"

type Student struct {
	*People
	Age int
}

type WhoSay interface {
	Say()
}

type People struct {
	Name string
}

func (p *People) Say() {
	fmt.Printf("%s say hello\n", p.Name)
}

type Animal struct {
	Name string
}

func (a *Animal) Say() {
	fmt.Printf("%s say hello\n", a.Name)
}

func main() {
	// 类似继承
	stu := &Student{
		&People{"Zhangsan"},
		11,
	}
	fmt.Println(stu.Name, stu.Age)
	stu.Say() // 继承了方法

	// 类似多态
	var ws WhoSay

	ws = &People{"Lisi"}
	ws.Say()

	ws = &Animal{"Dog"}
	ws.Say()
}
