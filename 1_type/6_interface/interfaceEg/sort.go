package interfaceEg

import (
	"fmt"
	"sort"
)

type Xi []int // 类型1
func (p Xi) Len() int {
	return len(p)
}
func (p Xi) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Xi) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Xs []string // 类型2
func (p Xs) Len() int {
	return len(p)
}
func (p Xs) Less(i, j int) bool {
	return p[i] < p[j] // 可以实现字符串比较
}
func (p Xs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func prints(s string) {
	fmt.Println("massage:", s)
}

func SortEg(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	fmt.Println("int类型排序")
	xi := Xi{32, 42, 4, 2, 5, 123, 1, 42, 34, 75, 67, 8, 65, 67, 567, 212, 6, 34, 3, 76, 88, 9, 90}
	fmt.Println("sort before:", xi)
	sort.Sort(xi) // xi类型实现了sort包的interface接口(Len,Less,Swap方法)，可以传参过去
	fmt.Println("sort after:", xi)

	fmt.Println("\n字符串类型排序")
	xs := Xs{"zhangfei", "guanyu", "liubei", "zhaoyun", "jiangwei", "huangzhong"}
	fmt.Println("sort before:", xs)
	sort.Sort(xs) // xs类型实现了sort包的interface接口(Len,Less,Swap方法)，可以传参过去
	fmt.Println("sort after:", xs)

	/*
		结果：
		int类型排序
		sort before: [32 42 4 2 5 123 1 42 34 75 67 8 65 67 567 212 6 34 3 76 88 9 90]
		sort after: [1 2 3 4 5 6 8 9 32 34 34 42 42 65 67 67 75 76 88 90 123 212 567]

		字符串类型排序
		sort before: [zhangfei guanyu liubei zhaoyun jiangwei huangzhong]
		sort after: [guanyu huangzhong jiangwei liubei zhangfei zhaoyun]
	*/
}
