package interfaceEg

import (
	"fmt"
)

type Xi []int           // 类型1
func (p Xi) Len() int { // 类型1实现方法1
	return len(p)
}
func (p Xi) Less(i, j int) bool { // 类型1实现方法2
	return p[j] < p[i]
}
func (p Xi) Swap(i, j int) { // 类型1实现方法3
	p[i], p[j] = p[j], p[i]
}

type Xs []string        // 类型2
func (p Xs) Len() int { // 类型2实现方法1
	return len(p)
}
func (p Xs) Less(i, j int) bool { // 类型2实现方法2
	return p[j] < p[i] // 可以实现字符串比较
}
func (p Xs) Swap(i, j int) { // 类型2实现方法3
	p[i], p[j] = p[j], p[i]
}

func prints(s string) {
	fmt.Println("massage:", s)
}

/*
其一，你只需要知道这个类实现了哪些方法，每个方法是啥含义就足够了。
其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理。接口由使用方按需定义，而不用事前规划。
其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口。
*/
type Sorter interface { // 定义接口sorter，接口实现Xi Xs的方法
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func SortEg() {
	xi := Xi{32, 42, 4, 2, 5, 123, 1, 42, 34, 75, 67, 8, 65, 67, 567, 212, 6, 34, 3, 76, 88, 9, 90}
	fmt.Println("sort before:", xi)
	sort(xi)
	fmt.Println(" sort after:", xi)
	fmt.Println()

	xs := Xs{"zhangfei", "guanyu", "liubei", "zhaoyun", "jiangwei", "huangzhong"}
	fmt.Println("sort before:", xs)
	sort(xs)
	fmt.Println(" sort after:", xs)
	fmt.Println()
}
