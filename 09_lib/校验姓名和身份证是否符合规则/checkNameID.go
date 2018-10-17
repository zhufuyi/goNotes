package main

import (
	"flag"
	"fmt"
	"strconv"
	"unicode"
)

type checkIDName struct {
	name string
	id   string
}

// 判断身份证是否符合中国身份证算法
func (c *checkIDName) checkId() bool {
	if len(c.id) != 18 {
		fmt.Println("id length != 18")
		return false
	}

	// 判断第18为是否符合规则
	if c.id[17] != 'X' && (c.id[17] < '0' || c.id[17] > '9') {
		fmt.Println("id[18] was forbidden characte.")
		return false
	}

	constNO := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2} // 常量

	// 计算校验和
	idSlice := []byte(c.id)
	sum := 0
	for i, v := range idSlice[:17] {
		no, _ := strconv.Atoi(string(v))
		sum += no * constNO[i]
	}
	checkCode := sum % 11

	constNO2 := [11]byte{1, 0, 'X', 9, 8, 7, 6, 5, 4, 3, 2}
	calcCode := constNO2[checkCode]
	var lastBit byte
	if calcCode != 'X' {
		lastBitInt, _ := strconv.Atoi(string(c.id[17]))
		lastBit = byte(lastBitInt)
	} else {
		lastBit = c.id[17]
	}

	if calcCode != lastBit {
		return false
	}
	return true
}

// 判断名字是否都是汉字
func (c *checkIDName) checkName() bool {
	if c.name == "" {
		return false
	}
	for _, r := range c.name {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

func (c *checkIDName) String() string {
	return fmt.Sprintf("name=%s, id=%s", c.name, c.id)
}

func main() {
	// 执行参数--name xxx --id xxx
	var name, idNumber string
	flag.StringVar(&name, "name", "", "姓名")
	flag.StringVar(&idNumber, "id", "", "身份证号码")
	flag.Parse()

	people := &checkIDName{name, idNumber}
	println(people.String())
	if people.checkName() {
		println("姓名正常")
	} else {
		println("姓名有非法字符")
	}
	if people.checkId() {
		println("身份证号正常")
	} else {
		println("身份证号不合法")
	}
}

//测试号码：34052419800101001X
//测试号码：511028199507215915
