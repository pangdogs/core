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

// EntityLib 是可并发访问的实体原型注册表。
type EntityLib interface {
	EntityPTProvider

	// ComponentLib 返回实体原型解析组件时使用的组件原型库。
	ComponentLib() ComponentLib
	// Declare 声明实体原型；同名声明会替换旧原型。
	Declare(prototype any, comps ...any) ec.EntityPT
	// Get 按原型名查询实体原型。
	Get(prototype string) (ec.EntityPT, bool)
	// List 返回当前全部实体原型的快照。
	List() []ec.EntityPT
	// Watch 依次发送当前快照及后续声明；ctx 结束后排空已排队项并关闭频道。
	Watch(ctx context.Context) <-chan ec.EntityPT
}

// NewEntityLib 使用 compLib 创建独立的空实体原型库；compLib 为 nil 时 panic。
func NewEntityLib(compLib ComponentLib) EntityLib {
	if compLib == nil {
		exception.Panicf("%w: %w: compLib is nil", ErrPt, exception.ErrArgs)
	}

	lib := &_EntityLib{compLib: compLib}
	lib.snapshot.Store(&_EntityLibSnapshot{
		entityPTIndex: map[string]ec.EntityPT{},
	})
	return lib
}

type _EntityLib struct {
	mutex       sync.Mutex
	compLib     ComponentLib
	snapshot    atomic.Pointer[_EntityLibSnapshot]
	eventStream generic.EventStream[ec.EntityPT]
}

// _EntityLibSnapshot 是实体原型库的只读快照，发布后不再修改。
type _EntityLibSnapshot struct {
	entityPTIndex map[string]ec.EntityPT
	entityPTList  []ec.EntityPT
}

func (snapshot *_EntityLibSnapshot) clone() *_EntityLibSnapshot {
	return &_EntityLibSnapshot{
		entityPTIndex: maps.Clone(snapshot.entityPTIndex),
		entityPTList:  slices.Clone(snapshot.entityPTList),
	}
}

// EntityLib 返回自身，以实现 EntityPTProvider。
func (lib *_EntityLib) EntityLib() EntityLib {
	return lib
}

// ComponentLib 返回实体原型解析组件时使用的组件原型库。
func (lib *_EntityLib) ComponentLib() ComponentLib {
	return lib.compLib
}

// Declare 声明实体原型；同名声明会替换旧原型并发布一次声明事件。
//
// prototype 支持原型名、EntityDescriptor 或其指针；comps 支持组件值、完整原型名、
// ComponentDescriptor 或其指针。参数无效或引用未声明的组件原型时 panic。
func (lib *_EntityLib) Declare(prototype any, comps ...any) ec.EntityPT {
	if prototype == nil {
		exception.Panicf("%w: %w: prototype is nil", ErrPt, exception.ErrArgs)
	}

	if slices.Contains(comps, nil) {
		exception.Panicf("%w: %w: comps contains nil", ErrPt, exception.ErrArgs)
	}

	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	var entityDescr EntityDescriptor

	switch v := prototype.(type) {
	case EntityDescriptor:
		entityDescr = v
	case *EntityDescriptor:
		entityDescr = *v
	case string:
		entityDescr = EntityDescriptor{Prototype: v}
	default:
		exception.Panicf("%w: invalid prototype type: %T", ErrPt, prototype)
	}

	if entityDescr.Prototype == "" {
		exception.Panicf("%w: prototype can't empty", ErrPt)
	}

	entityPT := &_Entity{
		prototype:                  entityDescr.Prototype,
		scope:                      entityDescr.Scope,
		componentAwakeOnFirstTouch: entityDescr.ComponentAwakeOnFirstTouch,
		componentUniqueID:          entityDescr.ComponentUniqueID,
		meta:                       entityDescr.Meta,
	}

	if entityDescr.Instance != nil {
		instanceRT, ok := entityDescr.Instance.(reflect.Type)
		if !ok {
			instanceRT = reflect.TypeOf(entityDescr.Instance)
		}

		for instanceRT.Kind() == reflect.Pointer {
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
		builtin := ec.BuiltinComponent{
			Offset: i,
		}

	retry:
		switch v := comp.(type) {
		case ComponentDescriptor:
			builtin.Name = v.Name
			builtin.Removable = v.Removable
			builtin.Meta = v.Meta
			comp = v.Instance
			goto retry
		case *ComponentDescriptor:
			comp = *v
			goto retry
		case string:
			compPT, ok := lib.compLib.Get(v)
			if !ok {
				exception.Panicf("%w: entity %q builtin component %q was not declared", ErrPt, prototype, v)
			}
			builtin.PT = compPT
		default:
			if v == nil {
				exception.Panicf("%w: entity %q builtin component is nil", ErrPt, prototype)
			}
			builtin.PT = lib.compLib.Declare(v)
		}

		if builtin.Name == "" {
			builtin.Name = types.NameRT(builtin.PT.InstanceRT().Elem())
		}

		entityPT.components = append(entityPT.components, builtin)
	}

	snapshot := lib.snapshot.Load()
	next := snapshot.clone()

	if _, ok := next.entityPTIndex[entityDescr.Prototype]; ok {
		next.entityPTList = slices.DeleteFunc(next.entityPTList, func(entityPT ec.EntityPT) bool {
			return entityPT.Prototype() == entityDescr.Prototype
		})
	}

	next.entityPTIndex[entityDescr.Prototype] = entityPT
	next.entityPTList = append(next.entityPTList, entityPT)
	lib.snapshot.Store(next)

	lib.eventStream.Publish(entityPT)

	return entityPT
}

// Get 按原型名查询实体原型。
func (lib *_EntityLib) Get(prototype string) (ec.EntityPT, bool) {
	entityPT, ok := lib.snapshot.Load().entityPTIndex[prototype]
	return entityPT, ok
}

// List 返回当前全部实体原型的快照。
func (lib *_EntityLib) List() []ec.EntityPT {
	return slices.Clone(lib.snapshot.Load().entityPTList)
}

// Watch 依次发送当前快照及后续声明；ctx 结束后排空已排队项并关闭频道。
// nil ctx 按 context.Background 处理。
func (lib *_EntityLib) Watch(ctx context.Context) <-chan ec.EntityPT {
	if ctx == nil {
		ctx = context.Background()
	}

	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	return lib.eventStream.Subscribe(ctx, lib.snapshot.Load().entityPTList...)
}
