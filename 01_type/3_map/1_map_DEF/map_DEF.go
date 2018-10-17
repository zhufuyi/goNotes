// map的声明和使用
// map是一种无序的键值对的集合，通过key来快速检索数据，每次遍历的顺序都可能不一样。

package main

import (
	"fmt"
)

type PeopleInfo struct {
	Name   string // 名字
	Gender string // 性别
	Age    uint8  // 年龄
}

var piMap map[int]*PeopleInfo

func printMap() {
	for k, v := range piMap { // 遍历
		fmt.Printf("%d : %v\n", k, v)
	}
	fmt.Println()
}

// map创建和初始化
func mapInit(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	piMap = make(map[int]*PeopleInfo) // 可以省略容量
	piMap[1] = &PeopleInfo{"刘备", "男", 42}
	piMap[2] = &PeopleInfo{"貂蝉", "女", 24}
	piMap[3] = &PeopleInfo{"周瑜", "男", 33}
	printMap()

	// 结果：
	// 1 : &{刘备 男 42}
	// 2 : &{貂蝉 女 24}
	// 3 : &{周瑜 男 33}
}

// 查找、修改和删除键值
func fcdMap(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	if pi, ok := piMap[1]; ok { // 判断key是否存在
		if pi.Name == "刘备" {
			pi.Name = "曹操" // 修改名字
		}
		printMap()
	}

	if _, ok := piMap[2]; ok { // 判断key是否存在
		piMap[2] = &PeopleInfo{"小乔", "女", 26} // 修改整个结构体信息
		printMap()
	}

	delete(piMap, 10) // 删除键值，即使键值不存在也不会有任何影响
	delete(piMap, 1)
	printMap()

	// 结果：
	// 1 : &{曹操 男 42}
	// 2 : &{貂蝉 女 24}
	// 3 : &{周瑜 男 33}

	// 1 : &{曹操 男 42}
	// 2 : &{小乔 女 26}
	// 3 : &{周瑜 男 33}

	// 2 : &{小乔 女 26}
	// 3 : &{周瑜 男 33}
}

func main() {
	mapInit("map创建和初始化")
	fcdMap("查找、修改和删除键值")
}
