package main

import "fmt"

type Student struct {
	*People
	School string
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
		&People{"张三"},
		"xxx小学",
	}
	fmt.Println(stu.Name, stu.School)
	stu.Say() // 继承了方法

	// 类似多态
	var ws WhoSay

	ws = &People{"李四"}
	ws.Say()

	ws = &Animal{"狗"}
	ws.Say()
}
