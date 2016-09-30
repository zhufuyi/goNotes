package main

import (
	"./exel"
	"./filename"
	"./sql"
	"fmt"
	"strings"
	"time"
)

var X sql.Xer

// 创建表
func createTables() {
	// 创建市场分类表
	sql := `CREATE TABLE market(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name CHAR(20) NOT NULL DEFAULT '' UNIQUE
	)ENGINE myisam CHARSET utf8`
	fmt.Println(X.Exec(sql))

	// 创建信用等级表
	sql = `CREATE TABLE credibility(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name CHAR(20) NOT NULL DEFAULT '' UNIQUE
	)ENGINE myisam CHARSET utf8`
	fmt.Println(X.Exec(sql))

	// 创建店铺信息表
	sql = `CREATE TABLE shop(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL DEFAULT '' UNIQUE ,
    addr VARCHAR(50) NOT NULL DEFAULT '',
    shop_site VARCHAR(100) NOT NULL DEFAULT '' UNIQUE,
    shop_credit TINYINT UNSIGNED NOT NULL DEFAULT 0,
    shop_score FLOAT NOT NULL DEFAULT 0.0,
    KEY addr(addr),
    KEY shop_credit(shop_credit),
    KEY shop_score(shop_score)
	)ENGINE myisam CHARSET utf8`
	fmt.Println(X.Exec(sql))

	// 创建商品信息表
	sql = `CREATE TABLE goods(
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    shop_id INT UNSIGNED NOT NULL DEFAULT 0,
    market_id INT UNSIGNED NOT NULL DEFAULT 0,
    goods_name CHAR(60) NOT NULL DEFAULT '',
    goods_site VARCHAR(100) NOT NULL DEFAULT '' UNIQUE,
    goods_praise FLOAT NOT NULL DEFAULT 0.0,
    KEY shop_id(shop_id),
    KEY market_id(market_id),
    KEY goods_praise(goods_praise),
    KEY goods_name(goods_name(10))
	)ENGINE myisam CHARSET utf8`
	fmt.Println(X.Exec(sql))
}

// 字符串拼接并执行sql语句
func insertData(sql string, d []string) error {
	for _, v := range d {
		sql += "('" + v + "')," // 字符串拼接
	}
	sql = strings.TrimRight(sql, ",") // 去掉最后的逗号
	_, err := X.Exec(sql)
	return err
}

// 字符串拼接并执行sql语句2
func insertData2(sql string, d map[string]string) error {
	cnt := 1
	str := ""
	for _, v := range d {
		cnt++
		str += "(" + v + ")," // 字符串拼接
		if cnt % 10 == 0 {
			// 每次插入10行
			str = strings.TrimRight(str, ",") // 去掉最后的逗号
			_, err := X.Exec(sql + str)
			if err != nil {
				return err
			}
			str = ""
		}
	}

	return nil
}

// 把exel文件内容插入数据库
func insert2database() {
	// 插入市场分类到数据库
	func() {
		sql := "INSERT INTO market (name) VALUES "
		market := []string{"服饰", "花鸟", "家电", "家居", "美妆", "母婴", "男包", "男鞋", "男装", "内衣", "女包", "女鞋", "女装", "配饰", "食品", "数码", "软件", "未分类"}
		err := insertData(sql, market)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok")
		}
		//fmt.Println(X.ExecString("SELECT * FROM market", "id", "name"))
	}()

	// 插入信誉等级分类到数据库
	func() {
		sql := "INSERT INTO credibility (name) VALUES "
		credit := []string{"一星", "二星", "三星", "四星", "五星", "一钻", "二钻", "三钻", "四钻", "五钻", "一黄冠", "二黄冠", "三黄冠", "四黄冠", "五黄冠", "一金冠", "二金冠", "三金冠", "四金冠", "五金冠", "天猫"}
		err := insertData(sql, credit)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok")
		}
		//fmt.Println(X.ExecString("SELECT * FROM credibility", "id", "name"))
	}()

	var goods [][]*exel.Goods
	sid := 0

	// 插入店铺信息到数据库
	func() {
		market := []string{"服饰", "花鸟", "家电", "家居", "美妆", "母婴", "男包", "男鞋", "男装", "内衣", "女包", "女鞋", "女装", "配饰", "食品", "数码", "软件", "未分类"}

		// 读取文件列表
		f := filename.Getfiles("./淘宝数据")
		files, err := f.GetExelFiles()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		shops := make(map[string]string)
		for _, fileName := range files {
			shps, gds, err := exel.OpenExel(fileName).Extract()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			s2i := make(map[string]int) // 关联shop名称和shopid
			for _, sl := range shps {
				if _, ok := shops[sl.ShopName]; !ok {
					// 不存在则写入
					sid++
					s2i[sl.ShopName] = sid
					shops[sl.ShopName] = fmt.Sprintf(" '%d','%s','%s','%s','%d','%f'", sid, sl.ShopName, sl.Addr, sl.ShopSite, sl.Credit, sl.Score)
				}
			}

			for mid, mk := range market {
				// 通过文件名来查找市场分类的序号
				if strings.Contains(fileName, mk) {
					for _, gd := range gds {
						gd.MarketId = mid + 1 // 补充该商品属于哪个市场分类，id从1开始
						gd.ShopId = s2i[gd.ShopName]
					}
					break
				}
			}

			goods = append(goods, gds)
		}

		sql := "INSERT INTO shop (id,name,addr,shop_site,shop_credit,shop_score) VALUES "
		err = insertData2(sql, shops)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok!")
		}
		//fmt.Println(X.ExecString("SELECT count(*) FROM shop", "count(*)"))
	}()

	// 插入商品信息到数据库
	func() {
		goodsmap := make(map[string]string)
		for _, gds := range goods {
			for _, gd := range gds {
				goodsmap[gd.GoodsSite] = fmt.Sprintf(" '%d','%d','%s','%s','%f'", gd.ShopId, gd.MarketId, gd.GoodsName, gd.GoodsSite, gd.Praise)
			}
		}

		sql := "INSERT INTO goods (shop_id,market_id,goods_name,goods_site,goods_praise) VALUES "
		err := insertData2(sql, goodsmap)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok!")
		}
		//fmt.Println(X.ExecString("SELECT count(*) FROM shop", "count(*)"))

	}()
}

