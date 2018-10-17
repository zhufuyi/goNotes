// slice的初始化和追加元素
/*
slice 是一种可变长的动态数组，属于引用类型
slice底层由三部分构成，分别是指向底层数组的指针、slice中元素的长度、slice 的容量(可供增长的最大值)
*/
package main

import (
	"fmt"
)

// slice创建和初始化
func sliceInit(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	var slice1 []int               // 声明一个len=0 cap=0的slice，此时slice1=nil
	slice2 := []int{1, 2, 3, 4, 5} // 声明一个len和cap都为5的slice
	// 用make声明slice，第一个参数为类型，第二个参数为slice的长度len（可省略），第三个参数为slice的容量cap，当第二个参数省略时len=cap
	slice3 := make([]string, 10) // 声明一个len=10 cap=10的slice
	slice4 := make([]int, 0, 10) // 声明一个len=0 cap=10的slic，

	//	slice1[0] = 1 	// len(slice1)=0 报错，索引超出范围
	//	slice4[0] = 2	// len(slice4)=0 报错，索引超出范围
	slice2[0] = 100      // len(slice2)=5
	slice3[0] = "golang" // len(slice3)=10
	fmt.Println("slice1 =", slice1, "\nslice2 =", slice2, "\nslice3 =", slice3, "\nslice4 =", slice4)

	// slice的截取,完整表示slice[low:high:max]，有多种写法，各种写法表示意义都相同
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // len=10, cap=10
	slice5 := sl[:]                           // [0,1,2,3,4,5,6,7,8,9], len=10, cap=10,相当于slice5=sl
	slice6 := sl[5:]                          // [5,6,7,8,9], len=5, cap=5，截取第5个元素后所有元素，包括第5个元素
	slice7 := sl[:6]                          // [0,1,2,3,4,5], len=6, cap=6，截取第6个元素之前的所有元素，不包括第6个元素
	slice8 := sl[2:6]                         // [2,3,4,5], len=4, cap=8，截取第2到第6个元素
	slice9 := sl[2:6:6]                       // [4,5], len=2, cap=6，截取第2到第6个元素，但限定cap=6
	fmt.Println("slice5 =", slice5, "\nslice6 =", slice6, "\nslice7 =", slice7, "\nslice8 =", slice8, "\nslice9 =", slice9)

	// 结果：
	// slice1 = []
	// slice2 = [100 2 3 4 5]
	// slice3 = [golang         ]
	// slice4 = []
	// slice5 = [0 1 2 3 4 5 6 7 8 9]
	// slice6 = [5 6 7 8 9]
	// slice7 = [0 1 2 3 4 5]
	// slice8 = [2 3 4 5]
	// slice9 = [2 3 4 5]
}

func chgSlice(sl []int) {
	// 追加元素，当所有元素总和超过cap时，cap自动翻倍，并且系统重新自动分配一块连续的内存，所以内存地址已经改变了，和原来最初的地址数据无关了
	sl = append(sl, 6, 7) // 追加元素
	if len(sl) > 3 {
		sl[3] = 333
		fmt.Printf("sl = %v, %p, %d, %d\n", sl, sl, len(sl), cap(sl))
	}
}

// slice追加元素
func sliceAppend(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	sl1 := []int{0, 1, 2, 3, 4, 5}
	chgSlice(sl1)
	fmt.Printf("sl1 = %v, %p, %d, %d\n\n", sl1, sl1, len(sl1), cap(sl1))

	sl2 := make([]int, 10)
	sl2 = append(sl2[:0], sl1...) // 当追加的是相同类型slice时，slice后面需要另外加个尾巴（...）,表示把slice打散为独立元素
	chgSlice(sl2)
	fmt.Printf("sl2 = %v, %p, %d, %d\n", sl2, sl2, len(sl2), cap(sl2))

	// 结果：
	// sl = [0 1 2 333 4 5 6 7], 0xc082066000, 8, 12
	// sl1 = [0 1 2 3 4 5], 0xc082062030, 6, 6

	// sl = [0 1 2 333 4 5 6 7], 0xc082050190, 8, 10
	// sl2 = [0 1 2 333 4 5], 0xc082050190, 6, 10
}

func main() {
	sliceInit("slice创建和初始化")
	sliceAppend("追加元素")
}
