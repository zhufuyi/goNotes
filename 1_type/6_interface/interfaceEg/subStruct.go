// 嵌入接口使用
/*
当结构体A是结构体B的子集时
（1）结构体B没有定义新方法情况，结构体B实现了结构体A的所有方法
（2）结构体B有定义新方法情况，结构体B既可以实现新方法，也可以实现结构体A的方法（需指定）
*/

package interfaceEg

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

// 发送邮件
func (u *User) SendEmail() {
	fmt.Printf("User:  send email to %s<%s>\n", u.Name, u.Email)
}

// 接收邮件
func (u *User) ReceiveEmail(email string) {
	fmt.Printf("User:  receive email from %s\n", email)
}

// 定义一个User实现的接口
type SRemailer interface {
	SendEmail()
	ReceiveEmail(email string)
}

func srEmail(sr SRemailer) {
	sr.SendEmail()
	sr.ReceiveEmail("匿名邮件")
}

// 新的结构体
type Admin struct {
	*User // 结构体User是Admin的子集
	Lever string
}

// 重新定义方法，和User的方法名称一样，但不会冲突，使用时会覆盖接口子集的同名方法
func (a *Admin) SendEmail() {
	fmt.Printf("Admin:  send email to %s<%s>\n", a.Name, a.Email)
}

func SubStuctEg(info string) {
	fmt.Printf("\n\n------------------------ %s ------------------------\n", info)

	user := &User{"Zhangsan", "zhangsan@126.com"}
	srEmail(user)

	fmt.Println()

	admin := &Admin{&User{"Lisi", "Lisi@126.com"}, "root"}
	srEmail(admin)

	// 结果：
	// User:  send email to Zhangsan<zhangsan@126.com>
	// User:  receive email from 匿名邮件

	// Admin:  send email to Lisi<Lisi@126.com>
	// User:  receive email from 匿名邮件
}
