// redis的商品表信息的添加和删除
package redisTable

import (
	"strconv"
	"fmt"
	"errors"
)

type Goodser interface {
	GetMaxId() (int, error)
	GetGoodsContent(min int, max int) ([][]string, error)
	AnalyzeValue(rs [][]string) error
	InsertGoodsRow() (int, error)
	DeleteGoodsRow(id int) error
	IdIsExist(id string) bool
}

// 实例化
func NewGoodser(columnName []string, rowKey, allIdKey, shopId2idKey, marketId2idKey, name2idKey, site2idKey, praise2idKey string) Goodser {
	return &goodsObj{
		columnName:columnName,
		rowKey:rowKey,
		allIdKey:allIdKey,
		shopId2idKey:shopId2idKey, // 后面补.id字符串，例如shop:name="apple".id
		marketId2idKey:marketId2idKey, // 后面补.id字符串
		name2idKey:name2idKey, // 后面补.id字符串
		site2idKey:site2idKey, // 后面补.id字符串
		praise2idKey:praise2idKey,
	}
	/*
		// 定义redis的key
		so := &ShopObj{
			columnName:		[]string{"id", "shop_id", "mark_id", "name", "site", "praise"},
			rowKey:			"goods:id=",
			allIdKey:		"goods:all_id",
			shopId2idKey:	"goods:shop_id=", 	// 后面补.id字符串，例如goods:shop_id="1111".id
			marketId2idKey:	"goods:market_id=", // 后面补.id字符串
			name2idKey:		"goods:name=", 		// 后面补.id字符串
			site2idKey:		"goods:site=", 		// 后面补.id字符串
			praise2idKey:	"goods:praise",
		}
	*/
}


// shop表对象
type goodsObj struct {
	columnName       []string        // 从mysql读取数据时输出列的名称和顺序

	rowKey           string          // 存储每一行goods表数据的key名
	rowValue         [][]interface{} // 待写入hash value

	allIdKey         string          // 存储goods所有id的key名
	allIdValue       [][]interface{} // 待写入list value

	shopId2idKey     string          // 存储商品id对应商品id的key名
	shopId2idValue   [][]interface{} // 待写入list value

	marketId2idKey   string          // 存储市场分类id对应商品id的key名
	marketId2idValue [][]interface{} // 待写入list value

	name2idKey       string          // 存储商品名字相同id的key名
	name2idValue     [][]interface{} // 待写入list value

	site2idKey       string          // 存储商品网址名对应id的key名
	site2idValue     [][]interface{} // 待写入string value

	praise2idKey     string          // 存储商品好评度分数对应id的key名
	praise2idValue   [][]interface{} // 待写入sort set value
}

// 从mysql获取商品最大id值
func (gdo *goodsObj)GetMaxId() (int, error) {
	sql := "SELECT max(id) FROM goods"
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
func (gdo *goodsObj)GetGoodsContent(min int, max int) ([][]string, error) {
	sql := "SELECT * FROM goods WHERE id>=" + fmt.Sprintf("%d", min) + " AND id<=" + fmt.Sprintf("%d", max)  //sql:="SELECT max(id) FROM shop"
	columnName := []string{"id", "shop_id", "market_id", "goods_name", "goods_site", "goods_praise"}
	rs, err := X.ExecString(sql, columnName...)
	if err != nil {
		return nil, err
	}
	return rs, err
}

// 解析为实际需要插入redis类型数据格式，输入多行的数据，如果是单行数据[]string，需要包装成多行数据[][]string
func (gdo *goodsObj)AnalyzeValue(rs [][]string) error {
	var rowValue [][]interface{}
	var allIdValue [][]interface{}
	var shopId2idValue [][]interface{}
	var marketId2idValue [][]interface{}
	var name2idValue [][]interface{}
	var site2idValue [][]interface{}
	var praise2idValue [][]interface{}

	for _, row := range rs {
		if len(gdo.columnName) != len(row) {
			// 插入列的数量必须等于mysql列的数量
			return errors.New("column lenght error")
		}

		var rowData []interface{}
		var allIdData []interface{}
		var shopId_id string
		var shopId2idData []interface{}
		var marketId_id string
		var marketId2idData []interface{}
		var name_id string
		var name2idData []interface{}
		var site_id string
		var site2idData []interface{}
		var praise_id string
		var praise2idData []interface{}

		for k, cel := range row {

			// 解析为rowValue(哈希)
			if k == 0 {
				// 每一行数据的第一个cell值为id
				rowData = append(rowData, gdo.rowKey + cel, gdo.columnName[k], cel)    // 插入rowData的key
			} else {
				rowData = append(rowData, gdo.columnName[k], cel)
			}

			// 解析为allIdValue(链表)
			if k == 0 {
				// 每一行数据的第一个cell值为id
				allIdData = append(allIdData, gdo.allIdKey, cel)
			}

			// 解析为shopId2idData(链表)
			if k == 0 {
				shopId_id = cel
			} else if k == 1 {
				shopId2idData = append(shopId2idData, gdo.shopId2idKey + cel + ".id", shopId_id)
			}

			// 解析为marketId2idValue(字符串)
			if k == 0 {
				marketId_id = cel
			} else if k == 2 {
				marketId2idData = append(marketId2idData, gdo.marketId2idKey + cel + ".id", marketId_id)
			}

			// 解析为name2idValue(链表)
			if k == 0 {
				name_id = cel
			} else if k == 3 {
				name2idData = append(name2idData, gdo.name2idKey + cel + ".id", name_id)
			}

			// 解析为site2idValue(字符串)
			if k == 0 {
				site_id = cel
			} else if k == 4 {
				site2idData = append(site2idData, gdo.site2idKey + cel + ".id", site_id)
			}

			// 解析为praise2idValue(有序集合)
			if k == 0 {
				praise_id = cel
			} else if k == 5 {
				praise2idData = append(praise2idData, gdo.praise2idKey, cel, praise_id)
			}
		}
		rowValue = append(rowValue, rowData)
		allIdValue = append(allIdValue, allIdData)
		shopId2idValue = append(shopId2idValue, shopId2idData)
		marketId2idValue = append(marketId2idValue, marketId2idData)
		name2idValue = append(name2idValue, name2idData)
		site2idValue = append(site2idValue, site2idData)
		praise2idValue = append(praise2idValue, praise2idData)
	}
	gdo.rowValue = rowValue
	gdo.allIdValue = allIdValue
	gdo.shopId2idValue = shopId2idValue
	gdo.marketId2idValue = marketId2idValue
	gdo.name2idValue = name2idValue
	gdo.site2idValue = site2idValue
	gdo.praise2idValue = praise2idValue
	return nil
}

// 添加redis数据，类似于mysql的shop表的一行数据，使用事务，要么一起失败，要么一起成功，但不支持回滚
// 返回添加成功的行数和错误
func (gdo *goodsObj)InsertGoodsRow() (int, error) {
	var err error
	sum := 0
	for i, v := range gdo.allIdValue {
		// 判断id是否存在,存在则直接则跳过
		if len(v) == 2 {
			if gdo.IdIsExist(v[1].(string)) {
				continue
			}
		}

		Rconn.Do("muilt")

		_, err = Rconn.Do("hmset", gdo.rowValue[i]...)   // 插入哈希值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("lpush", gdo.allIdValue[i]...) // 插入链表值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("lpush", gdo.shopId2idValue[i]...) // 插入链表
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("lpush", gdo.marketId2idValue[i]...) // 插入字链表
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}
		_, err = Rconn.Do("lpush", gdo.name2idValue[i]...)   // 插入链表值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("set", gdo.site2idValue[i]...)  // 插入字符串
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}

		_, err = Rconn.Do("zadd", gdo.praise2idValue[i]...)   // 插入有序集合值
		if err != nil {
			Rconn.Do("discard")
			return sum, err
		}
		sum++
		Rconn.Do("exec")
	}
	// 添加成功后删除buf
	gdo.deleteBuf()
	return sum, nil
}


