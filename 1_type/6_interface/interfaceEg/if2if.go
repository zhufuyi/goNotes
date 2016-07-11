package interfaceEg

import (
	"fmt"
)

//定义上述两个接口的实现类
type MyReadWrite struct{}

func (mrw *MyReadWrite) read() {
	fmt.Println("MyReadWrite...read")
}

func (mrw *MyReadWrite) write() {
	fmt.Println("MyReadWrite...write")
}

//接口中可以组合其它接口，这种方式等效于在接口中添加其它接口的方法
type Reader interface {
	read()
}

type Writer interface {
	write()
}

//定义一个接口，组合了上述两个接口
type ReadWriterV1 interface {
	Reader
	Writer
}

//上述接口等价于：
type ReadWriterV2 interface {
	read()
	write()
}

//ReadWriter和ReadWriterV2两个接口是等效的，因此可以相互赋值
func If2if() {
	mrw := &MyReadWrite{}
	//mrw对象实现了read()方法和write()方法，因此可以赋值给ReadWriter和ReadWriterV2
	var rw1 ReadWriterV1 = mrw
	rw1.read()
	rw1.write()

	fmt.Println("------")
	var rw2 ReadWriterV2 = mrw
	rw2.read()
	rw2.write()

	//同时，ReadWriter和ReadWriterV2两个接口对象可以相互赋值
	rw1 = rw2
	rw2 = rw1
}
