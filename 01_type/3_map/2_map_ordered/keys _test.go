package oMap

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func newKeys() Keyser {
	var cf = func(e1 interface{}, e2 interface{}) int8 {
		k1 := e1.(int)
		k2 := e2.(int)
		if k1 < k2 {
			return -1
		} else if k1 > k2 {
			return 1
		} else {
			return 0
		}
	}

	return NewKeys(cf, reflect.TypeOf(1))
}

func TestAdd(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(11)
	Convey("测试添加键", t, func() {
		So(fmt.Sprintf("%v", intKeys.GetAll()), ShouldEqual, "[11 30]")
	})
}

func TestSearch(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(11)
	intKeys.Add(52)
	intKeys.Add(23)
	intKeys.Add(44)

	index, ok := intKeys.Search(44)
	Convey("测试查找键是否存在1", t, func() {
		So(ok, ShouldEqual, true)
		So(index, ShouldEqual, 3)
	})

	index, ok = intKeys.Search(100)
	Convey("测试查找键是否存在2", t, func() {
		So(ok, ShouldEqual, false)
		So(index, ShouldEqual, 5)
	})
}

func TestRemove0(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(11)

	intKeys.Remove(30)

	_, ok := intKeys.Search(30)
	Convey("测试删除指定键", t, func() {
		So(ok, ShouldEqual, false)
	})
}

func TestClear0(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(11)

	intKeys.Clear()

	Convey("测试清除所有键", t, func() {
		So(intKeys.Len(), ShouldEqual, 0)
	})
}

func TestGet0(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(50)
	intKeys.Add(11)

	Convey("测试获取指定索引键", t, func() {
		So(intKeys.Get(1), ShouldEqual, 30)
		So(intKeys.Get(10), ShouldBeNil)
	})
}

func TestGetAll(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(11)
	intKeys.Add(52)
	intKeys.Add(23)
	intKeys.Add(44)

	Convey("测试获取所有键", t, func() {
		So(fmt.Sprintf("%v", intKeys.GetAll()), ShouldEqual, "[11 23 30 44 52]")
	})
}

func TestElemType(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)

	Convey("测试获取键的类型", t, func() {
		So(intKeys.ElemType().String(), ShouldEqual, "int")
	})
}

func TestString0(t *testing.T) {
	intKeys := newKeys()
	intKeys.Add(30)
	intKeys.Add(11)
	intKeys.Add(52)

	Convey("测试获取键的字符串形式", t, func() {
		So(intKeys.String(), ShouldEqual, "{int:[11 30 52]}")
	})
}