// 删除redis数据，类似于删除mysql的shop表的一行数据，使用事务
func (gdo *goodsObj)DeleteGoodsRow(id int) error {
	var err error
	id_string := fmt.Sprintf("%d", id)

	Rconn.Do("muilt")

	// 删除前先获取和该id相关的shop_id
	rs, err := Rconn.Do("hget", gdo.rowKey + id_string, "shop_id") // 获取id的name
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	shop_id, _ := Gstring(rs)
	// 删除前先获取和该id相关的market_id
	rs, err = Rconn.Do("hget", gdo.rowKey + id_string, "market_id") // 获取id的name
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	market_id, _ := Gstring(rs)
	// 删除前先获取和该id相关的name
	rs, err = Rconn.Do("hget", gdo.rowKey + id_string, "name") // 获取id的addr
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	name, _ := Gstring(rs)
	// 删除前先获取和该id相关的site
	rs, err = Rconn.Do("hget", gdo.rowKey + id_string, "site") // 获取id的site
	if err != nil {
		Rconn.Do("discard")
		return err
	}
	site, _ := Gstring(rs)

	_, err = Rconn.Do("del", gdo.rowKey + id_string)   // 删除哈希key,即一行数据
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("lrem", gdo.allIdKey, 1, id_string) // 在所有商品id链表中删除指定商品id
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("lrem", gdo.shopId2idKey + shop_id + ".id", 1, id_string) // 在店铺id链表中删除指定商品id
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("lrem", gdo.marketId2idKey + market_id + ".id", 1, id_string) // 在市场分类id链表中删除指定商品id
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("lrem", gdo.name2idKey + name + ".id", 1, id_string) // 在商品名称链表中删除指定商品id
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("del", gdo.site2idKey + site + ".id", id_string) // 删除site的key
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	_, err = Rconn.Do("zrem", gdo.praise2idKey, id_string)   // 在好评度有序集合中删除指定的商品id
	if err != nil {
		Rconn.Do("discard")
		return err
	}

	Rconn.Do("exec")

	return nil
}

// 判断id是否已经存在
func (gdo *goodsObj)IdIsExist(id string) bool {
	rs, _ := Rconn.Do("lrange", gdo.allIdKey, 0, -1)
	ids, _ := Gsstring(rs)
	for _, id_str := range ids {
		if id_str == id {
			return true
		}
	}
	return false
}

// 删除数据缓冲
func (gdo *goodsObj)deleteBuf() {
	gdo.rowValue = nil
	gdo.allIdValue = nil
	gdo.shopId2idValue = nil
	gdo.marketId2idValue = nil
	gdo.name2idValue = nil
	gdo.site2idValue = nil
	gdo.praise2idValue = nil
}



