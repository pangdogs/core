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

package ec

import (
	"encoding/json"
	"fmt"
	"reflect"

	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
)

// EntityPT 实体原型接口
type EntityPT interface {
	fmt.Stringer

	// Prototype 实体原型名称
	Prototype() string
	// InstanceRT 实体实例反射类型
	InstanceRT() reflect.Type
	// Scope 可访问作用域
	Scope() Scope
	// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentAwakeOnFirstTouch() bool
	// ComponentUniqueID 是否为实体组件分配唯一Id
	ComponentUniqueID() bool
	// Meta 原型Meta信息
	Meta() meta.Meta
	// CountComponents 组件数量
	CountComponents() int
	// GetComponent 获取组件
	GetComponent(idx int) BuiltinComponent
	// ListComponents 获取所有组件
	ListComponents() []BuiltinComponent
	// Construct 创建实体
	Construct(settings ...option.Setting[EntityOptions]) Entity
}

// BuiltinComponent 实体原型中的组件信息
type BuiltinComponent struct {
	PT        ComponentPT // 组件原型
	Offset    int         // 组件位置
	Name      string      // 组件名称
	Removable bool        // 可以删除
	Meta      meta.Meta   // 原型Meta信息
}

// String implements fmt.Stringer
func (bc BuiltinComponent) String() string {
	data, err := json.Marshal(bc)
	if err != nil {
		exception.Panicf("%w: unexpected failure marshaling builtin component: %s", ErrEC, err)
	}
	return string(data)
}

type _BuiltinComponentJSON struct {
	PT        ComponentPT    `json:"pt"`
	Offset    int            `json:"offset"`
	Name      string         `json:"name"`
	Removable bool           `json:"removable"`
	Meta      map[string]any `json:"meta"`
}

// MarshalJSON implements json.Marshaler
func (bc BuiltinComponent) MarshalJSON() ([]byte, error) {
	builtinComponentStringer := _BuiltinComponentJSON{
		PT:        bc.PT,
		Offset:    bc.Offset,
		Name:      bc.Name,
		Removable: bc.Removable,
		Meta:      bc.Meta.ToGoMap(),
	}

	data, err := json.Marshal(builtinComponentStringer)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ComponentPT 组件原型接口
type ComponentPT interface {
	fmt.Stringer

	// Prototype 组件原型名称
	Prototype() string
	// InstanceRT 组件实例反射类型
	InstanceRT() reflect.Type
	// Construct 创建组件
	Construct() Component
}

var (
	noneEntityPT         = &_NoneEntityPT{}
	noneComponentPT      = &_NoneComponentPT{}
	noneBuiltinComponent = &BuiltinComponent{PT: noneComponentPT, Offset: -1}
)

type _NoneEntityPT struct{}

// Prototype 实体原型名称
func (_NoneEntityPT) Prototype() string {
	return ""
}

// InstanceRT 实体实例反射类型
func (_NoneEntityPT) InstanceRT() reflect.Type {
	return nil
}

// Scope 可访问作用域
func (_NoneEntityPT) Scope() Scope {
	return Scope_Global
}

// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (_NoneEntityPT) ComponentAwakeOnFirstTouch() bool {
	return false
}

// ComponentUniqueID 是否为实体组件分配唯一Id
func (_NoneEntityPT) ComponentUniqueID() bool {
	return false
}

// Meta 原型Meta信息
func (_NoneEntityPT) Meta() meta.Meta {
	return nil
}

// CountComponents 组件数量
func (_NoneEntityPT) CountComponents() int {
	return 0
}

// GetComponent 获取组件
func (_NoneEntityPT) GetComponent(idx int) BuiltinComponent {
	exception.Panicf("%w: %w: idx out of range", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// ListComponents 获取所有组件
func (_NoneEntityPT) ListComponents() []BuiltinComponent {
	return nil
}

// Construct 创建实体
func (_NoneEntityPT) Construct(settings ...option.Setting[EntityOptions]) Entity {
	exception.Panicf("%w: %w: none prototype", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// String implements fmt.Stringer
func (_NoneEntityPT) String() string {
	return "null"
}

// MarshalJSON implements json.Marshaler
func (_NoneEntityPT) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

type _NoneComponentPT struct{}

// Prototype 组件原型名称
func (_NoneComponentPT) Prototype() string {
	return ""
}

// InstanceRT 组件实例反射类型
func (_NoneComponentPT) InstanceRT() reflect.Type {
	return nil
}

// Construct 创建组件
func (_NoneComponentPT) Construct() Component {
	exception.Panicf("%w: %w: none prototype", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// String implements fmt.Stringer
func (_NoneComponentPT) String() string {
	return "null"
}

// MarshalJSON implements json.Marshaler
func (_NoneComponentPT) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
