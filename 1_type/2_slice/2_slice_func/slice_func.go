// 切片的基本处理函数有
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// 比较两个slice顺序
func compare(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice1 := []byte("golang")
	slice2 := []byte("golang")
	slice3 := []byte("goLang...")

	// 大于0表示第一个在第二个slice顺序之后，小于0表示之前，等于0表示相同
	fmt.Println(bytes.Compare(slice1, slice2), bytes.Compare(slice2, slice3), bytes.Compare(slice3, slice1))

	// 结果：0 1 -1
}

// 判断两个slice是否相等
func equal(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice1 := []byte("golang")
	slice2 := []byte("Golang")
	fmt.Println(bytes.Equal(slice1, slice2), bytes.EqualFold(slice1, slice2)) // 区分大小写和不区分大小写

	// 结果：false  true
}

// 判断slice是否存在指定的前后缀
func preOrSuf(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("hello golang slice")
	prefix := []byte("golang")
	suffix := []byte("slice")
	fmt.Println(bytes.HasPrefix(slice, prefix), bytes.HasSuffix(slice, suffix)) // 前缀和后缀的判断

	// 结果：false  true
}

// 判断一个slice是否是另一个slice的子集
func contains(info string) {
	fmt.Printf("\n\n-------------------- %s --------------------\n", info)

	slice1 := []byte("goLang")
	slice2 := []byte("lan")
	slice3 := []byte("Lan")
	fmt.Println(bytes.Contains(slice1, slice2), bytes.Contains(slice1, slice3))

	// 结果：false  true
}

// 转换为URF-8编码
func runes(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("hello golang, 中国！")
	fmt.Println("转换前len =", len(slice))
	fmt.Println("转为后len =", len(bytes.Runes(slice))) // 把字符串转为UTF-8编码

	// 结果：
	// 转换前len = 23
	// 转为后len = 17
}

// 计算一个切片在另一个切片包含的个数
func count(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("banana")
	slice1 := []byte("an")
	slice2 := []byte("bn")
	fmt.Println("[\"banana\", \"ba\"]", bytes.Count(slice, slice1))
	fmt.Println("[\"banana\", \"bn\"]", bytes.Count(slice, slice2))

	// 结果:
	// ["banana", "ba"] 2
	// ["banana", "bn"] 0
}

// 获取字符或slice在另一个slice中第一次出现或最后一次出现的索引
func indexOrLast(info string) {
	fmt.Printf("\n\n---------- %s ----------\n", info)

	slice1 := []byte("01234567890123456789")
	fmt.Println(bytes.Index(slice1, []byte("345")), bytes.IndexByte(slice1, '3'), bytes.IndexAny(slice1, "345"))             // 第一次出现的索引
	fmt.Println(bytes.LastIndex(slice1, []byte("345")), bytes.LastIndexByte(slice1, '3'), bytes.LastIndexAny(slice1, "345")) // 最后一次出现

	slice2 := []byte("。。。哈哈。。。哈哈。。。")
	fmt.Println(bytes.IndexRune(slice2, '哈')) // 第一次出现的索引
	f := func(r rune) bool {
		if r == '哈' {
			return true
		} else {
			return false
		}
	}
	fmt.Println(bytes.LastIndexFunc(slice2, f)) // 最后一次出现的索引

	// 结果：
	// 3 3 3
	// 13 13 15
	// 9
	// 27
}

// 大小写转换
func lowerOrUpper(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("aBcDeFg")

	fmt.Println("转为小写：", string(bytes.ToLower(slice)))
	fmt.Println("转为大写：", string(bytes.ToUpper(slice)))

	// 结果：
	// 转为小写： abcdefg
	// 转为大写： ABCDEFG
}

// 重复复制切片n次
func repeat(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("^.^")
	fmt.Println("哈" + string(bytes.Repeat(slice, 5)))

	// 结果：哈^.^^.^^.^^.^^.^
}

// 替换切片匹配的内容
func replace(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	str := []byte("研究研究golang")
	oldStr := []byte("研究")
	newStr := []byte("学习")

	fmt.Println(string(bytes.Replace(str, oldStr, newStr, 1)))  // 如果存在，替换第一个匹配的slice
	fmt.Println(string(bytes.Replace(str, oldStr, newStr, -1))) // 如果存在，替换所有匹配的slice

	// 结果
	// 学习研究golang
	// 学习学习golang
}

