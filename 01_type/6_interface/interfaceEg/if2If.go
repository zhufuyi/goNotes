// 接口赋值或隐式转换

package interfaceEg

import (
	"fmt"
)

//定义上述两个接口的实现类
type MyRW struct{}

func (mrw *MyRW) read() string {
	return "read method"
}

func (mrw *MyRW) write() string {
	return "write method"
}

//接口中可以组合其它接口，这种方式等效于在接口中添加其它接口的方法
type Reader interface {
	read() string
}

type Writer interface {
	write() string
}

// 定义一个接口，组合了上述两个接口
type ReadWriterV1 interface {
	Reader
	Writer
}

// 定义一个接口,拥有read和write方法
type ReadWriterV2 interface {
	read() string
	write() string
}

func handleR(r Reader) {
	fmt.Printf("    Reader接口：%s\n", r.read())
}

func handleW(w Writer) {
	fmt.Printf("    Writer接口：%s\n", w.write())
}

func handleRW1(rw ReadWriterV1) {
	fmt.Printf("    ReadWriterV1接口：%s, %s\n", rw.read(), rw.write())
}

func handleRW2(rw ReadWriterV2) {
	fmt.Printf("    ReadWriterV2接口：%s, %s\n", rw.read(), rw.write())
}

func handleNull(t interface{}) {
	fmt.Println("    t =", t)
}

func If2if(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	myrw := &MyRW{}

	fmt.Println("(1) myrw对象有read()方法和write()方法，因此可以赋值给Reader、Writer、ReadWriterV1、ReadWriterV2接口")
	handleR(myrw)
	handleW(myrw)
	handleRW1(myrw)
	handleRW2(myrw)

	fmt.Println("\n(2) ReadWriterV1和ReadWriterV2两个接口是等效的，因此可以相互赋值")
	var rwV1 ReadWriterV1 = myrw
	handleRW2(rwV1)
	var rwV2 ReadWriterV2 = myrw
	handleRW1(rwV2)

	fmt.Println("\n(3) ReadWriterV1和ReadWriterV2两个接口的子接口有Reader、Writer，赋值时会隐式转换为子接口")
	handleR(rwV1)
	handleW(rwV1)
	handleR(rwV2)
	handleW(rwV2)

	fmt.Println("\n(4) 任何类型都实现空接口")
	handleNull(1)
	handleNull(3.14)
	handleNull("hello golang")
	handleNull([]int{1, 2, 3, 4, 5})
	handleNull(myrw)
	handleNull(rwV1)
	handleNull(rwV2)

	/*
	   结果：
	   (1) myrw对象有read()方法和write()方法，因此可以赋值给Reader、Writer、ReadWriterV1、ReadWriterV2接口
	       Reader接口：read method
	       Writer接口：write method
	       ReadWriterV1接口：read method, write method
	       ReadWriterV2接口：read method, write method

	   (2) ReadWriterV1和ReadWriterV2两个接口是等效的，因此可以相互赋值
	       ReadWriterV2接口：read method, write method
	       ReadWriterV1接口：read method, write method

	   (3) ReadWriterV1和ReadWriterV2两个接口的子接口有Reader、Writer，赋值时会隐式转换为子接口
	       Reader接口：read method
	       Writer接口：write method
	       Reader接口：read method
	       Writer接口：write method

	   (4) 任何类型都实现空接口
	       t = 1
	       t = 3.14
	       t = hello golang
	       t = [1 2 3 4 5]
	       t = &{}
	       t = &{}
	       t = &{}
	*/
}
