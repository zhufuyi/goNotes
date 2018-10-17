// redis的店铺表信息的添加和删除
package redisTable

import (
	"strconv"
	"fmt"
	"errors"
)

type Shoper interface {
	GetMaxId() (int, error)
	GetShopContent(min int, max int) ([][]string, error)
	AnalyzeValue(rs [][]string) error
	InsertShopRow() (int, error)
	DeleteShopRow(id int) error
	IdIsExist(id string) bool
}

// 实例化
func NewShop(columnName []string, rowKey, allIdKey, name2idKey, site2idKey, addr2idKey, credit2idKey, score2idKey string) Shoper {
	return &shopObj{
		columnName:columnName,
		rowKey:rowKey,
		allIdKey:allIdKey,
		name2idKey:name2idKey, // 后面补.id字符串，例如shop:name="apple".id
		site2idKey:site2idKey, // 后面补.id字符串
		addr2idKey:addr2idKey, // 后面补.id字符串
		credit2idKey:credit2idKey,
		score2idKey:score2idKey,
	}
	/*
		// 定义redis的key
		so := &shopObj{
			columnName:		[]string{"id", "name", "addr", "site", "credit", "score"},
			rowKey:			"shop:id=",
			allIdKey:		"shop:all_id",
			name2idKey:		"shop:name=", // 后面补.id字符串，例如shop:name="apple".id
			site2idKey:		"shop:site=", // 后面补.id字符串
			addr2idKey:		"shop:addr=", // 后面补.id字符串
			credit2idKey:	"shop:credit",
			score2idKey:	"shop:score",
		}
	*/

}


// shop表对象
type shopObj struct {
	columnName     []string        // 从mysql读取数据时输出列的名称和顺序

	rowKey         string          // 存储每一行shop表数据的key名
	rowValue       [][]interface{} // 待写入hash value

	allIdKey       string          // 存储店铺所有id的key名
	allIdValue     [][]interface{} // 待写入list value

	name2idKey     string          // 存储店铺名对应的店铺id的key名
	name2idValue   [][]interface{} // 待写入string value

	site2idKey     string          // 存储店铺网址名对应的店铺id的key名
	site2idValue   [][]interface{} // 待写入string value

	addr2idKey     string          // 存储相同店铺地址的店铺id的key名
	addr2idValue   [][]interface{} // 待写入list value

	credit2idKey   string          // 存储信用等级值和对应shopid的key名
	credit2idValue [][]interface{} // 待写入sort set value

	score2idKey    string          // 存储店铺评分值和对应shopid的key名
	score2idValue  [][]interface{} // 待写入sort set value
}

// 从mysql获取商品最大id值
func (so *shopObj)GetMaxId() (int, error) {
	sql := "SELECT max(id) FROM shop"
	rs, err := X.ExecString(sql)
	if err != nil {
		return -1, err
	}

	for _, num := range rs {
		for _, max := range num {
			return strconv.Atoi(max)
		}
	}
	return 0, nil
}

// 从mysql获取指定长度的shop表信息
func (so *shopObj)GetShopContent(min int, max int) ([][]string, error) {
	sql := "SELECT * FROM shop WHERE id>=" + fmt.Sprintf("%d", min) + " AND id<=" + fmt.Sprintf("%d", max)  //sql:="SELECT max(id) FROM shop"
	columnName := []string{"id", "name", "addr", "shop_site", "shop_credit", "shop_score"}
	rs, err := X.ExecString(sql, columnName...)
	if err != nil {
		return nil, err
	}
	return rs, err
}

