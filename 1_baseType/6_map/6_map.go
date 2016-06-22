/*
map 是一堆键值对的未排序集合
*/

package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	// 声明
	//	personDB := make(map[string]PersonInfo)	// 声明方式一
	personDB := map[string]PersonInfo{ // 声明方式二，声明并初始化，如果不初始化，{}不可省略
		"800004": PersonInfo{"800004", "Jiangwei", "Room 304"},
	}
	// 插入数据
	personDB["800001"] = PersonInfo{"800001", "Zhangfei", "Room 301"}
	personDB["800002"] = PersonInfo{"800002", "Guanyu", "Room 302"}
	personDB["800003"] = PersonInfo{"800003", "Liubei", "Room 303"}
	fmt.Println(personDB)

	// 从这个map查找是否有对应键的信息
	key := "800001"
	person, ok := personDB[key]
	if ok {
		fmt.Println("found person with", key, ", all information: ", person.Name, person.ID, person.Address)
	} else {
		fmt.Println("Dit not found person with", key)
	}

	// 删除一个键值（键值不存在也不会有任何影响）
	fmt.Println("Delete a key", key)
	delete(personDB, key)
	fmt.Println(personDB)
}
