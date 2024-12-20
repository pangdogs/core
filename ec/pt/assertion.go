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
	"git.golaxy.org/core/utils/types"
	"reflect"
	"strings"
)

// As 从实体提取一些需要的组件接口，复合在一起直接使用。
/*
示例：
	type IA interface {
		MethodA()
	}
	...
	type IB interface {
		MethodB()
	}
	...
	type CompositeAB struct {
		IA `ec:"A"`			// 指定组件名
		IB `ec:",CompB"`	// 指定组件原型名
	}
	...
	entity.AddComponent("A", compA)
	entity.AddComponent("B", compB)
	...
	v, ok := As[CompositeAB](entity)
	if ok {
		v.MethodA()
		v.MethodB()
	}

注意：
	1.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
	2.实体更新组件后，需要重新提取。
	3.需要使用tag标记组件名或组件原型名，若没有标记，将会尝试使用字段名作为组件名去查找。
*/
func As[T comparable](entity ec.Entity) (T, bool) {
	iface := types.ZeroT[T]()
	ifaceRV := reflect.ValueOf(&iface).Elem()

	if !as(entity, ifaceRV) {
		return types.ZeroT[T](), false
	}

	return iface, true
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic。
/*
示例：
	type IA interface {
		MethodA()
	}
	...
	type IB interface {
		MethodB()
	}
	...
	type CompositeAB struct {
		IA `ec:"A"`			// 指定组件名
		IB `ec:",CompB"`	// 指定组件原型名
	}
	...
	entity.AddComponent("A", compA)
	entity.AddComponent("B", compB)
	...
	Cast[CompositeAB](entity).MethodA()
	Cast[CompositeAB](entity).MethodB()

注意：
	1.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
	2.实体更新组件后，需要重新提取。
	3.需要使用tag标记组件名或组件原型名，若没有标记，将会尝试使用字段名作为组件名去查找。
*/
func Cast[T comparable](entity ec.Entity) T {
	iface, ok := As[T](entity)
	if !ok {
		exception.Panicf("%w: incorrect cast", ErrPt)
	}
	return iface
}

// Compose 创建组件复合器
func Compose[T comparable](entity ec.Entity) *Composite[T] {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrPt, exception.ErrArgs)
	}
	return &Composite[T]{
		entity: entity,
	}
}

// Composite 组件复合器，直接使用As()或Cast()时，无法检测提取后实体是否又更新组件，使用复合器可以解决此问题。
/*
示例：
	type IA interface {
		MethodA()
	}
	...
	type IB interface {
		MethodB()
	}
	...
	type CompositeAB struct {
		IA `ec:"A"`			// 指定组件名
		IB `ec:",CompB"`	// 指定组件原型名
	}
	...
	entity.AddComponent("A", compA)
	entity.AddComponent("B", compB)
	...
	cx := Compose[CompositeAB](entity)
	if v, ok := cx.As(); ok {
		v.MethodA()
		v.MethodB()
	}
	cx.Cast().MethodA()
	cx.Cast().MethodB()
*/
type Composite[T comparable] struct {
	entity  ec.Entity
	version int64
	iface   T
}

// Entity 实体
func (c *Composite[T]) Entity() ec.Entity {
	if c.entity == nil {
		exception.Panicf("%w: entity is nil", ErrPt)
	}
	return c.entity
}

// Changed 实体是否已更新组件
func (c *Composite[T]) Changed() bool {
	if c.entity == nil {
		exception.Panicf("%w: entity is nil", ErrPt)
	}
	return c.version != ec.UnsafeEntity(c.entity).GetVersion()
}

// As 从实体提取一些需要的组件接口，复合在一起直接使用（实体更新组件后，会自动重新提取）
func (c *Composite[T]) As() (T, bool) {
	if c.entity == nil {
		exception.Panicf("%w: entity is nil", ErrPt)
	}

	if c.iface != types.ZeroT[T]() && !c.Changed() {
		return c.iface, true
	}

	if !as(c.entity, reflect.ValueOf(c.iface)) {
		return types.ZeroT[T](), false
	}

	c.version = ec.UnsafeEntity(c.entity).GetVersion()

	return c.iface, true
}

// Cast 从实体提取一些需要的组件接口，复合在一起直接使用，提取失败会panic（实体更新组件后，会自动重新提取）
func (c *Composite[T]) Cast() T {
	iface, ok := c.As()
	if !ok {
		exception.Panicf("%w: incorrect cast", ErrPt)
	}
	return iface
}

func as(entity ec.Entity, ifaceRV reflect.Value) bool {
	if entity == nil {
		return false
	}

	ifaceRT := ifaceRV.Type()

	switch ifaceRV.Kind() {
	case reflect.Struct:
		for i := 0; i < ifaceRV.NumField(); i++ {
			field := ifaceRT.Field(i)
			fieldRV := ifaceRV.Field(i)

			tag := strings.TrimSpace(field.Tag.Get("ec"))
			if tag == "-" {
				continue
			}

			name, prototype, _ := strings.Cut(tag, ",")
			name = strings.TrimSpace(name)
			prototype = strings.TrimSpace(prototype)

			if name != "" {
				comp := entity.GetComponent(name)
				if comp == nil {
					return false
				}
				fieldRV.Set(comp.GetReflected())
				continue
			}

			if prototype != "" {
				comp := entity.GetComponentByPT(prototype)
				if comp == nil {
					return false
				}
				fieldRV.Set(comp.GetReflected())
				continue
			}

			comp := entity.GetComponent(field.Name)
			if comp == nil {
				return false
			}
			fieldRV.Set(comp.GetReflected())
		}

		return true

	default:
		return false
	}
}
