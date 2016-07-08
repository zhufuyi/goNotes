// 有序的string类型slice
package stsl

import (
	"bytes"
	"fmt"
	"sort"
)

// 对外接口
type StrSlicer interface {
	Len() int                     // 元素个数
	Add(vl string) bool           // 添加一个元素
	Adds(vls ...string) []string  // 批量添加元素
	Remove(vl string) bool        // 删除一个元素
	Clear()                       // 清空slice
	Get(index int) (string, bool) // 从索引获取元素
	GetAll() []string             // 获取所有元素
	Search(vl string) (int, bool) // 查找指定的元素，得到元素的索引
	String() string               // 元素的字符串形式
}

type StrSlice struct {
	content []string
}

// 比较函数
var compareFunc = func(str1, str2 string) int {
	return bytes.Compare([]byte(str1), []byte(str2))
}

// 实例化
func NewStrSlice() StrSlicer {
	return &StrSlice{content: make([]string, 0)}
}

// 元素个数
func (ss *StrSlice) Len() int {
	return len(ss.content)
}

// 比较元素大小
func (ss *StrSlice) Less(i int, j int) bool {
	return compareFunc(ss.content[i], ss.content[j]) < 0 // 升序
}

// 交换
func (ss *StrSlice) Swap(i int, j int) {
	ss.content[i], ss.content[j] = ss.content[j], ss.content[i]
}

// 查找指定的元素，得到元素的索引
func (ss *StrSlice) Search(vl string) (int, bool) {
	f := func(i int) bool {
		return compareFunc(ss.content[i], vl) >= 0
	}
	index := sort.Search(ss.Len(), f)
	if index < ss.Len() && ss.content[index] == vl {
		return index, true
	}
	return index, false
}

// 判断是否存在
func (ss *StrSlice) isExist(vl string) bool {
	_, ok := ss.Search(vl)
	return ok
}

// 添加一个元素
func (ss *StrSlice) Add(vl string) bool {
	if ss.isExist(vl) { // 如果实际需要唯一性，需要判断是否存在
		return false
	}
	ss.content = append(ss.content, vl)
	sort.Sort(ss)
	return true
}

// 批量添加元素,返回添加失败的元素
func (ss *StrSlice) Adds(vls ...string) []string {
	var failStr []string
	for _, vl := range vls {
		if !ss.Add(vl) {
			failStr = append(failStr, vl)
		}
	}
	return failStr
}

// 删除一个元素
func (ss *StrSlice) Remove(vl string) bool {
	index, ok := ss.Search(vl)
	if ok {
		ss.content = append(ss.content[:index], ss.content[index+1:]...) // 连接
		return true
	}
	return false
}

// 清空slice
func (ss *StrSlice) Clear() {
	ss.content = make([]string, 0)
}

// 从索引获取元素
func (ss *StrSlice) Get(index int) (string, bool) {
	if index > ss.Len()-1 || index < 0 {
		return "", false
	}
	return ss.content[index], true
}

// 获取所有元素
func (ss *StrSlice) GetAll() []string {
	return ss.content
}

// 所有元素的字符串形式
func (ss *StrSlice) String() string {
	var buf bytes.Buffer
	buf.WriteString("{string:[")
	first := true
	for _, v := range ss.content {
		if first {
			buf.WriteString(fmt.Sprintf("%s", v))
			first = false
		} else {
			buf.WriteString(fmt.Sprintf(" %s", v))
		}
	}
	buf.WriteString("]}")
	return buf.String()
}
