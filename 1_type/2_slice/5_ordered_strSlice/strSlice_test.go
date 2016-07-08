package stsl

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")
	Convey("测试添加值", t, func() {
		So(fmt.Sprintf("%v", ss.GetAll()), ShouldEqual, "[Abc aBc abC abc]")
	})
}

func TestSearch(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")

	index, ok := ss.Search("abc")
	Convey("测试查找\"abc\"是否存在", t, func() {
		So(ok, ShouldEqual, true)
		So(index, ShouldEqual, 3)
	})

	index, ok = ss.Search("ABC")
	Convey("测试查找\"ABC\"是否存在", t, func() {
		So(ok, ShouldEqual, false)
	})
}

func TestRemove(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")

	ss.Remove("abc")

	_, ok := ss.Search("abc")
	Convey("测试删除指定键", t, func() {
		So(ok, ShouldEqual, false)
	})
}

func TestClear(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")

	ss.Clear()

	Convey("测试清空所有值", t, func() {
		So(ss.Len(), ShouldEqual, 0)
	})
}

func TestGet(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")

	Convey("测试从指定索引获取值", t, func() {
		vl, _ := ss.Get(1)
		So(vl, ShouldEqual, "aBc")
		vl, _ = ss.Get(10)
		So(vl, ShouldEqual, "")
	})
}

func TestGetAll(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")

	Convey("测试获取所有值", t, func() {
		So(fmt.Sprintf("%v", ss.GetAll()), ShouldEqual, "[Abc aBc abC abc]")
	})
}

func TestString0(t *testing.T) {
	ss := NewStrSlice()
	ss.Adds("aBc", "abc", "Abc", "abC")

	Convey("测试获取值的字符串形式", t, func() {
		So(ss.String(), ShouldEqual, "{string:[Abc aBc abC abc]}")
	})
}
