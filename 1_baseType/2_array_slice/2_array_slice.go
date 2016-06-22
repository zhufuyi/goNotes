package main

import (
	"fmt"
)

func main() {
	// 1.数组切片
	var myArray [11]int = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11} //原始数组
	fmt.Println("myArray:", myArray)
	var mySlice1 []int = myArray[:5] // 截取数组前5个元素
	fmt.Println("mySlice1:", mySlice1)
	var mySlice2 []int = myArray[5:] // 截取数组第5个元素之后的元素
	fmt.Println("mySlice2:", mySlice2)
	var mySlice3 []int = myArray[5:8] // 截取第6到第8元素
	fmt.Println("mySlice3:", mySlice3)

	// 2.当没有事先创建数组时，可以通过内置函数 make() 灵活地创建数组切片
	mySlice4 := make([]int, 20) // 创建一个初始元素个数为20的数组切片，元素初始值为0
	fmt.Println("mySlice4:", mySlice4)
	mySlice5 := []int{1, 2, 3, 4, 5} //直接创建并初始化包含5个元素的数组切片(其中会创建一个不需要关心匿名数组被)
	fmt.Println("mySlice5:", mySlice5)
	mySlice6 := make([]int, 5, 10)                   // 创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	fmt.Println("mySlice6 len =", len(mySlice6))     // 切片实际长度
	fmt.Println("mySlice6 caption =", cap(mySlice6)) // 数组切片分配的空间大小

	// 3.数组切片追加新增元素（可以超过原分配空间大小，但建议一次分配够用的空间，可以提高效率，达到空间换时间的效果）
	mySlice6 = append(mySlice6, 1, 2, 3) // 追加元素
	fmt.Println("append mySlice6:", mySlice6)
	mySlice4 = append(mySlice5, mySlice6...) // 追加数字切片，第二个参数后面需要加上省略号，相当于把 mySlice2 包含的所有元素打散后传入。
	fmt.Println("append mySlice4:", mySlice4)

	// 4.数组切片之间内容复制(如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制)
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	slice3 := []int{9, 10, 11, 12}
	copy(slice1, slice2)
	fmt.Println("copy slice2 to slice1:", slice1)
	copy(slice2, slice3)
	fmt.Println("copy slice3 to slice2:", slice2)
}
