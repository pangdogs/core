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
	"slices"
	"sync"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// EntityLib 实体原型库
type EntityLib interface {
	EntityPTProvider

	// GetComponentLib 获取组件原型库
	GetComponentLib() ComponentLib
	// Declare 声明实体原型
	Declare(prototype any, comps ...any) ec.EntityPT
	// Get 获取实体原型
	Get(prototype string) (ec.EntityPT, bool)
	// List 获取所有实体原型
	List() []ec.EntityPT
	// EventStream 组件声明事件流
	EventStream(ctx context.Context) <-chan ec.EntityPT
}

// NewEntityLib 创建实体原型库
func NewEntityLib(compLib ComponentLib) EntityLib {
	if compLib == nil {
		exception.Panicf("%w: %w: compLib is nil", ErrPt, exception.ErrArgs)
	}

	return &_EntityLib{
		compLib:       compLib,
		entityPTIndex: map[string]int{},
	}
}

type _EntityLib struct {
	sync.RWMutex
	compLib       ComponentLib
	entityPTIndex map[string]int
	entityPTList  generic.FreeList[ec.EntityPT]
	eventStream   generic.EventStream[ec.EntityPT]
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
	if prototype == nil {
		exception.Panicf("%w: %w: prototype is nil", ErrPt, exception.ErrArgs)
	}

	if slices.Contains(comps, nil) {
		exception.Panicf("%w: %w: comps contains nil", ErrPt, exception.ErrArgs)
	}

	lib.Lock()
	defer lib.Unlock()

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

	if entityPTIdx, ok := lib.entityPTIndex[entityDescr.Prototype]; ok {
		lib.entityPTList.Release(entityPTIdx)
	}

	lib.entityPTIndex[entityDescr.Prototype] = lib.entityPTList.PushBack(entityPT).Index()

	lib.eventStream.Publish(entityPT)

	return entityPT
}

// Get 获取实体原型
func (lib *_EntityLib) Get(prototype string) (ec.EntityPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	entityPTIdx, ok := lib.entityPTIndex[prototype]
	if !ok {
		return nil, false
	}

	return lib.entityPTList.Get(entityPTIdx).V, true
}

// List 获取所有实体原型
func (lib *_EntityLib) List() []ec.EntityPT {
	lib.RLock()
	defer lib.RUnlock()

	return lib.entityPTList.ToSlice()
}

// EventStream 组件声明事件流
func (lib *_EntityLib) EventStream(ctx context.Context) <-chan ec.EntityPT {
	if ctx == nil {
		ctx = context.Background()
	}

	lib.Lock()
	defer lib.Unlock()

	return lib.eventStream.Subscribe(ctx, lib.entityPTList.ToSlice()...)
}
