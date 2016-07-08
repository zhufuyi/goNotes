// 有序的int类型slice
package insl

import (
	"bytes"
	"fmt"
	"sort"
)

// 对外接口
type IntSlicer interface {
	Len() int                  // 元素个数
	Add(vl int) bool           // 添加一个元素
	Remove(vl int) bool        // 删除一个元素
	Clear()                    // 清空slice
	Get(index int) (int, bool) // 从索引获取元素
	GetAll() []int             // 获取所有元素
	Search(vl int) (int, bool) // 查找指定的元素，得到元素的索引
	String() string            // 元素的字符串形式
}

type intSlice struct {
	content []int
}

// 实例化
func NewIntSlice() IntSlicer {
	return &intSlice{content: make([]int, 0)}
}

// 元素个数
func (is *intSlice) Len() int {
	return len(is.content)
}

// 比较元素大小
func (is *intSlice) Less(i int, j int) bool {
	return is.content[i]-is.content[j] < 0
}

// 交换
func (is *intSlice) Swap(i int, j int) {
	is.content[i], is.content[j] = is.content[j], is.content[i]
}

// 查找指定的元素，得到元素的索引
func (is *intSlice) Search(vl int) (int, bool) {
	f := func(i int) bool {
		return is.content[i]-vl >= 0
	}
	index := sort.Search(is.Len(), f)
	if index < is.Len() && is.content[index] == vl {
		return index, true
	}
	return index, false
}

// 判断是否存在
func (is *intSlice) isExist(vl int) bool {
	_, ok := is.Search(vl)
	return ok
}

// 添加一个元素
func (is *intSlice) Add(vl int) bool {
	if is.isExist(vl) { // 如果实际需要唯一性，需要判断是否存在
		return false
	}
	is.content = append(is.content, vl)
	sort.Sort(is)
	return true
}

// 删除一个元素
func (is *intSlice) Remove(vl int) bool {
	index, ok := is.Search(vl)
	if ok {
		is.content = append(is.content[:index], is.content[index+1:]...) // 连接
		return true
	}
	return false
}

// 清空slice
func (is *intSlice) Clear() {
	is.content = make([]int, 0)
}

// 从索引获取元素
func (is *intSlice) Get(index int) (int, bool) {
	if index > is.Len()-1 || index < 0 {
		return 0, false
	}
	return is.content[index], true
}

// 获取所有元素
func (is *intSlice) GetAll() []int {
	return is.content
}

// 所有元素的字符串形式
func (is *intSlice) String() string {
	var buf bytes.Buffer
	buf.WriteString("{int:[")
	first := true
	for _, v := range is.content {
		if first {
			buf.WriteString(fmt.Sprintf("%d", v))
			first = false
		} else {
			buf.WriteString(fmt.Sprintf(" %d", v))
		}
	}
	buf.WriteString("]}")
	return buf.String()
}
