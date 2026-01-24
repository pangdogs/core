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

package assertion

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"git.golaxy.org/core/service"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

// As 从实体中提取一些需要的组件，复合在一起直接使用。
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
	1.类型参数必须为结构体类型，可以是匿名结构体。
	2.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
	3.实体更新组件后，需要重新提取。
	4.需要使用tag标记组件名或组件原型名，若没有标记，将会尝试使用字段名作为组件名去查找。
	5.提取失败会返回false。
*/
func As[T any](entity ec.Entity) (*T, bool) {
	target := types.NewT[T]()

	if err := InjectRV(entity, reflect.ValueOf(target)); err != nil {
		return nil, false
	}

	return target, true
}

// Cast 从实体中提取一些需要的组件，复合在一起直接使用，提取失败会panic。
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
	1.类型参数必须为结构体类型，可以是匿名结构体。
	2.内部逻辑有使用反射，为了提高性能，可以使用一次后存储转换结果重复使用。
	3.实体更新组件后，需要重新提取。
	4.需要使用tag标记组件名或组件原型名，若没有标记，将会尝试使用字段名作为组件名去查找。
	5.提取失败会panic。
*/
func Cast[T any](entity ec.Entity) *T {
	target := types.NewT[T]()

	if err := InjectRV(entity, reflect.ValueOf(target)); err != nil {
		exception.Panicf("%w: incorrect cast, %w", exception.ErrCore, err)
	}

	return target
}

// Inject 向目标注入组件
func Inject(entity ec.Entity, target any) error {
	return InjectRV(entity, reflect.ValueOf(target))
}

// InjectRV 向目标注入组件
func InjectRV(entity ec.Entity, target reflect.Value) error {
	if entity == nil {
		return fmt.Errorf("%w: %w: entity is nil", exception.ErrCore, exception.ErrArgs)
	}

	targetRT := target.Type()

retry:
	switch target.Kind() {
	case reflect.Struct:
		for i := range target.NumField() {
			field := targetRT.Field(i)

			switch field.Type.Kind() {
			case reflect.Pointer:
				if field.Type.Elem().Kind() != reflect.Struct {
					continue
				}
			case reflect.Interface:
				break
			default:
				continue
			}

			tag := strings.TrimSpace(field.Tag.Get("ec"))
			if tag == "-" {
				continue
			}

			name, prototype, _ := strings.Cut(tag, ",")
			name = strings.TrimSpace(name)
			prototype = strings.TrimSpace(prototype)

			if name == "" && prototype == "" {
				switch field.Type.Kind() {
				case reflect.Pointer:
					fieldType := field.Type

					for fieldType.Kind() == reflect.Pointer {
						fieldType = fieldType.Elem()
					}

					name = types.NameRT(fieldType)
					prototype = types.FullNameRT(fieldType)
				default:
					continue
				}
			}

			if name != "" {
				comp := entity.GetComponent(name)
				if comp != nil && comp.GetReflected().Type().AssignableTo(field.Type) {
					if field.IsExported() {
						target.Field(i).Set(comp.GetReflected())
					} else {
						ptr := unsafe.Pointer(target.Field(i).UnsafeAddr())
						fieldPtr := reflect.NewAt(field.Type, ptr).Elem()
						fieldPtr.Set(comp.GetReflected())
					}
					continue
				}
			}

			if prototype != "" {
				comp := entity.GetComponentByPT(prototype)
				if comp != nil && comp.GetReflected().Type().AssignableTo(field.Type) {
					if field.IsExported() {
						target.Field(i).Set(comp.GetReflected())
					} else {
						ptr := unsafe.Pointer(target.Field(i).UnsafeAddr())
						fieldPtr := reflect.NewAt(field.Type, ptr).Elem()
						fieldPtr.Set(comp.GetReflected())
					}
					continue
				}

				compPT, ok := service.Current(entity).GetEntityLib().GetComponentLib().Get(prototype)
				if ok {
					sep := strings.LastIndexByte(prototype, '.')
					if sep >= 0 {
						comp := compPT.Construct()

						if err := entity.AddComponent(prototype[sep+1:], comp); err != nil {
							return fmt.Errorf("%w: %w", exception.ErrCore, err)
						}

						if field.IsExported() {
							target.Field(i).Set(comp.GetReflected())
						} else {
							ptr := unsafe.Pointer(target.Field(i).UnsafeAddr())
							fieldPtr := reflect.NewAt(field.Type, ptr).Elem()
							fieldPtr.Set(comp.GetReflected())
						}

						continue
					}
				}
			}
		}

		return nil

	case reflect.Pointer, reflect.Interface:
		if target.IsNil() {
			return fmt.Errorf("%w: target is nil", exception.ErrCore)
		}

		target = target.Elem()
		targetRT = target.Type()

		goto retry

	default:
		return fmt.Errorf("%w: invalid taget %s", exception.ErrCore, targetRT.Kind())
	}
}
