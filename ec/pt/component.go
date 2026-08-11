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
	"encoding/json"
	"reflect"
	"sync/atomic"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
)

type _Component struct {
	prototype     string
	instanceRT    reflect.Type
	builtin       *ec.BuiltinComponent
	stringerCache atomic.Pointer[string]
}

// Prototype 返回组件的完整原型名。
func (pt *_Component) Prototype() string {
	return pt.prototype
}

// InstanceRT 返回组件实例的指针类型。
func (pt *_Component) InstanceRT() reflect.Type {
	return reflect.PointerTo(pt.instanceRT)
}

// Construct 创建处于 Born 状态的组件，并绑定其组件原型。
func (pt *_Component) Construct() ec.Component {
	compRV := reflect.New(pt.instanceRT)

	comp := compRV.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetBuiltin(pt.builtin)
	ec.UnsafeComponent(comp).SetReflected(compRV)

	return comp
}

// String 返回组件原型的 JSON 文本；编码失败时 panic。
func (pt *_Component) String() string {
	if cached := pt.stringerCache.Load(); cached != nil {
		return *cached
	}

	data, err := json.Marshal(pt)
	if err != nil {
		exception.Panicf("%w: unexpected failure marshaling component: %s", ErrPt, err)
	}
	value := string(data)
	if pt.stringerCache.CompareAndSwap(nil, &value) {
		return value
	}
	return *pt.stringerCache.Load()
}

type _ComponentJSON struct {
	Prototype string `json:"prototype"`
	Instance  string `json:"instance"`
}

// MarshalJSON 将组件原型编码为 JSON。
func (pt *_Component) MarshalJSON() ([]byte, error) {
	compStringer := _ComponentJSON{
		Prototype: pt.prototype,
		Instance:  pt.instanceRT.String(),
	}

	data, err := json.Marshal(compStringer)
	if err != nil {
		return nil, err
	}

	return data, nil
}
