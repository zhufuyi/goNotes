// 有序的map

package oMap

import (
	"bytes"
	"fmt"
	"reflect"
)

// 对外接口
type OrderedMaper interface {
	Put(k interface{}, v interface{}) (interface{}, bool) // 添加一个键值对
	Remove(k interface{}) bool                            // 删除一个键值对
	Clear()                                               // 清空所有键值对
	Get(k interface{}) interface{}                        // 通过键获取值
	Len() int                                             // 键值对个数
	IsExist(k interface{}) bool                           // 判断键是否存在
	KeyType() reflect.Type                                // 键的类型
	ValueType() reflect.Type                              // 值的类型
	FirstKey() interface{}                                // 第一个键对应的值
	LastKey() interface{}                                 // 最后一个键对应的值
	String() string                                       // 键值对的字符形式

	SubOMap(startKey interface{}, endKey interface{}) OrderedMaper // 截取指定键的范围（包括边界）作为新的有序map
	HeadOMap(startKey interface{}) OrderedMaper                    // 截取指定键到结束（包括边界）作为新的有序map
	TailOMap(endKey interface{}) OrderedMaper                      // 截取开始到指定键（包括边界）作为新的有序map
}

type OrderedMap struct {
	keyser   Keyser
	elemType reflect.Type
	m        map[interface{}]interface{}
}

// 实例化
func NewOrderedMap(ks Keyser, et reflect.Type) OrderedMaper {
	return &OrderedMap{
		keyser:   ks,
		elemType: et,
		m:        make(map[interface{}]interface{}),
	}
}

// 添加键值对,返回旧的的值和添加结果
func (omap *OrderedMap) Put(k interface{}, v interface{}) (interface{}, bool) {
	_, ok := omap.keyser.Search(k)
	if ok {
		oldV := omap.m[k]
		omap.m[k] = v
		return oldV, true
	} else {
		if ok := omap.keyser.Add(k); ok {
			omap.m[k] = v
			return nil, true
		}
		return nil, false
	}
}

// 删除指定的键值对
func (omap *OrderedMap) Remove(k interface{}) bool {
	if ok := omap.keyser.Remove(k); ok {
		delete(omap.m, k)
		return true
	}
	return false
}

// 清除所有键值对
func (omap *OrderedMap) Clear() {
	omap.keyser.Clear()
	omap.m = make(map[interface{}]interface{})
}

// 获取键对应的元素的值
func (omap *OrderedMap) Get(k interface{}) interface{} {
	if v, ok := omap.m[k]; ok {
		return v
	} else {
		return nil
	}
}

// 获取键值对数量
func (omap *OrderedMap) Len() int {
	return len(omap.m)
}

// 判断是否包含给定的键值
func (omap *OrderedMap) IsExist(k interface{}) bool {
	_, ok := omap.m[k]
	return ok
}

// 获取键的类型
func (omap *OrderedMap) KeyType() reflect.Type {
	return omap.keyser.ElemType()
}

// 获取值的类型
func (omap *OrderedMap) ValueType() reflect.Type {
	return omap.elemType
}

// 获取第一个键值
func (omap *OrderedMap) FirstKey() interface{} {
	if omap.keyser.Len() > 0 {
		return omap.m[omap.keyser.Get(0)]
	} else {
		return nil
	}
}

// 获取最后一个键值
func (omap *OrderedMap) LastKey() interface{} {
	size := omap.keyser.Len()
	if size > 0 {
		return omap.m[omap.keyser.Get(size-1)]
	} else {
		return nil
	}
}

// 获取omap多所有键值对的字符串形式
func (omap *OrderedMap) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("OrderedMap<%v:%v>[", omap.keyser.ElemType(), omap.elemType))

	first := true
	var key, val interface{}
	for i := 0; i < omap.Len(); i++ {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		key = omap.keyser.Get(i)
		val = omap.m[key]
		buf.WriteString(fmt.Sprintf("%v:%v", key, val))
	}
	buf.WriteString("]")
	return buf.String()
}

// 截取指定键的范围（包括边界）作为新的有序map
func (omap *OrderedMap) SubOMap(startKey interface{}, endKey interface{}) OrderedMaper {
	o_Map := &OrderedMap{
		keyser:   NewKeys(omap.keyser.CompareFunc(), omap.keyser.ElemType()),
		elemType: omap.elemType,
		m:        make(map[interface{}]interface{}),
	}

	if omap.Len() == 0 {
		return o_Map
	}

	startIndex, ok := omap.keyser.Search(startKey) // 搜索起始的key是否存在
	if !ok {
		if startIndex < 0 {
			startIndex = 0
		}
		if startIndex > omap.keyser.Len() {
			startIndex = omap.keyser.Len()
		}
	}

	endIndex, ok := omap.keyser.Search(endKey) // 搜索起始的key是否存在
	if !ok {
		if startIndex < 0 {
			startIndex = 0
		}
		if endIndex > omap.keyser.Len() {
			endIndex = omap.keyser.Len()
		}
	} else {
		//		if endIndex == omap.keyser.Len()-1 { // 为了统一遍历
		//			endIndex = omap.keyser.Len()
		//		}
		endIndex++
	}

	var key, elem interface{}
	for i := startIndex; i < endIndex; i++ {
		key = omap.keyser.Get(i)
		elem = omap.m[key]
		o_Map.Put(key, elem)
	}

	return o_Map
}

// 截取指定键到结束（包括边界）作为新的有序map
func (omap *OrderedMap) HeadOMap(endKey interface{}) OrderedMaper {
	return omap.SubOMap(nil, endKey)
}

// 截取开始到指定键（包括边界）作为新的有序map
func (omap *OrderedMap) TailOMap(startKey interface{}) OrderedMaper {
	size := omap.keyser.Len()
	if size == 0 {
		return omap.SubOMap(nil, nil)
	} else {
		return omap.SubOMap(startKey, omap.keyser.Get(size-1))
	}
}
