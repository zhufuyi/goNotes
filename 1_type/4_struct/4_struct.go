// 结构体的声明和使用

/*
任何数据类型都可以作为结构体成员的类型，包括函数类型。

从面相对象看，结构体成员是对象的属性，结构体关联的方法是对象的动作行为。

结构体嵌入字段比接口嵌入类型更暴力，嵌入类型所有导出的成员和方法都无条件成为被嵌入结构体的成员和方法。
*/
package main

import (
	"fmt"
)

// 人的信息
type People struct {
	Name    string // 名字
	Gender  string // 性别
	Age     uint8  // 年龄
	Addr    string // 地址
	*Family        // 家庭成员（嵌入字段）
	*School        // 学校信息（嵌入字段）
}

// 人的睡觉行为
func (ppl *People) Sleep() {
	fmt.Println(ppl.Name + "每晚10点准时睡觉。")
}

// 人的吃饭行为
func (ppl *People) Eat(name string) {
	fmt.Println(ppl.Name + "和" + name + "在中秋一起吃团圆饭。")
}

// 人的输出期末考试成绩行为
func (ppl *People) GetScore() {
	fmt.Println(ppl.School.GetScore())
}

// 家庭成员
type Family struct {
	Dad string
	Mom string
}

// 家庭成员的输出信息行为
func (fml *Family) GetFamily() string {
	return fmt.Sprintf("%s、%s", fml.Dad, fml.Mom)
}

// 家庭成员的吃饭行为
func (fml *Family) Eat(name string) {
	fmt.Println(fml.GetFamily() + "和" + name + "在中秋一起吃团圆饭。")
}

// 学校信息
type School struct {
	SchoolName  string
	Headteacher string
	Grade       string
	*Score      // 成绩信息（嵌入字段）
}

// 学校输出信息行为
func (sch *School) PrintInfo() string {
	return fmt.Sprintf("学校：%d， 班主任：%d， 年级：%d， 分数：%s",
		sch.SchoolName, sch.Headteacher, sch.Grade, sch.GetScore())
}

// 成绩信息
type Score struct {
	Chinese uint8
	Math    uint8
	English uint8
	Music   uint8
	Sports  uint8
}

// 成绩的输出信息行为
func (sc *Score) GetScore() string {
	return fmt.Sprintf("语文：%d， 数学：%d， 英语：%d， 音乐：%d， 体育：%d",
		sc.Chinese, sc.Math, sc.English, sc.Music, sc.Sports)
}

func main() {
	// 结构体实例化时比较常用方式的是指针方式，避免作为参数时拷贝整个结构体的副本

	sc := new(Score) // 用new创建
	sc.Chinese = 98
	sc.English = 90
	sc.Math = 100
	sc.Music = 88
	sc.Sports = 99

	sch := &School{} // 用&创建
	sch.SchoolName = "三国小学"
	sch.Headteacher = "诸葛亮"
	sch.Grade = "5(2)班"
	sch.Score = sc

	fml := &Family{"周瑜", "小乔"} // 用&创建并初始化

	ppl := &People{"周循 ", "男", 11, "厦门", fml, sch} //实例化

	// People的方法
	ppl.Sleep()
	ppl.Eat(ppl.GetFamily())
	ppl.Family.Eat(ppl.Name)
	ppl.GetScore()
}
