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
	"reflect"
	"sync"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// ComponentLib 组件原型库
type ComponentLib interface {
	// Declare 声明组件原型
	Declare(comp any) ec.ComponentPT
	// Get 查询组件原型
	Get(prototype string) (ec.ComponentPT, bool)
	// List 获取所有组件原型
	List() []ec.ComponentPT
	// EventStream 组件声明事件流
	EventStream(ctx context.Context) <-chan ec.ComponentPT
}

var compLib = NewComponentLib()

// DefaultComponentLib 默认组件库
func DefaultComponentLib() ComponentLib {
	return compLib
}

// NewComponentLib 创建组件原型库
func NewComponentLib() ComponentLib {
	return &_ComponentLib{
		compPTNameIndex: map[string]int{},
	}
}

type _ComponentLib struct {
	sync.RWMutex
	compPTNameIndex map[string]int
	compPTList      generic.FreeList[ec.ComponentPT]
	eventStream     generic.EventStream[ec.ComponentPT]
}

// Declare 声明组件原型
func (lib *_ComponentLib) Declare(comp any) ec.ComponentPT {
	if comp == nil {
		exception.Panicf("%w: %w: comp is nil", ErrPt, exception.ErrArgs)
	}

	lib.Lock()
	defer lib.Unlock()

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

	compPTIdx, ok := lib.compPTNameIndex[prototype]
	if ok {
		return lib.compPTList.Get(compPTIdx).V
	}

	compPT := &_Component{
		prototype:  prototype,
		instanceRT: compRT,
	}
	compPT.builtin = &ec.BuiltinComponent{PT: compPT, Offset: -1}

	lib.compPTNameIndex[prototype] = lib.compPTList.PushBack(compPT).Index()

	lib.eventStream.Publish(compPT)

	return compPT
}

// Get 查询组件原型
func (lib *_ComponentLib) Get(prototype string) (ec.ComponentPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	compPTIdx, ok := lib.compPTNameIndex[prototype]
	if !ok {
		return nil, false
	}

	return lib.compPTList.Get(compPTIdx).V, true
}

// List 获取所有组件原型
func (lib *_ComponentLib) List() []ec.ComponentPT {
	lib.RLock()
	defer lib.RUnlock()

	return lib.compPTList.ToSlice()
}

// EventStream 组件声明事件流
func (lib *_ComponentLib) EventStream(ctx context.Context) <-chan ec.ComponentPT {
	if ctx == nil {
		ctx = context.Background()
	}

	lib.Lock()
	defer lib.Unlock()

	return lib.eventStream.Subscribe(ctx, lib.compPTList.ToSlice()...)
}
