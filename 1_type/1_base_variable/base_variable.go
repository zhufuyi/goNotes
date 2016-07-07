// 常用类型的简单使用
/*
golang内置的基础类型：
 布尔类型： bool
 整型：	 int8 、 byte 、 int16 、 int 、 uint 、 uintptr
 字符串：	 string
 浮点类型：float32 、 float64
 复数类型：complex64 、 complex128
 字符类型：rune
 错误类型：error
   指针：	 pointer

golang内置的复合类型：
 数组：	array
 切片：	slice(引用类型)
 字典：	map(引用类型)
 通道：	chan (引用类型)
 接口：	interface(引用类型)
   结构体：	struct

可以使用基础类型和符合类型组成更复杂的自定义类型，例如map[string][]string。
*/
// 变量名首字母大写为导出型，首字母为小写字母是包内私有。

package main

import (
	"errors"
	"fmt"
)

// 常量
func constType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	const ( // 多个变量或者多个常量可以用括号声明
		Pi   = 3.14
		Ip   = "192.168.1.111"
		Port = 8080
	)
	fmt.Printf("Pi = %.2f, Ip = %s, Port = %d\n", Pi, Ip, Port)

	// iota 比较特殊，可以被认为是一个可被编译器修改的常量，在每一个 const 关键字出现时被重置为0，然后在下一个 const 出现之前，每出现一次 iota ，其所代表的数字会自动增1。
	const ( // 出现const，iota 被重设为0
		d0 = iota // d0 == 0
		d1 = iota // d1 == 1
		d2 = iota // d2 == 2
	)

	const (
		d3 = 1 << iota // d3 == 1 (iota 在每个 const 开头被重设为0 )
		d4 = 1 << iota // d4 == 2
		d5 = 1 << iota // d5 == 4
	)
	fmt.Println("d0 =", d0, " d1 =", d1, " d2 =", d2, " d3 =", d3, " d4 =", d4, " d4 =", d5)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("Sunday =", Sunday, " Monday =", Monday, " Tuesday =", Tuesday, " Wednesday =", Wednesday,
		" Thursday =", Thursday, " Friday =", Friday, " Saturday =", Saturday)

	// 结果：
	// Pi = 3.14, Ip = 192.168.1.111, Port = 8080
	// d0 = 0  d1 = 1  d2 = 2  d3 = 1  d4 = 2  d4 = 4
	// Sunday = 0  Monday = 1  Tuesday = 2  Wednesday = 3  Thursday = 4  Friday = 5  Saturday = 6
}

/*
类型		字节数		范围
int8  	1  			-128 ～ 127
uint8 	1  			0 ～ 255
int16  	2  	   		-32768 ～ 32767
uint16  2 		  	0 ～ 65 535
int32  	4  	   		-2147483648 ～ 2147483647
uint32  4  			0 ～ 4294967295
int64  	8  	  		-9223372036854775808 ～ 9223372036854775807
uint64  8  		 	0 ～ 18446744073709551615
int  	平台相关  	平台相关
uint  	平台相关  	平台相关
uintptr 同指针		在32位平台下为4字节，64位平台下为8字节
*/
// 整形变量
func integerType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	// 4种声明方式
	var v1 int
	v1 = 1
	var v2 int = 1
	var v3 = 1
	v4 := 1 // 常用局部变量声明方式，不能用来声明全局变量， go语言自动从右到左推导出变量类型

	var v5 int8 = 127
	var v6 uint8 = 255
	var v7 int16 = 32767
	var v8 uint16 = 65535
	var v9 int32 = 2147483647
	var v10 uint32 = 4294967295
	var v11 int64 = 9223372036854775807
	var v12 uint64 = 18446744073709551615
	var v13 byte = 'g'
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13)

	// 结果：1 1 1 1 127 255 32767 65535 2147483647 4294967295 9223372036854775807 18446744073709551615 103
}

