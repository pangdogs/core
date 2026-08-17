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

// EntityPT 描述可构造实体及其内建组件的原型。
type EntityPT interface {
	fmt.Stringer

	// Prototype 返回实体原型名。
	Prototype() string
	// InstanceRT 返回实际实体实例的指针类型；使用默认实体实现时返回 nil。
	InstanceRT() reflect.Type
	// Scope 返回原型的默认实体作用域。
	Scope() Scope
	// ComponentAwakeOnFirstTouch 报告正常激活期间被访问的组件是否优先执行 Awake。
	ComponentAwakeOnFirstTouch() bool
	// ComponentUniqueID 报告是否为组件分配唯一 ID。
	ComponentUniqueID() bool
	// Meta 返回原型元数据。
	Meta() meta.Meta
	// CountComponents 返回内建组件数。
	CountComponents() int
	// GetComponent 返回指定位置的内建组件描述；索引越界时 panic。
	GetComponent(idx int) BuiltinComponent
	// ListComponents 返回全部内建组件描述的副本。
	ListComponents() []BuiltinComponent
	// Construct 根据原型创建处于 Born 状态的实体，并应用额外选项。
	Construct(settings ...option.Setting[EntityOptions]) Entity
}

// BuiltinComponent 描述实体原型中的一个内建组件。
type BuiltinComponent struct {
	PT        ComponentPT // PT 是组件原型。
	Offset    int         // Offset 是组件在实体原型中的位置。
	Name      string      // Name 是组件加入实体时使用的名称。
	Removable bool        // Removable 指示组件是否允许动态删除。
	Meta      meta.Meta   // Meta 是该内建组件的原型元数据。
}

// String 返回内建组件描述的 JSON 文本；编码失败时 panic。
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

// MarshalJSON 将内建组件描述编码为 JSON。
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

// ComponentPT 描述可构造组件的原型。
type ComponentPT interface {
	fmt.Stringer

	// Prototype 返回组件原型名。
	Prototype() string
	// InstanceRT 返回实际组件实例的指针类型。
	InstanceRT() reflect.Type
	// Construct 根据原型创建处于 Born 状态的组件。
	Construct() Component
}

var (
	noneEntityPT         = &_NoneEntityPT{}
	noneComponentPT      = &_NoneComponentPT{}
	noneBuiltinComponent = &BuiltinComponent{PT: noneComponentPT, Removable: true, Offset: -1}
)

type _NoneEntityPT struct{}

// Prototype 返回空实体原型的名称。
func (_NoneEntityPT) Prototype() string {
	return ""
}

// InstanceRT 返回 nil，表示空实体原型没有实例类型。
func (_NoneEntityPT) InstanceRT() reflect.Type {
	return nil
}

// Scope 返回空实体原型的默认全局作用域。
func (_NoneEntityPT) Scope() Scope {
	return Scope_Global
}

// ComponentAwakeOnFirstTouch 对空实体原型返回 false。
func (_NoneEntityPT) ComponentAwakeOnFirstTouch() bool {
	return false
}

// ComponentUniqueID 对空实体原型返回 false。
func (_NoneEntityPT) ComponentUniqueID() bool {
	return false
}

// Meta 对空实体原型返回 nil。
func (_NoneEntityPT) Meta() meta.Meta {
	return nil
}

// CountComponents 对空实体原型返回 0。
func (_NoneEntityPT) CountComponents() int {
	return 0
}

// GetComponent 对空实体原型始终 panic。
func (_NoneEntityPT) GetComponent(idx int) BuiltinComponent {
	exception.Panicf("%w: %w: idx out of range", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// ListComponents 对空实体原型返回 nil。
func (_NoneEntityPT) ListComponents() []BuiltinComponent {
	return nil
}

// Construct 对空实体原型始终 panic。
func (_NoneEntityPT) Construct(settings ...option.Setting[EntityOptions]) Entity {
	exception.Panicf("%w: %w: none prototype", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// String 返回 JSON 空值文本。
func (_NoneEntityPT) String() string {
	return "null"
}

// MarshalJSON 返回 JSON 空值。
func (_NoneEntityPT) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

type _NoneComponentPT struct{}

// Prototype 返回空组件原型的名称。
func (_NoneComponentPT) Prototype() string {
	return ""
}

// InstanceRT 返回 nil，表示空组件原型没有实例类型。
func (_NoneComponentPT) InstanceRT() reflect.Type {
	return nil
}

// Construct 对空组件原型始终 panic。
func (_NoneComponentPT) Construct() Component {
	exception.Panicf("%w: %w: none prototype", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// String 返回 JSON 空值文本。
func (_NoneComponentPT) String() string {
	return "null"
}

// MarshalJSON 返回 JSON 空值。
func (_NoneComponentPT) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}