// 解析为实际需要插入redis类型数据格式，输入多行的数据，如果是单行数据[]string，需要包装成多行数据[][]string
func (so *shopObj)AnalyzeValue(rs [][]string) error {
	var rowValue [][]interface{}
	var allIdValue [][]interface{}
	var name2idValue [][]interface{}
	var site2idValue [][]interface{}
	var addr2idValue [][]interface{}
	var credit2idValue [][]interface{}
	var score2idValue [][]interface{}

	for _, row := range rs {
		if len(so.columnName) != len(row) {
			// 插入列的数量必须等于mysql列的数量
			return errors.New("column lenght error")
		}

		var rowData []interface{}
		var allIdData []interface{}
		var name_Id string
		var name2idData []interface{}
		var site_Id string
		var site2idData []interface{}
		var addr_Id string
		var addr2idData []interface{}
		var credit_id string
		var credit2idData []interface{}
		var score_id string
		var score2idData []interface{}

		for k, cel := range row {

			// 解析为rowValue(哈希)
			if k == 0 {
				// 每一行数据的第一个cell值为id
				rowData = append(rowData, so.rowKey + cel, so.columnName[k], cel)    // 插入rowData的key
			} else {
				rowData = append(rowData, so.columnName[k], cel)
			}

			// 解析为allIdValue(链表)
			if k == 0 {
				// 每一行数据的第一个cell值为id
				allIdData = append(allIdData, so.allIdKey, cel)
			}

			// 解析为name2idValue(字符串)
			if k == 0 {
				name_Id = cel
			} else if k == 1 {
				name2idData = append(name2idData, so.name2idKey + cel + ".id", name_Id)
			}

			// 解析为site2idValue(字符串)
			if k == 0 {
				site_Id = cel
			} else if k == 3 {
				site2idData = append(site2idData, so.site2idKey + cel + ".id", site_Id)
			}

			// 解析为addr2idValue(链表)
			if k == 0 {
				addr_Id = cel
			} else if k == 2 {
				addr2idData = append(addr2idData, so.addr2idKey + cel + ".id", addr_Id)
			}

			// 解析为credit2idValue(有序集合)
			if k == 0 {
				credit_id = cel
			} else if k == 4 {
				credit2idData = append(credit2idData, so.credit2idKey, cel, credit_id)
			}

			// 解析为credit2idValue(有序集合)
			if k == 0 {
				score_id = cel
			} else if k == 5 {
				score2idData = append(score2idData, so.score2idKey, cel, score_id)
			}
		}
		rowValue = append(rowValue, rowData)
		allIdValue = append(allIdValue, allIdData)
		name2idValue = append(name2idValue, name2idData)
		site2idValue = append(site2idValue, site2idData)
		addr2idValue = append(addr2idValue, addr2idData)
		credit2idValue = append(credit2idValue, credit2idData)
		score2idValue = append(score2idValue, score2idData)
	}
	so.rowValue = rowValue
	so.allIdValue = allIdValue
	so.name2idValue = name2idValue
	so.site2idValue = site2idValue
	so.addr2idValue = addr2idValue
	so.credit2idValue = credit2idValue
	so.score2idValue = score2idValue
	return nil
}

// 添加redis数据，类似于mysql的shop表的一行数据，使用事务，要么一起失败，要么一起成功，但不支持回滚
// 返回添加成功的行数和错误
func (so *shopObj)InsertShopRow() (int, error) {
	var err error
	sum := 0
	for i, v := range so.allIdValue {
		// 判断id是否存在
		if len(v) == 2 {
			if so.IdIsExist(v[1].(string)) {
				continue
			}
		}

		Rconn.Do("muilt")

		_, err = Rconn.Do("hmset", so.rowValue[i]...)   // 插入哈希值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("lpush", so.allIdValue[i]...) // 插入链表值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("set", so.name2idValue[i]...) // 插入字符串值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("set", so.site2idValue[i]...) // 插入字符串值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}
		_, err = Rconn.Do("lpush", so.addr2idValue[i]...)   // 插入链表值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("zadd", so.credit2idValue[i]...)  // 插入有序集合值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("zadd", so.score2idValue[i]...)   // 插入有序集合值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}
		sum++
		Rconn.Do("exec")
	}
	// 添加成功后删除buf
	so.deleteBuf()
	return sum, nil
}


// 删除redis数据，类似于删除mysql的shop表的一行数据，使用事务
func (so *shopObj)DeleteShopRow(id int) error {
	var err error
	id_string := fmt.Sprintf("%d", id)

	Rconn.Do("muilt")

	// 删除前先获取和该id相关的name
	rs, err := Rconn.Do("hget", so.rowKey + id_string, "name") // 获取id的name
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	name, _ := Gstring(rs)
	// 删除前先获取和该id相关的addr
	rs, err = Rconn.Do("hget", so.rowKey + id_string, "addr") // 获取id的addr
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	addr, _ := Gstring(rs)
	// 删除前先获取和该id相关的site
	rs, err = Rconn.Do("hget", so.rowKey + id_string, "site") // 获取id的site
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	site, _ := Gstring(rs)

	_, err = Rconn.Do("del", so.rowKey + id_string)   // 删除哈希key
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("lrem", so.allIdKey, 1, id_string) // 删除链表key里的一个值
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("del", so.name2idKey + name + ".id", id_string) // 删除name的key
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("del", so.site2idKey + site + ".id", id_string) // 删除site的key
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	_, err = Rconn.Do("lrem", so.addr2idKey + addr + ".id", 1, id_string)   // 删除addr链表key里的值
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	_, err = Rconn.Do("zrem", so.credit2idKey, id_string)  // 在信誉等级有序集合中删除指定的店铺id
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	_, err = Rconn.Do("zrem", so.score2idKey, id_string)   // 在评分有序集合中删除指定的店铺id
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	Rconn.Do("exec")

	return nil
}

// 判断id是否已经存在
func (so *shopObj)IdIsExist(id string) bool {
	rs, _ := Rconn.Do("lrange", "shop:all_Id", 0, -1)
	ids, _ := Gsstring(rs)
	for _, id_str := range ids {
		if id_str == id {
			return true
		}
	}
	return false
}

// 删除数据缓冲
func (so *shopObj)deleteBuf() {
	so.rowValue = nil
	so.allIdValue = nil
	so.name2idValue = nil
	so.site2idValue = nil
	so.addr2idValue = nil
	so.credit2idValue = nil
	so.score2idValue = nil
}


