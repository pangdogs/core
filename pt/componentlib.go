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
	"fmt"
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
	// Declare 声明组件原型
	Declare(comp any, aliases ...string) ComponentPT
	// Undeclare 取消声明组件原型
	Undeclare(name string)
	// Get 获取组件原型
	Get(name string) (ComponentPT, bool)
	// AliasGet 使用别名获取组件原型
	AliasGet(alias string) []ComponentPT
	// Range 遍历所有已注册的组件原型
	Range(fun generic.Func1[ComponentPT, bool])
	// ReversedRange 反向遍历所有已注册的组件原型
	ReversedRange(fun generic.Func1[ComponentPT, bool])
}

var compLib = NewComponentLib()

// DefaultComponentLib 默认组件库
func DefaultComponentLib() ComponentLib {
	return compLib
}

// NewComponentLib 创建组件原型库
func NewComponentLib() ComponentLib {
	return &_ComponentLib{
		compIdx:  map[string]*ComponentPT{},
		aliasTab: map[string][]*ComponentPT{},
	}
}

type _ComponentLib struct {
	sync.RWMutex
	compIdx  map[string]*ComponentPT
	compList []*ComponentPT
	aliasTab map[string][]*ComponentPT
}

// Declare 声明组件原型
func (lib *_ComponentLib) Declare(comp any, aliases ...string) ComponentPT {
	if comp == nil {
		panic(fmt.Errorf("%w: %w: comp is nil", ErrPt, exception.ErrArgs))
	}

	tfComp, ok := comp.(reflect.Type)
	if !ok {
		tfComp = reflect.TypeOf(comp)
	}

	return lib.declare(tfComp, aliases)
}

// Undeclare 取消声明组件原型
func (lib *_ComponentLib) Undeclare(name string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.compIdx, name)

	lib.compList = slices.DeleteFunc(lib.compList, func(pt *ComponentPT) bool {
		return pt.Name == name
	})

	lib.cleanAliases(name)
}

// Get 获取组件原型
func (lib *_ComponentLib) Get(name string) (ComponentPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	comp, ok := lib.compIdx[name]
	if !ok {
		return ComponentPT{}, false
	}

	return *comp, ok
}

// AliasGet 使用别名获取组件原型
func (lib *_ComponentLib) AliasGet(alias string) []ComponentPT {
	lib.RLock()
	defer lib.RUnlock()

	comps := lib.aliasTab[alias]
	ret := make([]ComponentPT, 0, len(comps))

	for _, comp := range comps {
		ret = append(ret, *comp)
	}

	return ret
}

// Range 遍历所有已注册的组件原型
func (lib *_ComponentLib) Range(fun generic.Func1[ComponentPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.compList)
	lib.RUnlock()

	for i := range copied {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}

// ReversedRange 反向遍历所有已注册的组件原型
func (lib *_ComponentLib) ReversedRange(fun generic.Func1[ComponentPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.compList)
	lib.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}

func (lib *_ComponentLib) declare(tfComp reflect.Type, aliases []string) ComponentPT {
	lib.Lock()
	defer lib.Unlock()

	for tfComp.Kind() == reflect.Pointer || tfComp.Kind() == reflect.Interface {
		tfComp = tfComp.Elem()
	}

	if tfComp.Name() == "" {
		panic(fmt.Errorf("%w: anonymous component not allowed", ErrPt))
	}

	compName := types.FullName(tfComp)

	if !reflect.PointerTo(tfComp).Implements(reflect.TypeFor[ec.Component]()) {
		panic(fmt.Errorf("%w: component %q not implement ec.Component", ErrPt, compName))
	}

	comp, ok := lib.compIdx[compName]
	if ok {
		lib.addAliases(comp, aliases)
		return *comp
	}

	comp = &ComponentPT{
		Name:       compName,
		InstanceRT: tfComp,
	}

	lib.compIdx[compName] = comp
	lib.compList = append(lib.compList, comp)
	lib.addAliases(comp, aliases)

	return *comp
}

func (lib *_ComponentLib) addAliases(comp *ComponentPT, aliases []string) {
	for _, alias := range aliases {
		if !slices.ContainsFunc(lib.aliasTab[alias], func(pt *ComponentPT) bool {
			return pt.Name == comp.Name
		}) {
			lib.aliasTab[alias] = append(lib.aliasTab[alias], comp)
		}
	}
}

func (lib *_ComponentLib) cleanAliases(compName string) {
	for alias, comps := range lib.aliasTab {
		lib.aliasTab[alias] = slices.DeleteFunc(comps, func(pt *ComponentPT) bool {
			return pt.Name == compName
		})
	}
}
