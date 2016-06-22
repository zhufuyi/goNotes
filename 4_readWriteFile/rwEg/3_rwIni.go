package rwEg

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func RwIniFile() {
	// 1 初始化
	fc, err := goconfig.LoadConfigFile("./conf.ini")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 2 读配置信息
	fmt.Printf("\n获取所有段名： %q\n\n", fc.GetSectionList())

	fmt.Printf("\n获取某个段的所有Key名： %q\n\n", fc.GetKeyList("server"))

	sec, _ := fc.GetSection("server")
	fmt.Printf("\n获取某个段里所有Key内容： %q\n\n", sec)

	val, _ := fc.GetValue("server", "ip")
	fmt.Printf("\n获取某个段里某个Key的值： %q\n\n", val)

	fmt.Printf("\n获取某个段的注释： %q\n\n", fc.GetSectionComments("server"))

	fmt.Printf("\n获取某个key的注释： %q\n\n", fc.GetKeyComments("server", "ip"))

	// 3 修改或增加配置信息（当设置的Key不存在是会自动新建）
	fc.SetValue("server", "ip", "192.168.1.5") //只在内存里修改
	v1, _ := fc.GetValue("server", "ip")
	fmt.Printf("\n修改后key(ip)的内容：%#v\n\n", v1)

	fc.SetSectionComments("server", "change server info")
	fmt.Printf("\n修改后段(server)的注释： %q\n\n", fc.GetSectionComments("server"))

	fc.SetKeyComments("server", "ip", "change ip addr")
	fmt.Printf("\n修改后key(ip)的注释： %q\n\n", fc.GetKeyComments("server", "ip"))

	fc.SetValue("client", "ip", "202.202.171.99") // 当段或Key不存在是会自动创建
	v2, _ := fc.GetValue("client", "ip")
	fmt.Printf("\n新增key(ip)的内容：%#v\n\n", v2)

	// 4 删除一个配置信息(注：只修改内存数据没有更改实际文件)
	fc.DeleteKey("server", "fp")
	fmt.Printf("\n删除后key(fp)： %q\n\n", fc.GetKeyComments("server", "fp"))

	fc.DeleteSection("Demo")
	fmt.Printf("\n删除段(Demo)： %q\n\n", fc.GetSectionComments("Demo"))

	// 5 保存修改后的配置文件
	goconfig.SaveConfigFile(fc, "./conf.ini")

	// 6 获取所有配置信息(只读取内存数据)

}
