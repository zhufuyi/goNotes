// 有序的slice，支持能比较大小的类型，比如整形，字符串类型等

package oMap

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

// 对外接口
type Keyser interface {
	Len() int                                        // 元素的个数
	Add(k interface{}) bool                          // 添加一个元素
	Remove(k interface{}) bool                       // 删除一个元素
	Clear()                                          // 清空所有元素
	Get(index int) interface{}                       // 从索引获取元素
	GetAll() []interface{}                           // 获取所有元素
	Search(k interface{}) (index int, contains bool) // 查找指定的元素，得到元素的索引
	ElemType() reflect.Type                          // 元素的类型
	CompareFunc() CompareFunction                    // 比较元素大小
	String() string                                  // 元素的字符串形式
}

type myKeys struct {
	container   []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType    reflect.Type
}

type CompareFunction func(interface{}, interface{}) int8

// 实例化
func NewKeys(cf CompareFunction, et reflect.Type) Keyser {
	return &myKeys{
		container:   make([]interface{}, 0),
		compareFunc: cf,
		elemType:    et,
	}
}

// 字典元素个数
func (keys *myKeys) Len() int {
	return len(keys.container)
}

// 比较元素大小
func (keys *myKeys) Less(i int, j int) bool {
	return keys.compareFunc(keys.container[i], keys.container[j]) == -1
}

// 交换
func (keys *myKeys) Swap(i int, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

// 判断类型是否符合要求
func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	if reflect.TypeOf(k) != keys.elemType { //判断类型是否一致
		return false
	}
	return true
}

// 添加元素
func (keys *myKeys) Add(k interface{}) bool {
	if !keys.isAcceptableElem(k) {
		return false
	}

	//	for _, key := range keys.container { // 排除重复的key
	//		if key == k {
	//			return true
	//		}
	//	}
	if _, ok := keys.Search(k); ok { //判断添加的元素是否存在
		return true
	}

	keys.container = append(keys.container, k)
	sort.Sort(keys)
	return true
}

// 查找元素,返回元素的索引和查找结果
func (keys *myKeys) Search(k interface{}) (int, bool) {
	if !keys.isAcceptableElem(k) {
		return 0, false
	}

	f := func(i int) bool {
		return keys.compareFunc(keys.container[i], k) >= 0
	}
	index := sort.Search(keys.Len(), f) // 要求已经排序过的slice

	if index < keys.Len() && keys.container[index] == k {
		return index, true
	}
	return index, false
}

// 删除元素
func (keys *myKeys) Remove(k interface{}) bool {
	if !keys.isAcceptableElem(k) {
		return false
	}

	index, ok := keys.Search(k)
	if !ok {
		return false
	}

	keys.container = append(keys.container[:index], keys.container[index+1:]...)
	return true
}

// 清除所有元素
func (keys *myKeys) Clear() {
	keys.container = make([]interface{}, 0)
}

// 通过索引获取元素值
func (keys *myKeys) Get(index int) interface{} {
	if index >= keys.Len() || index < 0 {
		return nil
	}
	return keys.container[index]
}

// 获取所有元素值
func (keys *myKeys) GetAll() []interface{} {
	initLen := keys.Len()
	snapshot := make([]interface{}, initLen)
	actualLen := 0
	for _, key := range keys.container {
		if actualLen < initLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

// 获取key的类型
func (keys *myKeys) ElemType() reflect.Type {
	return keys.elemType
}

// 比较函数类型
func (keys *myKeys) CompareFunc() CompareFunction {
	return keys.compareFunc
}

// 获取元素的字符串形式
func (keys *myKeys) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("{%v:[", keys.elemType))
	first := true
	for _, key := range keys.container {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("]}")
	return buf.String()
}
