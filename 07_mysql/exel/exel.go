package exel

import (
	"github.com/tealeg/xlsx"
	"strconv"
	"math"
	"strings"
)


// 店铺信息
type Shop struct {
	ShopName string  // 店铺名称
	ShopSite string  // 店铺网址
	Credit   int     // 信用值
	Addr     string  // 店铺地址
	Score    float64 // 店铺评分
}

// 商品信息
type Goods struct {
	ShopId    int     // 店铺id
	ShopName  string  // 店铺名
	MarketId  int     // 市场分类id
	GoodsName string  // 商品名称
	Praise    float64 // 好评度
	GoodsSite string  // 商品网址
}

type Exeler interface {
	Extract() ([]*Shop, []*Goods, error)
}

func OpenExel(filename string) Exeler {
	return &exelV{
		filename:filename,
		content:make([][]string, 0),
	}
}

type exelV struct {
	filename string     // 文件名称
	content  [][]string // 存储exel文件各单元内容
}


// 获取exel文件的内容
func (e *exelV)readContent(filename string) error {
	// 打开exel文件
	xlFile, err := xlsx.OpenFile(e.filename)
	if err != nil {
		return err
	}

	// 按行读取exel内容
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 行
			s := make([]string, len(row.Cells))
			for n, cell := range row.Cells {
				// 单元
				s[n] = cell.Value
			}
			e.content = append(e.content, s)
		}
	}
	return nil
}



// 提取数据
func (e *exelV)Extract() ([]*Shop, []*Goods, error) {
	// 获取exel文件的内容
	err := e.readContent(e.filename)
	if err != nil {
		return nil, nil, err
	}

	var shops []*Shop
	var goods []*Goods

	for n, s := range e.content {
		if n == 0 {
			// 去掉第一行内容
			continue
		}

		sl := &Shop{}
		gd := &Goods{}
		if len(s) == 10 {
			credit := getCreditVal(s[3])
			if credit < 0 {
				continue
			}
			sl.Credit = credit
			sl.ShopName = s[1]
			sl.ShopSite = s[2]
			sl.Addr = s[6]
			score, _ := strconv.ParseFloat(s[8], 32)
			sl.Score = floatPointN(score, 2)

			gd.ShopName = s[1]
			//gd.GoodsName=s[4]
			gd.GoodsName = strings.Replace(s[4], "'", `\'`, -1)
			gd.Praise = cut2float(s[7])
			gd.GoodsSite = s[9]
		}
		if sl.ShopName != ""&&sl.ShopSite != ""&&sl.Addr != "" {
			// 去掉空内容的行
			shops = append(shops, sl)
		}
		if gd.GoodsName != ""&&gd.GoodsSite != "" {
			// 去掉空内容的行
			goods = append(goods, gd)
		}
	}
	return shops, goods, nil
}



// 获取店铺信用对应的序号
func getCreditVal(cn string) int {
	// 注：由于源文件的信用等级“天猫卖家”的值不唯一，通过这里统一值为21，id从1开始
	credit := []string{"一星卖家", "二星卖家", "三星卖家", "四星卖家", "五星卖家", "一钻卖家", "二钻卖家", "三钻卖家", "四钻", "五钻卖家", "一蓝冠卖家", "二蓝冠卖家", "三蓝冠卖家", "四蓝冠卖家", "五蓝冠卖家", "一黄冠卖家", "二黄冠卖家", "三黄冠卖家", "四黄冠卖家", "五黄冠卖家", "天猫卖家"}
	for index, v := range credit {
		if v == cn {
			return index + 1
		}
	}
	return -1
}

// 从字符串截取转浮点数
func cut2float(praise string) float64 {
	if praise == "" || praise == "null" {
		return floatPointN(0.0, 2)
	}
	pr, _ := strconv.ParseFloat(strings.TrimRight(praise, "%"), 32)
	return floatPointN(pr * 100, 2)
}

// 浮点数保留小数点位数
func floatPointN(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f + 0.5 / pow10_n) * pow10_n) / pow10_n
}
