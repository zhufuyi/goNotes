// 有序的slice，支持能比较大小的类型，比如整形，字符串类型等

package orsl

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

// 对外接口
type SliceOrder interface {
	sort.Interface
	Add(k interface{}) bool                          // 添加一个元素
	Remove(k interface{}) bool                       // 删除一个元素
	Clear()                                          // 清空slice
	Get(index int) interface{}                       // 从索引获取元素
	GetAll() []interface{}                           // 获取所有元素
	Search(k interface{}) (index int, contains bool) // 查找指定的元素，得到元素的索引
	ElemType() reflect.Type                          // 元素的类型
	CompareFunc() CompareFunction                    // 比较元素大小
	String() string                                  // 元素的字符串形式
}

type slicer struct {
	container   []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType    reflect.Type
}

type CompareFunction func(interface{}, interface{}) int8

// 实例化
func NewSlice(cf CompareFunction, et reflect.Type) SliceOrder {
	return &slicer{
		container:   make([]interface{}, 0),
		compareFunc: cf,
		elemType:    et,
	}
}

// slice元素个数
func (sl *slicer) Len() int {
	return len(sl.container)
}

// 比较元素大小
func (sl *slicer) Less(i int, j int) bool {
	return sl.compareFunc(sl.container[i], sl.container[j]) == -1
}

// 交换
func (sl *slicer) Swap(i int, j int) {
	sl.container[i], sl.container[j] = sl.container[j], sl.container[i]
}

// 判断类型是否符合要求
func (sl *slicer) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	if reflect.TypeOf(k) != sl.elemType { //判断类型是否一致
		return false
	}
	return true
}

// 添加元素
func (sl *slicer) Add(k interface{}) bool {
	if !sl.isAcceptableElem(k) {
		return false
	}

	//	for _, key := range sl.container { // 排除重复的key
	//		if key == k {
	//			return true
	//		}
	//	}
	if _, ok := sl.Search(k); ok { //判断添加的元素是否存在
		return true
	}

	sl.container = append(sl.container, k)
	sort.Sort(sl) // 没插入一个元素就排序一次
	return true
}

// 查找元素,返回元素的索引和查找结果
func (sl *slicer) Search(k interface{}) (int, bool) {
	if !sl.isAcceptableElem(k) {
		return 0, false
	}

	f := func(i int) bool {
		return sl.compareFunc(sl.container[i], k) >= 0
	}
	index := sort.Search(sl.Len(), f) // 要求已经排序过的slice

	if index < sl.Len() && sl.container[index] == k {
		return index, true
	}
	return index, false
}

// 删除元素
func (sl *slicer) Remove(k interface{}) bool {
	if !sl.isAcceptableElem(k) {
		return false
	}

	index, ok := sl.Search(k)
	if !ok {
		return false
	}

	sl.container = append(sl.container[:index], sl.container[index+1:]...) // 连接
	return true
}

// 清除所有元素
func (sl *slicer) Clear() {
	sl.container = make([]interface{}, 0)
}

// 通过索引获取元素值
func (sl *slicer) Get(index int) interface{} {
	if index >= sl.Len() || index < 0 {
		return nil
	}
	return sl.container[index]
}

// 获取所有元素值
func (sl *slicer) GetAll() []interface{} {
	initLen := sl.Len()
	snapshot := make([]interface{}, initLen)
	actualLen := 0
	for _, key := range sl.container {
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
func (sl *slicer) ElemType() reflect.Type {
	return sl.elemType
}

// 比较函数类型
func (sl *slicer) CompareFunc() CompareFunction {
	return sl.compareFunc
}

// 获取元素的字符串形式
func (sl *slicer) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("{%v:[", sl.elemType))
	first := true
	for _, key := range sl.container {
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