// 布尔类型
func boolType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	var b1, b2 bool // 默认为false
	b1 = (0 != 1)
	b2 = (0 == 1)
	fmt.Println("b1 =", b1, "  b2 =", b2)

	// 结果：b1 = true   b2 = false
}

// 浮点类型
func floatType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	var pi float32 = 3.14159
	var v float64 = 10.0 / 3

	fmt.Printf("%.2f, %.3f\n", pi, v) // %.2f表示是保留2位小数点

	// 结果：3.14, 3.333
}

func stringType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	var str1 string // 字符串
	str1 = "hello golang"
	str2 := "中国"
	str3 := str1 + ", " + str2

	ch := str1[0]
	runeT := rune('人') // rune类型
	str4 := str2 + string(runeT)

	fmt.Printf("str1 = %s\nstr2 = %s\nstr3 = %s\nstr4 = %s\n", str1, str2, str3, str4)
	fmt.Printf("ch = %c", ch)

	// 结果：
	// str1 = hello golang
	// str2 = 中国
	// str3 = hello golang, 中国
	// str4 = 中国人
	// ch1 = h
}

// 复数类型
func complexType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	var cpl1 complex64 = 1.1 + 2.2i //  由 2 个 float32 构成的复数类型
	var cpl2 complex128 = 3.3 + 4.4i
	cpl3 := complex(5.5, 6.6)
	fmt.Printf("x1=%f, y1=%f, x2=%f, y2=%f, x3=%f, y3=%f\n", real(cpl1), imag(cpl1), real(cpl2), imag(cpl2), real(cpl3), imag(cpl3))

	// 结果：x1=1.100000, y1=2.200000, x2=3.300000, y2=4.400000, x3=5.500000, y3=6.600000
}

// 数组类型
func arrayType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	array3 := [5]string{"张飞", "关羽", "刘备", "赵云", "黄忠"}

	array1[2] = 33333
	array2[0][2] = 3333
	array3[2] = "曹操"

	fmt.Printf("array1 = %v\narray2 = %v\narray3 = %q\v", array1, array2, array3)

	// 结果：
	// array1 = [1 2 33333 4 5]
	// array2 = [[1 2 3333 4] [5 6 7 8]]
	// array3 = ["张飞" "关羽" "曹操" "赵云" "黄忠"]
}

func division(x int, y int) (int, error) { // 除法
	if y == 0 {
		return 0, errors.New("除数不能为0")
	}
	return x / y, nil
}

// 错误类型
func errorType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	x, y := 18, 0
	result1, err := division(x, y)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d/%d = %d\n", x, y, result1)
	}

	y = 3
	result2, err := division(x, y)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d/%d = %d\n", x, y, result2)
	}

	// 结果：
	// 除数不能为0
	// 18/3 = 6
}

func getRGB() (red, green, blue uint8) {
	return 255, 0, 255 // 紫色
}

// 匿名变量
func anonymousType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	red, _, _ := getRGB() // 下划线是匿名变量，从返回值中忽略不关心匿名变量，只获取有用的信息
	_, _, blue := getRGB()

	fmt.Println("red =", red, ", blue =", blue)

	// 结果：red = 255 , blue = 255
}

// 指针
func pointerType(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	x := 1
	px := &x
	*px = 10
	fmt.Printf("x=%d, px=%p\n", x, px)

	ppx := &px
	**ppx = 100
	fmt.Printf("x=%d, *ppx=%p, ppx=%p\n", x, *ppx, ppx)

	// 结果：
	// x=10, px=0xc0820365b0
	// x=100, *ppx=0xc0820365b0, ppx=0xc082056020

}

func main() {
	constType("常量")
	integerType("整型")
	boolType("布尔类型")
	floatType("浮点类型")
	stringType("字符串")
	complexType("复数类型")
	arrayType("数组类型")
	errorType("错误类型")
	anonymousType("匿名变量")
	pointerType("指针类型")
}
