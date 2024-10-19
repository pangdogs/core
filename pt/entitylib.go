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

	// Declare 声明实体原型
	Declare(prototype any, comps ...any) EntityPT
	// Redeclare 重新声明实体原型
	Redeclare(prototype any, comps ...any) EntityPT
	// Undeclare 取消声明实体原型
	Undeclare(prototype string)
	// Get 获取实体原型
	Get(prototype string) (EntityPT, bool)
	// Range 遍历所有已注册的实体原型
	Range(fun generic.Func1[EntityPT, bool])
	// ReversedRange 反向遍历所有已注册的实体原型
	ReversedRange(fun generic.Func1[EntityPT, bool])
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
		entityIdx: map[string]*_EntityPT{},
	}
}

type _EntityLib struct {
	sync.RWMutex
	compLib    ComponentLib
	entityIdx  map[string]*_EntityPT
	entityList []*_EntityPT
}

// GetEntityLib 获取实体原型库
func (lib *_EntityLib) GetEntityLib() EntityLib {
	return lib
}

// Declare 声明实体原型
func (lib *_EntityLib) Declare(prototype any, comps ...any) EntityPT {
	return lib.declare(false, prototype, comps...)
}

// Redeclare 重新声明实体原型
func (lib *_EntityLib) Redeclare(prototype any, comps ...any) EntityPT {
	return lib.declare(true, prototype, comps...)
}

// Undeclare 取消声明实体原型
func (lib *_EntityLib) Undeclare(prototype string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.entityIdx, prototype)

	lib.entityList = slices.DeleteFunc(lib.entityList, func(pt *_EntityPT) bool {
		return pt.prototype == prototype
	})
}

// Get 获取实体原型
func (lib *_EntityLib) Get(prototype string) (EntityPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	entity, ok := lib.entityIdx[prototype]
	if !ok {
		return nil, false
	}

	return entity, ok
}

// Range 遍历所有已注册的实体原型
func (lib *_EntityLib) Range(fun generic.Func1[EntityPT, bool]) {
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
func (lib *_EntityLib) ReversedRange(fun generic.Func1[EntityPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.entityList)
	lib.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

func (lib *_EntityLib) declare(re bool, prototype any, comps ...any) EntityPT {
	if prototype == nil {
		exception.Panicf("%w: %w: prototype is nil", ErrPt, exception.ErrArgs)
	}

	if pie.Contains(comps, nil) {
		exception.Panicf("%w: %w: comps contains nil", ErrPt, exception.ErrArgs)
	}

	lib.Lock()
	defer lib.Unlock()

	var entityAtti EntityAtti

	switch v := prototype.(type) {
	case EntityAtti:
		entityAtti = v
	case *EntityAtti:
		entityAtti = *v
	case string:
		entityAtti = EntityAtti{Prototype: v}
	default:
		exception.Panicf("%w: invalid prototype type: %T", ErrPt, prototype)
	}

	if entityAtti.Prototype == "" {
		exception.Panicf("%w: prototype can't empty", ErrPt)
	}

	entityPT := &_EntityPT{
		prototype:          entityAtti.Prototype,
		scope:              entityAtti.Scope,
		awakeOnFirstAccess: entityAtti.AwakeOnFirstAccess,
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

	for _, comp := range comps {
		compDesc := ComponentDesc{NonRemovable: true}

	retry:
		switch v := comp.(type) {
		case CompAtti:
			compDesc.Name = v.Name
			compDesc.NonRemovable = v.NonRemovable
			comp = v.Instance
			goto retry
		case *CompAtti:
			compDesc.Name = v.Name
			compDesc.NonRemovable = v.NonRemovable
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
			lib.entityList = slices.DeleteFunc(lib.entityList, func(pt *_EntityPT) bool {
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
