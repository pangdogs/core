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

// As 创建 T 并从 entity 向其字段注入匹配组件。
//
// T 必须是结构体。字段标签格式为 `ec:"组件名,完整组件原型名"`；指针字段未标注
// 标签时会按字段类型推导名称和原型。未找到组件的字段保持零值，不影响 ok；只有注入
// 参数或动态添加组件出错时 ok 才为 false。返回值反映调用时的组件快照，实体组件变化后
// 应重新提取。
func As[T any](entity ec.Entity) (*T, bool) {
	target := types.New[T]()

	if err := InjectRV(entity, reflect.ValueOf(target)); err != nil {
		return nil, false
	}

	return target, true
}

// Cast 与 As 相同地创建并注入 T，但在 InjectRV 返回错误时 panic。
// 未找到匹配组件的字段仍保持零值，不会单独触发 panic。
func Cast[T any](entity ec.Entity) *T {
	target := types.New[T]()

	if err := InjectRV(entity, reflect.ValueOf(target)); err != nil {
		exception.Panicf("%w: incorrect cast, %w", exception.ErrCore, err)
	}

	return target
}

// Inject 使用反射向非 nil 结构体指针 target 注入 entity 的匹配组件。
// 该操作可能按已声明原型创建并添加缺失组件。
func Inject(entity ec.Entity, target any) error {
	return InjectRV(entity, reflect.ValueOf(target))
}

// InjectRV 向可寻址结构体或可解引用到结构体的 target 注入匹配组件。
// 未找到匹配组件时保留字段原值；entity 为 nil、target 可解引用但为 nil、
// target 类型不受支持或动态添加失败时返回错误。无效的 reflect.Value 会在读取类型时 panic。
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
				if comp != nil && comp.Reflected().Type().AssignableTo(field.Type) {
					if field.IsExported() {
						target.Field(i).Set(comp.Reflected())
					} else {
						ptr := unsafe.Pointer(target.Field(i).UnsafeAddr())
						fieldPtr := reflect.NewAt(field.Type, ptr).Elem()
						fieldPtr.Set(comp.Reflected())
					}
					continue
				}
			}

			if prototype != "" {
				comp := entity.GetComponentByPT(prototype)
				if comp != nil && comp.Reflected().Type().AssignableTo(field.Type) {
					if field.IsExported() {
						target.Field(i).Set(comp.Reflected())
					} else {
						ptr := unsafe.Pointer(target.Field(i).UnsafeAddr())
						fieldPtr := reflect.NewAt(field.Type, ptr).Elem()
						fieldPtr.Set(comp.Reflected())
					}
					continue
				}

				compPT, ok := service.Current(entity).EntityLib().ComponentLib().Get(prototype)
				if ok {
					sep := strings.LastIndexByte(prototype, '.')
					if sep >= 0 {
						comp := compPT.Construct()

						if err := entity.AddComponent(prototype[sep+1:], comp); err != nil {
							return fmt.Errorf("%w: %w", exception.ErrCore, err)
						}

						if field.IsExported() {
							target.Field(i).Set(comp.Reflected())
						} else {
							ptr := unsafe.Pointer(target.Field(i).UnsafeAddr())
							fieldPtr := reflect.NewAt(field.Type, ptr).Elem()
							fieldPtr.Set(comp.Reflected())
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
