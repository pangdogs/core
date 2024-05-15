package event

import (
	"fmt"
	"git.golaxy.org/core/util/types"
	"hash/fnv"
	"reflect"
	"sync"
)

// IEventTab 本地事件表接口，方便管理多个事件
/*
使用方式：
	1.在定义事件的源码文件（.go）头部添加以下注释，在编译前自动化生成代码：
	//go:generate go run git.golaxy.org/core/event/eventcode gen_eventtab --name={事件表名称}

定义事件的选项（添加到定义事件的注释里）：
	1.事件表初始化时，该事件使用的递归处理方式，不填表示使用事件表初始化参数值
		[EventRecursion_Allow]
		[EventRecursion_Disallow]
		[EventRecursion_Discard]
		[EventRecursion_Truncate]
		[EventRecursion_Deepest]
*/
type IEventTab interface {
	IEventCtrl
	// Get 获取事件
	Get(id uint64) IEvent
}

// MakeEventTabId 创建事件表Id
func MakeEventTabId(eventTab IEventTab) uint64 {
	hash := fnv.New32a()
	rt := reflect.ValueOf(eventTab).Type()
	if rt.PkgPath() == "" || rt.Name() == "" {
		panic("unsupported type")
	}
	hash.Write([]byte(types.FullNameRT(rt)))
	return uint64(hash.Sum32()) << 32
}

// MakeEventTabIdT 创建事件表Id
func MakeEventTabIdT[T any]() uint64 {
	hash := fnv.New32a()
	rt := reflect.TypeFor[T]()
	if rt.PkgPath() == "" || rt.Name() == "" || !reflect.PointerTo(rt).Implements(reflect.TypeFor[IEventTab]()) {
		panic("unsupported type")
	}
	hash.Write([]byte(types.FullNameRT(rt)))
	return uint64(hash.Sum32()) << 32
}

// MakeEventId 创建事件Id
func MakeEventId(eventTab IEventTab, pos int32) uint64 {
	return MakeEventTabId(eventTab) + uint64(pos)
}

// MakeEventIdT 创建事件Id
func MakeEventIdT[T any](pos int32) uint64 {
	return MakeEventTabIdT[T]() + uint64(pos)
}

var (
	declareEventTabs = &sync.Map{}
	declareEvents    = &sync.Map{}
)

// DeclareEventTabId 声明事件表Id
func DeclareEventTabId(eventTab IEventTab) uint64 {
	id := MakeEventTabId(eventTab)
	if name, loaded := declareEventTabs.LoadOrStore(id, types.FullNameRT(reflect.TypeOf(eventTab).Elem())); loaded {
		panic(fmt.Errorf("event_tab(%d) has already been declared by %q", id, name))
	}
	return id
}

// DeclareEventTabIdT 声明事件表Id
func DeclareEventTabIdT[T any]() uint64 {
	id := MakeEventTabIdT[T]()
	if name, loaded := declareEventTabs.LoadOrStore(id, types.FullNameRT(reflect.TypeFor[T]())); loaded {
		panic(fmt.Errorf("event_tab(%d) has already been declared by %q", id, name))
	}
	return id
}

// DeclareEventId 声明事件Id
func DeclareEventId(eventTab IEventTab, pos int32) uint64 {
	id := MakeEventTabId(eventTab) + uint64(pos)
	if name, loaded := declareEvents.LoadOrStore(id, types.FullNameRT(reflect.TypeOf(eventTab).Elem())); loaded {
		panic(fmt.Errorf("event(%d) has already been declared by %q", id, name))
	}
	return id
}

// DeclareEventIdT 声明事件Id
func DeclareEventIdT[T any](pos int32) uint64 {
	id := MakeEventTabIdT[T]() + uint64(pos)
	if name, loaded := declareEvents.LoadOrStore(id, types.FullNameRT(reflect.TypeFor[T]())); loaded {
		panic(fmt.Errorf("event(%d) has already been declared by %q", id, name))
	}
	return id
}
