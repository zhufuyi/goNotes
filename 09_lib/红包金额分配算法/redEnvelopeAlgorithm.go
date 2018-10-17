package main

import (
	"math/rand"
	"time"
)

type RedEnvelope struct {
	remainNum   int // 剩余红包个数
	remainMoney int // 剩余魅钻数量
}

// 获取随机魅钻
func (r *RedEnvelope) getRandomMoney() int {
	if r.remainNum == 1 {
		return r.remainMoney
	}
	min := 1                               // 最小值
	max := r.remainMoney / r.remainNum * 2 // 最大值

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	money := rd.Intn(100) * max / 100 // 随机值
	if money < min {
		money = 1
	}
	r.remainNum--
	r.remainMoney -= money
	return money

}

// 判断是否分配红包完毕，true表示抢光了，false表示还有剩余红包可抢
func (r *RedEnvelope) isRobbedOver() bool {
	if r.remainNum == 0 {
		return true
	}
	return false
}

func main() {
	re := &RedEnvelope{10, 100}
	array := []int{}
	for i := 0; i < 10; i++ {
		if re.isRobbedOver() {
			break
		}
		val := re.getRandomMoney()
		array = append(array, val)
		println(val)
	}
	println(array)
}
