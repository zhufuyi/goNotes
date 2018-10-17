package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
)

// 执行原生sql接口
type Xer interface {
	Exec(sql string) ([]map[string][]byte, error)
	ExecString(sql string, Columns ...string) ([][]string, error)
}

// orm对象
type xORM struct {
	orm *xorm.Engine
}


// 实例化
func NewXorm(driverName, user, password, host, port, dbname string) (Xer, error) {
	x := &xORM{}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, dbname)
	err := x.connectDB(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return x, nil
}

// 连接数据库
func (x *xORM)connectDB(driverName string, dataSourceName string) error {
	var err error
	x.orm, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = x.orm.Ping()// 测试是否连接数据库
	if err != nil {
		return err
	}
	return nil
}

// 执行原生sql
func (x *xORM)Exec(sql string) ([]map[string][]byte, error) {
	return x.orm.Query(sql)
}

// 查询结果字符串形式,第二个参数为选择打印的列，如果不指定列输出全部列（随机顺序）
func (x *xORM)ExecString(sql string, Columns ...string) ([][]string, error) {
	result, err := x.orm.Query(sql)
	if err != nil {
		return nil, err
	}

	var vals [][]string
	if Columns == nil {
		// 如果不指定打印全部列值
		for _, res := range result {
			var tmp []string
			for _, vl := range res {
				tmp = append(tmp, string(vl))
			}
			vals = append(vals, tmp)
		}
		return vals, nil
	}

	for _, res := range result {
		var tmp []string
		for _, vl := range Columns {
			// 判断列名是否存在
			if v, ok := res[vl]; ok {
				tmp = append(tmp, string(v))
			}
		}
		vals = append(vals, tmp)
	}
	return vals, nil
}

