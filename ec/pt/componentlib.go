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
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
	"reflect"
	"slices"
	"sync"
)

// ComponentLib 组件原型库
type ComponentLib interface {
	// Declare 声明组件原型（可以重复声明）
	Declare(comp any) ec.ComponentPT
	// Undeclare 取消声明组件原型
	Undeclare(prototype string)
	// Get 获取组件原型
	Get(prototype string) (ec.ComponentPT, bool)
	// Range 遍历所有已注册的组件原型
	Range(fun generic.Func1[ec.ComponentPT, bool])
	// ReversedRange 反向遍历所有已注册的组件原型
	ReversedRange(fun generic.Func1[ec.ComponentPT, bool])
}

var compLib = NewComponentLib()

// DefaultComponentLib 默认组件库
func DefaultComponentLib() ComponentLib {
	return compLib
}

// NewComponentLib 创建组件原型库
func NewComponentLib() ComponentLib {
	return &_ComponentLib{
		compIndex: map[string]*_Component{},
	}
}

type _ComponentLib struct {
	sync.RWMutex
	compIndex map[string]*_Component
	compList  []*_Component
}

// Declare 声明组件原型（可以重复声明）
func (lib *_ComponentLib) Declare(comp any) ec.ComponentPT {
	return lib.declare(comp)
}

// Undeclare 取消声明组件原型
func (lib *_ComponentLib) Undeclare(prototype string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.compIndex, prototype)

	lib.compList = slices.DeleteFunc(lib.compList, func(pt *_Component) bool {
		return pt.Prototype() == prototype
	})
}

// Get 获取组件原型
func (lib *_ComponentLib) Get(prototype string) (ec.ComponentPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	compPT, ok := lib.compIndex[prototype]
	if !ok {
		return nil, false
	}

	return compPT, ok
}

// Range 遍历所有已注册的组件原型
func (lib *_ComponentLib) Range(fun generic.Func1[ec.ComponentPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.compList)
	lib.RUnlock()

	for i := range copied {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

// ReversedRange 反向遍历所有已注册的组件原型
func (lib *_ComponentLib) ReversedRange(fun generic.Func1[ec.ComponentPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.compList)
	lib.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

func (lib *_ComponentLib) declare(comp any) ec.ComponentPT {
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

	compPT, ok := lib.compIndex[prototype]
	if ok {
		return compPT
	}

	compPT = &_Component{
		prototype:  prototype,
		instanceRT: compRT,
	}
	compPT.builtin = &ec.BuiltinComponent{PT: compPT, Offset: -1}

	lib.compIndex[prototype] = compPT
	lib.compList = append(lib.compList, compPT)

	return compPT
}
