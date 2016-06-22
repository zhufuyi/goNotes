// 切片的基本处理函数有
package main

import (
	"bytes"
	"fmt"
)

// Contains函数，判断切片是否包含在另一个切片中
func bytes_Contains() {
	s := []byte("Golang")
	s1 := []byte("go")
	s2 := []byte("Go")
	fmt.Println("[\"Golang\", \"Go\"]", bytes.Contains(s, s1))
	fmt.Println("[\"Golang\", \"go\"]", bytes.Contains(s, s2))
}

// Count函数，计算一个切片在另一个切片包含的个数
func bytes_Count() {
	s := []byte("banana")
	s1 := []byte("ba")
	s2 := []byte("na")
	s3 := []byte("a")
	s4 := []byte("123")
	fmt.Println("[\"banana\", \"ba\"]", bytes.Count(s, s1))
	fmt.Println("[\"banana\", \"na\"]", bytes.Count(s, s2))
	fmt.Println("[\"banana\", \"a\"]", bytes.Count(s, s3))
	fmt.Println("[\"banana\", \"123\"]", bytes.Count(s, s4))
}

// Map函数，把s转为UTF-8编码，使用mapping函数把s中所有字符串映射成对应的字符，把映射结果存放到一个新的切片中
func bytes_Map() {
	s := []byte("hello，张三。")
	m := func(r rune) rune { // 匿名函数
		if r == '张' {
			r = '李'
		}
		if r == '三' {
			r = '四'
		}
		return r
	}
	fmt.Println(string(s))
	fmt.Println(string(bytes.Map(m, s)))
}

// 重复复制切片n次
func bytes_Repeat() {
	s := []byte("哈")
	r := []byte("。")
	fmt.Println(string(s))
	fmt.Println(string(s) + string(bytes.Repeat(r, 6)))
}

// 替换切片指定的所有元素(n<0)
func bytes_Replace() {
	str := []byte("安安静静")
	ol := []byte("静")
	ne := []byte("稳")
	fmt.Println(string(str))
	fmt.Println(string(bytes.Replace(str, ol, ne, 1))) // 替换切片指定元素n个(n>0)
	fmt.Println(string(bytes.Replace(str, ol, ne, -1)))
}

func bytes_Runes() {
	s := []byte("今天天气不错")
	fmt.Println("没有转换长度 =", len(s))
	fmt.Println("转为UTF-8后长度 =", len(bytes.Runes(s))) // 把字符串转为UTF-8编码
}

func main() {
	bytes_Contains()
	bytes_Count()
	bytes_Map()
	bytes_Repeat()
	bytes_Replace()
	bytes_Runes()
}
