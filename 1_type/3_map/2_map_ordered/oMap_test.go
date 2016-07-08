package oMap

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func newOmap() OrderedMaper {
	var cmpr = func(e1 interface{}, e2 interface{}) int8 {
		k1 := e1.(string)
		k2 := e2.(string)
		if k1 < k2 {
			return -1
		} else if k1 > k2 {
			return 1
		} else {
			return 0
		}
	}
	var kType = reflect.TypeOf("a")
	var vType = reflect.TypeOf(1)

	return NewOrderedMap(NewKeys(cmpr, kType), vType)
}

func TestPut(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("lisi", 200)

	Convey("测试添加键值", t, func() {
		So(o_Map.Get("lisi"), ShouldEqual, 200)
	})
}

func TestRemove(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhangsan", 100)
	o_Map.Remove("zhangsan")

	Convey("测试删除指定键值", t, func() {
		So(o_Map.Get("zhangsan"), ShouldBeNil)
	})
}

func TestClear(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhangsan", 100)
	o_Map.Put("lisi", 200)
	o_Map.Clear()

	Convey("测试清空键值", t, func() {
		So(o_Map.Get("zhangsan"), ShouldBeNil)
		So(o_Map.Get("lisi"), ShouldBeNil)
		So(o_Map.Len(), ShouldEqual, 0)
	})
}

func TestMapLen(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhangsan", 100)
	o_Map.Put("lisi", 200)
	o_Map.Put("wangwu", 300)

	Convey("测试map键值对数量", t, func() {
		So(o_Map.Len(), ShouldEqual, 3)
	})
}

func TestGet(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhangsan", 100)
	Convey("测试获取map值", t, func() {
		So(o_Map.Get("zhangsan"), ShouldEqual, 100)
	})
}

func TestIsExist(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhangsan", 100)

	Convey("测试map的键是否存在", t, func() {
		So(o_Map.IsExist("zhangsan"), ShouldBeTrue)
		So(o_Map.IsExist("lisi"), ShouldBeFalse)
	})
}

func TestFirstKey(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhangsan", 100)
	o_Map.Put("lisi", 200)
	o_Map.Put("wangwu", 300)

	if o_Map.FirstKey() != 200 {
		t.Error("测试失败")
	}
	Convey("测试有序map的第一个键值", t, func() {
		So(o_Map.FirstKey(), ShouldEqual, 200)
	})
}

func TestLastKey(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("wangwu", 300)
	o_Map.Put("zhangsan", 100)
	o_Map.Put("lisi", 200)

	Convey("测试有序map的最后一个键值", t, func() {
		So(o_Map.LastKey(), ShouldEqual, 100)
	})
}

func TestSubOMap(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("e", 150)
	o_Map.Put("a", 300)
	o_Map.Put("c", 100)
	o_Map.Put("g", 200)

	Convey("测试输出指定key范围的有序map,包括边界", t, func() {
		So(o_Map.SubOMap("A", "Z").String(), ShouldEqual, "OrderedMap<string:int>[]")
		So(o_Map.SubOMap("Z", "a").String(), ShouldEqual, "OrderedMap<string:int>[a:300]")
		So(o_Map.SubOMap("a", "a").String(), ShouldEqual, "OrderedMap<string:int>[a:300]")
		So(o_Map.SubOMap("a", "b").String(), ShouldEqual, "OrderedMap<string:int>[a:300]")
		So(o_Map.SubOMap("b", "e").String(), ShouldEqual, "OrderedMap<string:int>[c:100 e:150]")
		So(o_Map.SubOMap("g", "z").String(), ShouldEqual, "OrderedMap<string:int>[g:200]")
	})
}

func TestHeadOMap(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("e", 150)
	o_Map.Put("a", 300)
	o_Map.Put("c", 100)
	o_Map.Put("g", 200)

	Convey("测试输出<=key的有序map", t, func() {
		So(o_Map.HeadOMap("Z").String(), ShouldEqual, "OrderedMap<string:int>[]")
		So(o_Map.HeadOMap("a").String(), ShouldEqual, "OrderedMap<string:int>[a:300]")
		So(o_Map.HeadOMap("b").String(), ShouldEqual, "OrderedMap<string:int>[a:300]")
		So(o_Map.HeadOMap("g").String(), ShouldEqual, "OrderedMap<string:int>[a:300 c:100 e:150 g:200]")
		So(o_Map.HeadOMap("z").String(), ShouldEqual, "OrderedMap<string:int>[a:300 c:100 e:150 g:200]")
	})
}

func TestTailOMap(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("e", 150)
	o_Map.Put("a", 300)
	o_Map.Put("c", 100)
	o_Map.Put("g", 200)

	Convey("测试输出>=key的有序map", t, func() {
		So(o_Map.TailOMap("Z").String(), ShouldEqual, "OrderedMap<string:int>[a:300 c:100 e:150 g:200]")
		So(o_Map.TailOMap("a").String(), ShouldEqual, "OrderedMap<string:int>[a:300 c:100 e:150 g:200]")
		So(o_Map.TailOMap("b").String(), ShouldEqual, "OrderedMap<string:int>[c:100 e:150 g:200]")
		So(o_Map.TailOMap("g").String(), ShouldEqual, "OrderedMap<string:int>[g:200]")
		So(o_Map.TailOMap("z").String(), ShouldEqual, "OrderedMap<string:int>[]")
	})
}

func TestString(t *testing.T) {
	o_Map := newOmap()
	o_Map.Put("zhang", 100)
	o_Map.Put("li", 200)

	Convey("测试有序map的字符串形式", t, func() {
		So(o_Map.String(), ShouldEqual, "OrderedMap<string:int>[li:200 zhang:100]")
	})
}