// 去掉slice内容
func trim(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice1 := []byte("研究研究golang")
	fmt.Println(string(bytes.Trim(slice1, "lang"))) // 去掉所有匹配的字符串

	slice2 := []byte("        golang")
	fmt.Println(string(bytes.TrimSpace(slice2))) // 去掉左边的空格

	f := func(r rune) bool {
		if r == '研' {
			return true
		}
		return false
	}
	fmt.Println(string(bytes.TrimFunc(slice1, f))) // 去掉左边匹配的unicode

	slice3 := []byte("hello golang")
	fmt.Println(string(bytes.TrimLeft(slice3, "he")))             // 去掉左边匹配的字符串
	fmt.Println(string(bytes.TrimPrefix(slice3, []byte("he"))))   // 去掉左边匹配的slice
	fmt.Println(string(bytes.TrimRight(slice3, "lang")))          // 去掉右边匹配的字符串
	fmt.Println(string(bytes.TrimSuffix(slice3, []byte("lang")))) // 去掉右边匹配的slice

	// 结果：
	// 研究研究go
	// golang
	// 究研究golang
	// llo golang
	// llo golang
	// hello go
	// hello go
}

func printSlice(ss [][]byte) {
	for _, sl := range ss {
		fmt.Println(string(sl))
	}
	fmt.Println()
}

// 以空间隔白拆分slice
func fields(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("hello golang 中国")
	printSlice(bytes.Fields(slice))

	// 结果：
	// hello
	// golang
	// 中国

}

// slice拆分
func split(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	slice := []byte("golang--golang--golang")
	printSlice(bytes.Split(slice, []byte("--"))) // 拆分所有匹配的slice

	printSlice(bytes.SplitN(slice, []byte("--"), 2)) // n>0最多拆分n个匹配的slice，n<0拆分所有匹配的slice，n=0返回nil

	printSlice(bytes.SplitAfter(slice, []byte("--"))) // 拆分所有匹配slice，但保留匹配的slice

	printSlice(bytes.SplitAfterN(slice, []byte("--"), 2)) // n>0最多拆分n个匹配的slice，n<0拆分所有匹配的slice，n=0返回nil，保留匹配的slice

	// 结果：
	// golang
	// golang
	// golang

	// golang
	// golang--golang

	// golang--
	// golang--
	// golang

	// golang--
	// golang--golang
}

// 二维slice以指定的分割符连接为一维的slice
func join(info string) {
	fmt.Printf("\n\n----------------- %s -----------------\n", info)

	sslice := [][]byte{
		[]byte("hello"),
		[]byte("golang"),
	}
	fmt.Println(string(bytes.Join(sslice, []byte("_")))) // 以分割符连接返回一维的slice

	// 结果：hello_golang
}

// 去除重复元素，返回去除重复后的slice和重复的元素
func removeDuplicate(info string) ([]string, []string) {
	fmt.Printf("\n\n----------------- %s -----------------\n", info)
	slice := []string{"a", "b", "c", "a", "d", "e", "c"}
	var x []string
	var duplicate []string
	for _, i := range slice {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					duplicate = append(duplicate, v)
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	fmt.Println(x, "\n", duplicate)
	return x, duplicate

	// 结果：
	// 去除重复后：[a b c d e]
	// 重复的元素：[a c]
}

// 打乱顺序
func randSlice(info string) []string {
	fmt.Printf("\n\n----------------- %s -----------------\n", info)

	slice := []string{"a", "b", "c", "e", "f", "g"}
	rt := rand.New(rand.NewSource(time.Now().UnixNano()))
	randV := rt.Perm(len(slice))
	randSlice := make([]string, len(slice))
	for k, _ := range slice {
		randSlice[k] = slice[randV[k]]
	}
	fmt.Println(randSlice)

	return randSlice

	// 结果：
	// 随机一个结果：[g c e f a b]
}

func main() {
	compare("比较两个slice的顺序")
	equal("判断两个slice是否相等")
	preOrSuf("判断slice是否存在指定的前后缀")
	contains("判断一个slice是否是另一个slice的子集")
	runes("转换为URF-8编码")
	count("计算一个切片在另一个切片重复的次数")
	indexOrLast("获取字符或slice在另一个slice中第一次出现或最后一次出现的索引")
	lowerOrUpper("大小写转换")
	repeat("重复复制切片n次")
	replace("替换切片匹配的内容")
	trim("裁剪slice内容")
	fields("以空白间隔拆分slice")
	split("slice拆分")
	join("二维slice以指定的分割符连接为一维的slice")
	removeDuplicate("去除相同的元素")
	randSlice("打乱slice元素的顺序")
}
