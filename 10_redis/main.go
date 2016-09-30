package main

import (
	"fmt"
	"./redisTable"
	"strconv"
	"sort"
	"time"
	"strings"
)

// 添加数据到redis
func addRowData() {
	// 添加市场分类信息到redis
	market := []string{"1:服饰", "2:花鸟", "3:家电", "4:家居", "5:美妆", "6:母婴", "7:男包", "8:男鞋", "9:男装", "10:内衣", "11:女包", "12:女鞋", "13:女装", "14:配饰", "15:食品", "16:数码", "17:软件", "18:未分类"}
	fmt.Println(redisTable.InsertSortSet("market", market))

	// 添加信用等级分类信息到redis
	credit := []string{"1:一星", "2:二星", "3:三星", "4:四星", "5:五星", "6:一钻", "7:二钻", "8:三钻", "9:四钻", "10:五钻", "11:一黄冠", "12:二黄冠", "13:三黄冠", "14:四黄冠", "15:五黄冠", "16:一金冠", "17:二金冠", "18:三金冠", "19:四金冠", "20:五金冠", "21:天猫"}
	fmt.Println(redisTable.InsertSortSet("credit", credit))

	// 添加商店信息到redis
	so := redisTable.NewShop([]string{"id", "name", "addr", "site", "credit", "score"}, "shop:id=", "shop:all_id", "shop:name=", "shop:site=", "shop:addr=", "shop:credit", "shop:score")

	maxRow, err := so.GetMaxId()    // 获取mysql最大id值
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if maxRow > 0 {
			// 大于0说明mysql表有数据
			rs, _ := so.GetShopContent(1, maxRow)    // 获取mysql的id范围[min,max]数据
			so.AnalyzeValue(rs)    // 解析mysql行的数据为redis数据
		}
	}

	fmt.Println(so.InsertShopRow())  // 利用事务同时插入数据解析后的shop表的行数据
	//so.DeleteShopRow(1)	// 删除一行数据

	// 添加商品信息到redis
	gdo := redisTable.NewGoodser([]string{"id", "shop_id", "market_id", "name", "site", "praise"}, "goods:id=", "goods:all_id", "goods:shop_id=", "goods:market_id=", "goods:name=", "goods:site=", "goods:praise")
	maxRow, err = gdo.GetMaxId()    // 获取mysql最大id值
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if maxRow > 0 {
			// 大于0说明mysql表有数据
			rs, _ := gdo.GetGoodsContent(1, maxRow)    // 获取mysql的id范围[min,max]数据
			gdo.AnalyzeValue(rs)    // 解析mysql行的数据为redis数据
		}
	}
	fmt.Println(gdo.InsertGoodsRow())  // 利用事务同时插入数据解析后的shop表的行数据
	//gdo.DeleteGoodsRow(1)	// 删除指定行数据	
}


