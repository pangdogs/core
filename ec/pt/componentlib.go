/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package pt

import (
	"context"
	"maps"
	"reflect"
	"slices"
	"sync"
	"sync/atomic"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// ComponentLib 是可并发访问的组件原型注册表。
type ComponentLib interface {
	// Declare 声明组件原型；同一具名类型重复声明时返回已有原型。
	Declare(comp any) ec.ComponentPT
	// Get 按完整原型名查询组件原型。
	Get(prototype string) (ec.ComponentPT, bool)
	// List 返回当前全部组件原型的快照。
	List() []ec.ComponentPT
	// Watch 依次发送当前快照及后续的新声明；ctx 结束后排空已排队项并关闭频道。
	Watch(ctx context.Context) <-chan ec.ComponentPT
}

var compLib = NewComponentLib()

// DefaultComponentLib 返回进程级默认组件原型库。
func DefaultComponentLib() ComponentLib {
	return compLib
}

// NewComponentLib 创建独立的空组件原型库。
func NewComponentLib() ComponentLib {
	lib := &_ComponentLib{}
	lib.snapshot.Store(&_ComponentLibSnapshot{
		compPTIndex: map[string]ec.ComponentPT{},
	})
	return lib
}

type _ComponentLib struct {
	mutex       sync.Mutex
	snapshot    atomic.Pointer[_ComponentLibSnapshot]
	eventStream generic.EventStream[ec.ComponentPT]
}

// _ComponentLibSnapshot 是组件原型库的只读快照，发布后不再修改。
type _ComponentLibSnapshot struct {
	compPTIndex map[string]ec.ComponentPT
	compPTList  []ec.ComponentPT
}

func (snapshot *_ComponentLibSnapshot) clone() *_ComponentLibSnapshot {
	return &_ComponentLibSnapshot{
		compPTIndex: maps.Clone(snapshot.compPTIndex),
		compPTList:  slices.Clone(snapshot.compPTList),
	}
}

// Declare 声明组件原型；同一具名类型重复声明时返回已有原型。
//
// comp 可以是组件值或 reflect.Type。匿名类型、nil 以及未实现 ec.Component 的类型
// 会导致 panic。
func (lib *_ComponentLib) Declare(comp any) ec.ComponentPT {
	if comp == nil {
		exception.Panicf("%w: %w: comp is nil", ErrPt, exception.ErrArgs)
	}

	compRT, ok := comp.(reflect.Type)
	if !ok {
		compRT = reflect.TypeOf(comp)
	}

	for compRT.Kind() == reflect.Pointer {
		compRT = compRT.Elem()
	}

	if compRT.Name() == "" {
		exception.Panicf("%w: anonymous component not allowed", ErrPt)
	}

	prototype := types.FullNameRT(compRT)

	if !reflect.PointerTo(compRT).Implements(reflect.TypeFor[ec.Component]()) {
		exception.Panicf("%w: component %q not implement ec.Component", ErrPt, prototype)
	}

	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	snapshot := lib.snapshot.Load()
	if compPT, ok := snapshot.compPTIndex[prototype]; ok {
		return compPT
	}

	compPT := &_Component{
		prototype:  prototype,
		instanceRT: compRT,
	}
	compPT.builtin = &ec.BuiltinComponent{PT: compPT, Offset: -1}

	next := snapshot.clone()
	next.compPTIndex[prototype] = compPT
	next.compPTList = append(next.compPTList, compPT)
	lib.snapshot.Store(next)

	lib.eventStream.Publish(compPT)

	return compPT
}

// Get 按完整原型名查询组件原型。
func (lib *_ComponentLib) Get(prototype string) (ec.ComponentPT, bool) {
	compPT, ok := lib.snapshot.Load().compPTIndex[prototype]
	return compPT, ok
}

// List 返回当前全部组件原型的快照。
func (lib *_ComponentLib) List() []ec.ComponentPT {
	return slices.Clone(lib.snapshot.Load().compPTList)
}

// Watch 依次发送当前快照及后续的新声明；ctx 结束后排空已排队项并关闭频道。
// nil ctx 按 context.Background 处理。
func (lib *_ComponentLib) Watch(ctx context.Context) <-chan ec.ComponentPT {
	if ctx == nil {
		ctx = context.Background()
	}

	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	return lib.eventStream.Subscribe(ctx, lib.snapshot.Load().compPTList...)
}
