
/*
Go语言内置以下这些基础类型：
 布尔类型： bool
 整型： 	 int8 、 byte 、 int16 、 int 、 uint 、 uintptr 等
 浮点类型： float32 、 float64
 复数类型： complex64 、 complex128
 字符串： 	 string
 字符类型： rune
 错误类型： error

此外，Go语言也支持以下这些复合类型：
 指针（pointer）
 数组（array）
 切片（slice）
 字典（map）
 通道（chan）
 结构体（struct）
 接口（interface）
*/
// 变量名首字母大写为导出型，首字母为小写字母是包内私有。

package main

import (
	"fmt"
)

// 多变量或者多常量可以这样申明
/*const(
    n = 100
    str = "你好，微度网络"
)
var(
    m int
    ui string
    name float32
)*/

func add(a int, b int) (int, error) {
	return a + b, nil
}

func getName() (firstName, lastName, nickName string) {
	return "Chen", "GangSheng", "ChengLong"
}

func main() {

	// 1.变量初始化
	var v1 int
	v1 = 1
	var v2 int = 2
	var v3 = 3
	v4 := 4                                                 // go语言自动从右到左推导出变量类型，默认32位
	v1v, v2v, v3v, v4v := 1, 2.2, 3, "variable declaration" // 多个变量一起声明，而且可以不同类型
	fmt.Println(v1, v2, v3, v4, v1v, v2v, v3v, v4v)         // go自动识别类型来打印
	fmt.Println()                                           // 换行

	// 整形(int)
	/*
		int8  	  1  		128 ~ 127
		uint8 	 1  		0 ~ 255
		int16  	 2  	   32 768 ~ 32 767
		uint16  2 		  0 ~ 65 535
		int32  	 4  	   2 147 483 648 ~ 2 147 483 647
		uint32  4  		 0 ~ 4 294 967 295
		int64  	 8  	  9 223 372 036 854 775 808 ~ 9 223 372 036 854 775 807
		uint64  8  		 0 ~ 18 446 744 073 709 551 615
		int  	平台相关  	平台相关
		uint  	平台相关  	平台相关
		uintptr 同指针		在32位平台下为4字节，64位平台下为8字节
	*/
	var v5 int8 = 5
	var v6 int16 = 6
	var v7 int32 = 7
	var v8 int64 = 8
	fmt.Println(v5, v6, v7, v8)
	fmt.Println() // 换行

	// 2.浮点数(float)
	var v9 float32 = 9.1
	var v10 float64 = 10.2
	fmt.Println(v9, v10)
	fmt.Println() // 换行

	// 3.布尔类型(bool)
	var b1, b2 bool
	b1 = (0 != 1)
	b2 = (0 == 1)
	fmt.Println("1!=0 result:", b1, "   ", "1==0 result:", b2)
	fmt.Println() // 换行

	// 4.字符串(string)
	var str1 string // 字符串
	str1 = "string variable+"
	str2 := "hello golang"
	ch := str2[0]
	fmt.Printf("str1 = %s, str2 = %s, ch = %c\n", str1, str2, ch)
	str3 := str1 + str2
	str3len := len(str3)
	fmt.Println(str3, "   str3len = ", str3len)
	fmt.Println() // 换行

	// 5.数组([]类型)
	/*
	   [8]byte         	//  长度为8的数组，每个元素为一个字节
	   [8]int,[8]int8, [8]int16, [8]int32, [8]int64
	   [2*N] struct { x, y int32 } //  复杂类型数组
	   [1000]*float64      //  指针数组
	   [3][5]int        	//  二维数组
	*/
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Elements of array1: ")
	for _, v := range array1 {
		fmt.Print(v, " ")
	}
	fmt.Println() // 换行

	array2 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	for i := 0; i < len(array2); i++ {
		fmt.Println("Array2 element[", i, "]=", array2[i])
	}

	str_arr := [5]string{"zhangfei", "guanyu", "liubei", "zhaoyun", "huagzhong"}
	str_arr[2] = "jiangwei"
	fmt.Println(str_arr)
	fmt.Println() // 换行

	// 6.复数类型(complex)
	var value1 complex64 //  由 2 个 float32 构成的复数类型
	value1 = 2.1 + 12.1i
	value2 := 4.2 + 24.2i        // value2 是 complex128 类型
	value3 := complex(8.4, 48.4) // value3 结果同  value2
	fmt.Printf("x1=%f, y1=%f, x2=%f, y2=%f, x3=%f, y3=%f\n", real(value1), imag(value1), real(value2), imag(value2), real(value3), imag(value3))
	fmt.Println() // 换行

	// 7.错误类型(error)
	sum, err := add(5, 6)
	if err != nil {
		fmt.Println("run error!")
	} else {
		fmt.Println("a+b =", sum)
	}
	fmt.Println() // 换行

	// 8.匿名变量
	_, _, nickName := getName() // 下划线是匿名变量，这种用法可以让代码非常清晰，基本上屏蔽掉了可能混淆代码阅读者视线的内容，只获取有用的信息
	fmt.Println("nickName:", nickName)
	fmt.Println() // 换行

	// 9.常量(const)
	const cn int = 100
	const Pi float64 = 3.14159265358979323846
	const STR string = "const string"
	fmt.Println(cn, Pi, STR)
	fmt.Println() // 换行

	// 10.预定义常量（true、false、iota）
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
	fmt.Println() // 换行

	// 12.枚举
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays //  这个常量没有导出, 为包内私有。
	)
	fmt.Println("Sunday =", Sunday, " Monday =", Monday, " Tuesday =", Tuesday, " Wednesday =", Wednesday, " Thursday =", Thursday, " Friday =", Friday, " Saturday =", Saturday, " numberOfDays =", numberOfDays)
}
