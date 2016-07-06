/*
Go语言的结构体（struct）和其他语言的类（class）有同等的地位，但Go语言放弃了包括继
承在内的大量面向对象特性，只保留了组合（composition）这个最基础的特性。
*/
package main

import (
	"fmt"
)

type Cuboid struct {
	length, width, height float64
}

func (a *Cuboid) Area() float64 { // 定义Cuboid结构体类型的方法Area，该方法接受者是Cuboid类型
	return 2 * (a.length*a.height + a.length*a.width + a.width*a.height)
}

func (v *Cuboid) Volume() (vl float64, err error) { // 定义Cuboid结构体类型的方法Area，该方法接受者是Cuboid类型，多个返回值写法
	return v.length * v.height * v.length, err // 对应返回值个数
}

type D2 struct {
	x, y float64
}

type D3 struct {
	D2
	z float64
}

func (m *D3) Mul() float64 { // 定义D3结构体类型的方法Mul，该方法接受者是D3类型
	return m.x * m.y * m.z
}

func main() {
	// 初始化方法
	cb1 := new(Cuboid)
	cb2 := &Cuboid{}
	cb3 := &Cuboid{3, 4, 5}
	cb4 := &Cuboid{length: 6, width: 7, height: 8} // 指向结构体的指针，函数传参时直接使用
	cb5 := Cuboid{length: 6, width: 7, height: 8}  // 结构体，函数传参时前面需要加取址符号&

	// 调用
	cb1.length = 1 // 给结构体的元素赋值
	cb1.width = 2
	cb1.height = 3
	fmt.Println("cb1 =", cb1)
	fmt.Println("cb2 =", cb2)
	ar := cb3.Area() // 结构体Cuboid的Area方法调用
	fmt.Println("Area =", ar)
	vl, _ := cb4.Volume() // 结构体Cuboid的Volume方法调用
	fmt.Println("Volume =", vl)
	fmt.Println("cb5", cb5)

	// 结构体内置时可以直接外层结构体直接到具体元素，如果嵌入式的结构没有数据结构的名字 就默认是类型名字D2:D2
	d1 := D3{D2: D2{x: 3, y: 4}, z: 5} // 初始化
	d2 := D2{1, 2}                     // 初始化
	d3 := D3{d2, 3}
	fmt.Println("d1 =", d1, " d3 =", d3)
	mul := d3.Mul()
	fmt.Println("mul =", mul)
}
