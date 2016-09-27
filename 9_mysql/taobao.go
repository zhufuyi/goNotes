/*读取exel文件数据内容，把内容插入mysql数据库，并查询测试，注：exel数据文件来自网络*/
package main

import (
	"fmt"
	"./sql"
	"strings"
	"./filename"
	"./exel"
)

var X sql.Xer


// 字符串拼接并执行sql语句
func insertData(sql string, d []string) error {
	for _, v := range d {
		sql += "('" + v + "'),"    // 字符串拼接
	}
	sql = strings.TrimRight(sql, ",")    // 去掉最后的逗号
	_, err := X.Exec(sql)
	return err
}



// 字符串拼接并执行sql语句2
func insertData2(sql string, d map[string]string) error {
	for _, v := range d {
		sql += "(" + v + "),"    // 字符串拼接
	}
	sql = strings.TrimRight(sql, ",")    // 去掉最后的逗号
	_, err := X.Exec(sql)
	return err
}

// 把exel文件内容插入数据库
func insert2database() {
	// 插入市场分类到数据库
	func() {
		sql := "INSERT INTO marketgroup (name) VALUES "
		market := []string{"服饰", "花鸟", "家电", "家居", "美妆", "母婴", "男包", "男鞋", "男装", "内衣", "女包", "女鞋", "女装", "配饰", "食品", "数码", "软件", "未分类"}
		err := insertData(sql, market)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok")
		}
		//fmt.Println(X.ExecString("SELECT * FROM marketgroup", "id", "name"))
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


	// 插入卖家信息到数据库
	func() {
		market := []string{"服饰", "花鸟", "家电", "家居", "美妆", "母婴", "男包", "男鞋", "男装", "内衣", "女包", "女鞋", "女装", "配饰", "食品", "数码", "软件", "未分类"}

		// 读取文件列表
		f := filename.Getfiles("./淘宝卖家数据库")
		files, err := f.GetExelFiles()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Println(exels)

		sellers := make(map[string]string)
		for _, fileName := range files {
			seller, gds, err := exel.OpenExel(fileName).Extract()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			s2i := make(map[string]int)
			for _, sl := range seller {
				if _, ok := sellers[sl.SellerName]; !ok {
					// 不存在则写入
					sid++
					s2i[sl.SellerName] = sid
					sellers[sl.SellerName] = fmt.Sprintf(" '%d','%s','%s','%s','%d','%f'", sid, sl.SellerName, sl.Addr, sl.ShopSite, sl.Credit, sl.Score)
				}
			}

			for mid, mk := range market {
				// 通过文件名来查找市场分类的序号
				if strings.Contains(fileName, mk) {
					for _, gd := range gds {
						gd.MarketId = mid + 1    // 补充该商品属于哪个市场分类，id从1开始
						gd.SellerId = s2i[gd.SellerNam]
					}
					break
				}
			}

			goods = append(goods, gds)
		}

		sql := "INSERT INTO seller (id,name,addr,shop_site,credit_val,shop_score) VALUES "
		err = insertData2(sql, sellers)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok!")
		}
		//fmt.Println(X.ExecString("SELECT count(*) FROM seller", "count(*)"))
	}()


	// 插入商品信息到数据库
	func() {
		goodsmap := make(map[string]string)
		for _, gds := range goods {
			for _, gd := range gds {
				goodsmap[gd.GoodsSite] = fmt.Sprintf(" '%d','%d','%s','%s','%f'", gd.SellerId, gd.MarketId, gd.GoodsName, gd.GoodsSite, gd.Praise)
			}
		}

		sql := "INSERT INTO goods (seller_id,market_id,goods_name,goods_site,goods_praise) VALUES "
		err := insertData2(sql, goodsmap)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert ok!")
		}
		//fmt.Println(X.ExecString("SELECT count(*) FROM seller", "count(*)"))

	}()
}


// 测试查询
func testQuery() {
	var sql string
	var rs [][]string

	// (1) 查看一共有多少个卖家
	fmt.Println("\n\n-------------查看一共有多少个卖家-----------------")
	sql = "SELECT count(*) FROM seller"
	rs, _ = X.ExecString(sql)
	fmt.Println("卖家数量 =", rs[0][0])

	// (2) 查看一共有多少个商品
	fmt.Println("\n\n-------------查看一共有多少个商品-----------------")
	sql = "SELECT count(*) FROM goods"
	rs, _ = X.ExecString(sql)
	fmt.Println("商品数量 =", rs[0][0])

	// (3) 查看有多少个天猫卖家
	fmt.Println("\n\n--------------查看有多少个天猫卖家----------------")
	sql = "SELECT count(*) FROM seller WHERE credit_val = (SELECT id FROM credibility WHERE name='天猫')"
	rs, _ = X.ExecString(sql)
	fmt.Println("天猫卖家数量 =", rs[0][0])

	// (4) 查找地址在广州并且店铺评分大于4.95分的皇冠以上卖家
	fmt.Println("\n\n--------------查找地址在广州并且店铺评分大于4.95分的皇冠以上卖家----------------")
	sql = "SELECT * FROM seller WHERE (addr REGEXP '广州') AND (shop_score>4.95) AND (credit_val>(SELECT id FROM credibility WHERE name='一黄冠'))"
	rs, _ = X.ExecString(sql, "name", "addr", "credit_val", "shop_score")
	for _, v := range rs {
		fmt.Println(v)
	}

	// (5) 查询好评率大于98%的手机相关商品
	fmt.Println("\n\n--------------查询好评率大于98%的手机商品----------------")
	sql = "SELECT * FROM goods WHERE (goods_name REGEXP '手机') AND (goods_praise>98) AND (market_id=(SELECT id FROM marketgroup WHERE name='数码'))"
	rs, _ = X.ExecString(sql, "goods_name", "goods_praise")
	for _, v := range rs {
		fmt.Println(v)
	}

	// (6) 查询至少有6个商品以上的卖家信息
	fmt.Println("\n\n--------------查询至少有6个商品以上的卖家信息----------------")
	sql = "SELECT * FROM seller RIGHT JOIN (SELECT seller_id,count(*) AS cnt FROM goods GROUP BY seller_id HAVING cnt >=6) AS multigoods ON seller.id=multigoods.seller_id HAVING multigoods.seller_id>0"
	rs, _ = X.ExecString(sql, "id", "name", "addr", "credit_val", "shop_score", "cnt")
	for _, v := range rs {
		fmt.Println(v)
	}

	// (7)查询各市场分类中好评度最好的商品的卖家名称
	fmt.Println("\n\n--------------查询各市场分类中好评度最好的商品的卖家名称----------------")
	sql = "SELECT name FROM seller RIGHT JOIN (SELECT seller_id,max(goods_praise) FROM goods GROUP BY market_id) AS bestgoods ON seller.id=bestgoods.seller_id"
	rs, _ = X.ExecString(sql, "id", "name", "addr", "credit_val", "shop_score", "cnt")
	for _, v := range rs {
		fmt.Println(v)
	}

}

func main() {
	// 连接数据库
	var err error
						//驱动名   用户名	密码		主机	  端口	 数据库名	
	X, err = sql.NewXorm("mysql", "root", "123456", "localhost", "3306", "taobao")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 插入数据
	insert2database()

	// 测试查询
	testQuery()
}