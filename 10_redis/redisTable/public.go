package redisTable

import (
	"github.com/garyburd/redigo/redis"
	"../sql"
	"fmt"
	"errors"
	"strings"
)

var Rconn redis.Conn
var X sql.Xer

func Init_mysql_redis() {
	var err error

	// 连接mysql服务器
	X, err = sql.NewXorm("mysql", "root", "123456", "192.168.8.102", "3306", "taobao")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 连接redis服务器
	Rconn, err = redis.Dial("tcp", "192.168.8.102:6379")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// 关闭mysql和redis连接
func CloseConn() {
	X.Close()
	Rconn.Close()
}



// 获取int64整形值，测试：fmt.Println(Gint64(int64(123)))
func Gint64(rs interface{}) (int64, error) {
	if rs == nil {
		return 0, nil
	}
	switch rs.(type) {
	case int64:
		return rs.(int64), nil
	}
	return -1, errors.New("variable is not int64")
}

// 获取字符串值，测试：fmt.Println(Gstring([]uint8("hello")))
func Gstring(rs interface{}) (string, error) {
	if rs == nil {
		return "", nil
	}
	switch rs.(type) {
	case []uint8:
		return string(rs.([]uint8)), nil
	}
	return "", errors.New("variable is not []uint8")
}

// 获取多返回值的字符串形式，测试：fmt.Println(Gsstring([]interface{}{[]uint8("hello"),[]uint8("word")}))
func Gsstring(rs interface{}) ([]string, error) {
	if rs == nil {
		return nil, nil
	}
	switch rs.(type) {
	case []interface{}:
		var vals []string
		for _, v := range rs.([]interface{}) {
			vals = append(vals, string(v.([]uint8)))
		}
		return vals, nil
	}
	return nil, errors.New("variable is not []interface{}")
}

// 添加有序序列
func InsertSortSet(zName string, market []string) (int, error) {
	sum := 0
	for _, mk := range market {
		kv := strings.Split(mk, ":")
		if len(kv) == 2 {
			rs, err := Rconn.Do("zscore", zName, kv[1])
			if err != nil {
				return sum, err
			}
			val, _ := Gstring(rs)
			if val != "" {
				// 说明key已经存在
				continue
			} else {
				rs, err = Rconn.Do("zadd", zName, kv[0], kv[1])
				if err != nil {
					return sum, err
				}
				sum++
			}
		}
	}

	return sum, nil
}

// 删除有序序列元素
func DeleteSortSet(zName string, val string) (interface{}, error) {
	return Rconn.Do("zrem", zName, 1, val)
}