package interfaceEg

import (
	"fmt"
)

// 1.定义结构体
type User struct {
	Name  string
	Email string
}

//2. 定义User结构体类型的方法Notify，该方法接受者是User类型，要调用 Notify 方法我们需要一个 User 类型的值或者指针
func (u *User) Notify() error { // (指针)
	fmt.Printf("User: Sending User Email To %s<%s>\n", u.Name, u.Email)
	return nil
}

/*
func (u User) Notify() error { // （值，传递的是副本，修改副本不改变原值）
	fmt.Printf("User: Sending User Email To %s<%s>\n", u.Name, u.Email)
	return nil
}
*/

// 3. 定义一个叫做 Notifier 的接口并包含一个 Notify 方法
type Notifier interface {
	Notify() error
}

// 4. 定义一个函数来接受任意一个实现了接口 Notifier 的类型的值或者指针
func SendNotification(notify Notifier) error {
	return notify.Notify()
}

// 新的结构体
type Admin struct {
	User  // 结构体User是Admin的子集
	Lever string
}

func (a *Admin) Notify() error { // 和User的方法名称一样，但不会冲突，同时使用时执行的是外部方法，只有具体指定user时才执行内部方法
	fmt.Printf("Admin: Sending Admin Email To %s<%s>\n", a.Name, a.Email)
	return nil
}

func SubStuctEg() {
	// 5. 用 User 类型来实现该接口并且传入一个 User 类型来调用 SendNotification 方法
	user := &User{"Guanyu", "guanyu@126.com"} // user为指向结构体User指针，对应方法Notify接受也是结构体User指针
	user.Notify()
	SendNotification(user)

	fmt.Println()

	admin := &Admin{User: User{"Liubei", "liubei@126.com"}, Lever: "master"}
	SendNotification(admin) //（如果定义新方法，则使用新方法，否则使用User方法）
	admin.Notify()          // （如果定义新方法，则使用新方法，否则使用User方法）这些字段和方法也同样被提升到了外部类型
	admin.User.Notify()     // 嵌入类型的名字充当着字段名，同时嵌入类型作为内部类型存在

	// 例子说明当结构体A是结构体B的子集时
	// 1）结构体B没有定义新方法情况，结构体B实现了结构体A的所有方法
	// 2）结构体B有定义新方法情况，结构体B既可以实现新方法，也可以实现结构体A的方法（需指定），
}
