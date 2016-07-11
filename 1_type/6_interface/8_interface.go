package main

import (
	"./interfaceEg"
	"fmt"
)

func main() {
	fmt.Println("同名接口：")
	interfaceEg.SortEg()

	fmt.Println("\n嵌入接口：")
	interfaceEg.SubStuctEg()

	fmt.Println("\n接口赋值：")
	interfaceEg.If2if()
}
