package insl

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(22)
	intSlice.Add(44)
	intSlice.Add(33)
	intSlice.Add(11)
	Convey("测试添加值", t, func() {
		So(fmt.Sprintf("%v", intSlice.GetAll()), ShouldEqual, "[11 22 33 44]")
	})
}

func TestSearch(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(30)
	intSlice.Add(11)
	intSlice.Add(52)
	intSlice.Add(23)
	intSlice.Add(44)

	index, ok := intSlice.Search(44)
	Convey("测试查找值44是否存在", t, func() {
		So(ok, ShouldEqual, true)
		So(index, ShouldEqual, 3)
	})

	index, ok = intSlice.Search(100)
	Convey("测试查找100是否存在", t, func() {
		So(ok, ShouldEqual, false)
		So(index, ShouldEqual, 5)
	})
}

func TestRemove(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(30)
	intSlice.Add(11)
	intSlice.Add(52)
	intSlice.Add(23)
	intSlice.Add(44)

	intSlice.Remove(30)

	_, ok := intSlice.Search(30)
	Convey("测试删除指定键", t, func() {
		So(ok, ShouldEqual, false)
	})
}

func TestClear(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(30)
	intSlice.Add(11)

	intSlice.Clear()

	Convey("测试清空所有值", t, func() {
		So(intSlice.Len(), ShouldEqual, 0)
	})
}

func TestGet(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(30)
	intSlice.Add(50)
	intSlice.Add(11)

	Convey("测试从指定索引获取值", t, func() {
		vl, _ := intSlice.Get(1)
		So(vl, ShouldEqual, 30)
		vl, _ = intSlice.Get(10)
		So(vl, ShouldEqual, 0)
	})
}

func TestGetAll(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(30)
	intSlice.Add(11)
	intSlice.Add(52)
	intSlice.Add(23)
	intSlice.Add(44)

	Convey("测试获取所有值", t, func() {
		So(fmt.Sprintf("%v", intSlice.GetAll()), ShouldEqual, "[11 23 30 44 52]")
	})
}

func TestString0(t *testing.T) {
	intSlice := NewIntSlice()
	intSlice.Add(30)
	intSlice.Add(11)
	intSlice.Add(52)

	Convey("测试获取值的字符串形式", t, func() {
		So(intSlice.String(), ShouldEqual, "{int:[11 30 52]}")
	})
}
