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
	"github.com/elliotchance/pie/v2"
	"reflect"
	"slices"
	"sync"
)

// EntityLib 实体原型库
type EntityLib interface {
	EntityPTProvider

	// GetComponentLib 获取组件原型库
	GetComponentLib() ComponentLib
	// Declare 声明实体原型
	Declare(prototype any, comps ...any) ec.EntityPT
	// Redeclare 重声明实体原型
	Redeclare(prototype any, comps ...any) ec.EntityPT
	// Undeclare 取消声明实体原型
	Undeclare(prototype string)
	// Get 获取实体原型
	Get(prototype string) (ec.EntityPT, bool)
	// Range 遍历所有已注册的实体原型
	Range(fun generic.Func1[ec.EntityPT, bool])
	// ReversedRange 反向遍历所有已注册的实体原型
	ReversedRange(fun generic.Func1[ec.EntityPT, bool])
}

var entityLib = NewEntityLib(DefaultComponentLib())

// DefaultEntityLib 默认实体库
func DefaultEntityLib() EntityLib {
	return entityLib
}

// NewEntityLib 创建实体原型库
func NewEntityLib(compLib ComponentLib) EntityLib {
	if compLib == nil {
		exception.Panicf("%w: %w: compLib is nil", ErrPt, exception.ErrArgs)
	}

	return &_EntityLib{
		compLib:   compLib,
		entityIdx: map[string]*_Entity{},
	}
}

type _EntityLib struct {
	sync.RWMutex
	compLib    ComponentLib
	entityIdx  map[string]*_Entity
	entityList []*_Entity
}

// GetEntityLib 获取实体原型库
func (lib *_EntityLib) GetEntityLib() EntityLib {
	return lib
}

// GetComponentLib 获取组件原型库
func (lib *_EntityLib) GetComponentLib() ComponentLib {
	return lib.compLib
}

// Declare 声明实体原型
func (lib *_EntityLib) Declare(prototype any, comps ...any) ec.EntityPT {
	return lib.declare(false, prototype, comps...)
}

// Redeclare 重声明实体原型
func (lib *_EntityLib) Redeclare(prototype any, comps ...any) ec.EntityPT {
	return lib.declare(true, prototype, comps...)
}

// Undeclare 取消声明实体原型
func (lib *_EntityLib) Undeclare(prototype string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.entityIdx, prototype)

	lib.entityList = slices.DeleteFunc(lib.entityList, func(pt *_Entity) bool {
		return pt.prototype == prototype
	})
}

// Get 获取实体原型
func (lib *_EntityLib) Get(prototype string) (ec.EntityPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	entity, ok := lib.entityIdx[prototype]
	if !ok {
		return nil, false
	}

	return entity, ok
}

// Range 遍历所有已注册的实体原型
func (lib *_EntityLib) Range(fun generic.Func1[ec.EntityPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.entityList)
	lib.RUnlock()

	for i := range copied {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

// ReversedRange 反向遍历所有已注册的实体原型
func (lib *_EntityLib) ReversedRange(fun generic.Func1[ec.EntityPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.entityList)
	lib.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

func (lib *_EntityLib) declare(re bool, prototype any, comps ...any) ec.EntityPT {
	if prototype == nil {
		exception.Panicf("%w: %w: prototype is nil", ErrPt, exception.ErrArgs)
	}

	if pie.Contains(comps, nil) {
		exception.Panicf("%w: %w: comps contains nil", ErrPt, exception.ErrArgs)
	}

	lib.Lock()
	defer lib.Unlock()

	var entityAtti EntityAttribute

	switch v := prototype.(type) {
	case EntityAttribute:
		entityAtti = v
	case *EntityAttribute:
		entityAtti = *v
	case string:
		entityAtti = EntityAttribute{Prototype: v}
	default:
		exception.Panicf("%w: invalid prototype type: %T", ErrPt, prototype)
	}

	if entityAtti.Prototype == "" {
		exception.Panicf("%w: prototype can't empty", ErrPt)
	}

	entityPT := &_Entity{
		prototype:                  entityAtti.Prototype,
		scope:                      entityAtti.Scope,
		componentAwakeOnFirstTouch: entityAtti.ComponentAwakeOnFirstTouch,
		componentUniqueID:          entityAtti.ComponentUniqueID,
		extra:                      entityAtti.Extra,
	}

	if entityAtti.Instance != nil {
		instanceRT, ok := entityAtti.Instance.(reflect.Type)
		if !ok {
			instanceRT = reflect.TypeOf(entityAtti.Instance)
		}

		for instanceRT.Kind() == reflect.Pointer || instanceRT.Kind() == reflect.Interface {
			instanceRT = instanceRT.Elem()
		}

		if instanceRT.Name() == "" {
			exception.Panicf("%w: anonymous entity instance not allowed", ErrPt)
		}

		if !reflect.PointerTo(instanceRT).Implements(reflect.TypeFor[ec.Entity]()) {
			exception.Panicf("%w: entity instance %q not implement ec.Entity", ErrPt, types.FullNameRT(instanceRT))
		}

		entityPT.instanceRT = instanceRT
	}

	for i, comp := range comps {
		compDesc := ec.ComponentDesc{
			Offset:       i,
			NonRemovable: true,
		}

	retry:
		switch v := comp.(type) {
		case ComponentAttribute:
			compDesc.Name = v.Name
			compDesc.NonRemovable = v.NonRemovable
			compDesc.Extra = v.Extra
			comp = v.Instance
			goto retry
		case *ComponentAttribute:
			compDesc.Name = v.Name
			compDesc.NonRemovable = v.NonRemovable
			compDesc.Extra = v.Extra
			comp = v.Instance
			goto retry
		case string:
			compPT, ok := lib.compLib.Get(v)
			if !ok {
				exception.Panicf("%w: entity %q component %q was not declared", ErrPt, prototype, v)
			}
			compDesc.PT = compPT
		default:
			compDesc.PT = lib.compLib.Declare(v)
		}

		if compDesc.Name == "" {
			compDesc.Name = compDesc.PT.Prototype()
		}

		entityPT.components = append(entityPT.components, compDesc)
	}

	if _, ok := lib.entityIdx[entityAtti.Prototype]; ok {
		if re {
			lib.entityList = slices.DeleteFunc(lib.entityList, func(pt *_Entity) bool {
				return pt.prototype == prototype
			})
		} else {
			exception.Panicf("%w: entity %q is already declared", ErrPt, prototype)
		}
	}

	lib.entityIdx[entityAtti.Prototype] = entityPT
	lib.entityList = append(lib.entityList, entityPT)

	return entityPT
}