// 测试redis查询
func testQuery() {

	// (1) 查看一共有多少个店铺
	fmt.Println("\n\n-------------查看一共有多少个店铺-----------------")
	t := time.Now()
	func() {
		rs, _ := redisTable.Rconn.Do("llen", "shop:all_id")
		sum, _ := redisTable.Gint64(rs)
		fmt.Println("店铺数量 =", sum)
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (2) 查看一共有多少个商品
	fmt.Println("\n\n-------------查看一共有多少个商品-----------------")
	t = time.Now()
	func() {
		rs, _ := redisTable.Rconn.Do("llen", "goods:all_id")
		sum, _ := redisTable.Gint64(rs)
		fmt.Println("商品数量 =", sum)
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (3) 查看有多少个天猫店铺
	fmt.Println("\n\n--------------查看有多少个天猫店铺----------------")
	t = time.Now()
	func() {
		// 获取天猫所对应的信用等级值
		rs, _ := redisTable.Rconn.Do("zscore", "credit", "天猫")
		tm, _ := redisTable.Gstring(rs)
		// 通过值获取天猫店铺在有序集合的个数
		rs, _ = redisTable.Rconn.Do("zrangebyscore", "shop:credit", tm, tm)
		sum, _ := redisTable.Gsstring(rs)
		fmt.Println("天猫店铺数量 =", len(sum))
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())


	// (4) 查找地址在广州并且店铺评分大于4.95分的皇冠以上店铺名称
	fmt.Println("\n\n--------------查找地址在广州并且店铺评分大于4.95分的皇冠以上店铺名称----------------")
	t = time.Now()
	func() {
		// 1)首先取出在广州地区的所有店铺id
		rs, _ := redisTable.Rconn.Do("lrange", "shop:addr=" + "广东广州.id", "0", "-1")
		addrIds, _ := redisTable.Gsstring(rs)
		// 2)获取店铺评分大于4.95的店铺id
		rs, _ = redisTable.Rconn.Do("zrangebyscore", "shop:score", "4.95", "5.1")
		scoreIds, _ := redisTable.Gsstring(rs)

		// 3)获取皇冠以上的店铺id
		// 获取一皇冠所对应的信用等级值
		rs, _ = redisTable.Rconn.Do("zscore", "credit", "一黄冠")
		min, _ := redisTable.Gstring(rs)
		rs, _ = redisTable.Rconn.Do("zscore", "credit", "天猫")
		max, _ := redisTable.Gstring(rs)
		// 通过值获取一皇冠店铺在有序集合的店铺id
		rs, _ = redisTable.Rconn.Do("zrangebyscore", "shop:credit", min, max)
		creditIds, _ := redisTable.Gsstring(rs)

		// 4)求交集方式一，写入redis求交集，方式二，直接扫描判断
		var interIds []string
		for _, addrId := range addrIds {
			for _, scoreId := range scoreIds {
				if addrId == scoreId {
					for _, creditId := range creditIds {
						if addrId == creditId {
							interIds = append(interIds, addrId)
						}
					}
				}
			}
		}

		// 5)输出店铺名称
		for _, shopid := range interIds {
			rs, _ = redisTable.Rconn.Do("hget", "shop:id=" + shopid, "name")
			name, _ := redisTable.Gstring(rs)
			fmt.Println(name)
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())


	// (5) 查看好评率99%以上的手机相关商品网址
	fmt.Println("\n\n--------------查看好评率99%以上的手机相关商品网址----------------")
	t = time.Now()
	func() {
		// 1)获取大于98%的商品id
		rs, _ := redisTable.Rconn.Do("zrangebyscore", "goods:praise", "99", "101")
		praiseIds, _ := redisTable.Gsstring(rs)

		// 2)获取名称匹配手机字段的商品id
		// 获取手机相关的商品名称key
		rs, _ = redisTable.Rconn.Do("keys", "goods:name=*手机*.id")
		names, _ := redisTable.Gsstring(rs)
		// 通过商品名key获取商品id
		var nameIds []string
		for _, name := range names {
			rs, _ := redisTable.Rconn.Do("lrange", name, "0", "-1")
			nameId, _ := redisTable.Gsstring(rs)
			nameIds = append(nameIds, nameId...)
		}

		// 3)求交集
		var interIds []string
		for _, praiseId := range praiseIds {
			for _, nameId := range nameIds {
				if praiseId == nameId {
					interIds = append(interIds, praiseId)
				}
			}
		}

		// 4) 输出地址
		for _, interId := range interIds {
			rs, _ := redisTable.Rconn.Do("hget", "goods:id=" + interId, "site")
			site, _ := redisTable.Gstring(rs)
			fmt.Println(site)
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())

	// (6) 查询至少有6个商品以上的店铺名称
	fmt.Println("\n\n--------------查询至少有6个商品以上的店铺名称----------------")
	t = time.Now()
	func() {
		// 1)获取店铺id名称key
		rs, _ := redisTable.Rconn.Do("keys", "goods:shop_id=*")
		shop_id_Keys, _ := redisTable.Gsstring(rs)

		// 2)获取商品数6个以上的店铺id
		var shop_ids []string
		for _, shop_id := range shop_id_Keys {
			rs, _ := redisTable.Rconn.Do("llen", shop_id)// 查询花费时间比较多
			size, _ := redisTable.Gint64(rs)
			if size >= 6 {
				shop_ids = append(shop_ids, strings.TrimRight(strings.TrimLeft(shop_id, "goods:shop_id="), ".id"))
			}
		}

		// 3)输出店铺名称
		for _, shopid := range shop_ids {
			rs, _ = redisTable.Rconn.Do("hget", "shop:id=" + shopid, "name")
			name, _ := redisTable.Gstring(rs)
			if name != "" {
				fmt.Println(name)
			}
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())


	// (7)查询各市场分类中好评度最好最新的商品的店铺名称
	fmt.Println("\n\n--------------查询各市场分类中好评度最好最新的商品的店铺名称----------------")
	t = time.Now()
	func() {
		// 1) 获取市场分类id
		rs, _ := redisTable.Rconn.Do("zrangebyscore", "market", "0", "100", "withscores")
		marketNameIds, _ := redisTable.Gsstring(rs)
		marketIdsMap := make(map[string]string)
		name := ""
		for n, nameId := range marketNameIds {
			if n % 2 == 0 {
				name = nameId
			} else {
				marketIdsMap[nameId] = name
			}
		}


		// 2) 获取好评度最好的商品id
		rs, _ = redisTable.Rconn.Do("zrangebyscore", "goods:praise", "100", "101")
		goodsIds, _ := redisTable.Gsstring(rs)


		// 3) 通过最好商品id进行市场分类id
		classify := make(map[string][]int)
		for _, goodsId := range goodsIds {
			// 从goodsId获取markeyId
			rs, _ = redisTable.Rconn.Do("hget", "goods:id=" + goodsId, "market_id")
			market_id, _ := redisTable.Gstring(rs)

			for marketId, _ := range marketIdsMap {
				if marketId == market_id {
					goodsIdInt, _ := strconv.Atoi(goodsId)
					if mk, ok := classify[marketId]; ok {
						mk = append(mk, goodsIdInt)
						classify[marketId] = mk
					} else {
						classify[marketId] = []int{goodsIdInt}
					}
				}
			}
		}


		// 4)在市场分类中的商品id取出最新商品id
		maxMarketGoodsId := make(map[string]string)
		for markId, ids := range classify {
			sort.Ints(ids)    // 排序
			maxMarketGoodsId[markId] = fmt.Sprintf("%d", ids[len(ids) - 1])
		}

		// 5)输出店铺名称
		for marketid, goodsid := range maxMarketGoodsId {
			rs, _ = redisTable.Rconn.Do("hget", "goods:id=" + goodsid, "shop_id")
			shop_id, _ := redisTable.Gstring(rs)
			rs, _ = redisTable.Rconn.Do("hget", "shop:id=" + shop_id, "name")
			name, _ := redisTable.Gstring(rs)
			fmt.Println(marketIdsMap[marketid] + "市场：", name)
		}
	}()
	fmt.Println("使用时间(ns)：", time.Now().Sub(t).Nanoseconds())
}

func main() {
	// 连接mysql和redis数据库
	redisTable.Init_mysql_redis()

	// 添加表
	//addRowData()

	// 查询测试
	testQuery()
}



