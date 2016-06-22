package main

import (
	"./rwEg"
	"fmt"
)

func main() {
	fmt.Println("------读写文件：")
	rwEg.ReadWriteFile()

	fmt.Println("\n\n------ 按行读取文件内容：")
	rwEg.ReadLineContent()

	fmt.Println("\n\n------ 使用第三方包编辑ini配置文件：")
	rwEg.RwIniFile()

	fmt.Println("\n\n------ 比较三种读取文件的速度：")
	rwEg.ReadRateCompare()
}