// 测试查询
func testQuery() {
	var sql string
	var rs [][]string

	// (1) 查看一共有多少个店铺
	fmt.Println("\n\n-------------查看一共有多少个店铺-----------------")
	t := time.Now()
	func() {
		sql = "SELECT count(*) FROM shop"
		rs, _ = X.ExecString(sql)
		fmt.Println("店铺数量 =", rs[0][0])
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (2) 查看一共有多少个商品
	fmt.Println("\n\n-------------查看一共有多少个商品-----------------")
	t = time.Now()
	func() {
		sql = "SELECT count(*) FROM goods"
		rs, _ = X.ExecString(sql)
		fmt.Println("商品数量 =", rs[0][0])
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (3) 查看有多少个天猫店铺
	fmt.Println("\n\n--------------查看有多少个天猫店铺----------------")
	t = time.Now()
	func() {
		sql = "SELECT count(*) FROM shop WHERE shop_credit = (SELECT id FROM credibility WHERE name='天猫')"
		rs, _ = X.ExecString(sql)
		fmt.Println("天猫店铺数量 =", rs[0][0])
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (4) 查找地址在广州并且店铺评分至少4.95分的皇冠以上店铺名称
	fmt.Println("\n\n--------------查找地址在广州并且店铺评分至少4.95分的皇冠以上店铺名称----------------")
	t = time.Now()
	func() {
		sql = "SELECT * FROM shop WHERE (addr REGEXP '广州') AND (shop_score>=4.949) AND (shop_credit>=(SELECT id FROM credibility WHERE name='一黄冠'))"
		rs, _ = X.ExecString(sql, "name")
		for _, v := range rs {
			fmt.Println(v[0])
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (5) 查询好评率99%以上的手机相关商品的网址
	fmt.Println("\n\n--------------查看好评率99%以上的手机相关商品的网址----------------")
	t = time.Now()
	func() {
		sql = "SELECT * FROM goods WHERE (goods_name REGEXP '手机') AND (goods_praise>=99)"
		rs, _ = X.ExecString(sql, "goods_site")
		for _, v := range rs {
			fmt.Println(v[0])
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (6) 查询至少有6个商品以上的店铺名称
	fmt.Println("\n\n--------------查询至少有6个商品以上的店铺名称----------------")
	t = time.Now()
	func() {
		sql = "SELECT * FROM shop RIGHT JOIN (SELECT shop_id,count(*) AS cnt FROM goods GROUP BY shop_id HAVING cnt >=6) AS multigoods ON shop.id=multigoods.shop_id HAVING multigoods.shop_id>0"
		rs, _ = X.ExecString(sql, "name")
		for _, v := range rs {
			fmt.Println(v[0])
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (7)查询各市场分类中好评度最好最新的商品的店铺名称
	fmt.Println("\n\n--------------查询各市场分类中好评度最好最新的商品的店铺名称----------------")
	t = time.Now()
	func() {
		sql = `SELECT market.name,bestshop.name FROM market RIGHT JOIN
    (SELECT id,name,market_id FROM shop RIGHT JOIN
        (SELECT shop_id,market_id FROM goods WHERE id IN (SELECT max(id) FROM goods WHERE goods_praise>=100 GROUP BY market_id)
        ) AS bestgoods ON shop.id=bestgoods.shop_id) AS bestshop
    ON market.id=bestshop.market_id`
		rs, _ = X.ExecString(sql, "name")
		for _, v := range rs {
			fmt.Println(v[0])
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())
}

func main() {
	// (1)连接数据库
	var err error
	X, err = sql.NewXorm("mysql", "root", "123456", "192.168.8.102", "3306", "taobao")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// (2)创建表
	//createTables()

	// (3)插入数据
	//insert2database()

	// (4)测试查询
	testQuery()
}
