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
	"sync"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
)

type _Component struct {
	prototype     string
	instanceRT    reflect.Type
	builtin       *ec.BuiltinComponent
	stringerOnce  sync.Once
	stringerCache string
}

// Prototype 组件原型名称
func (pt *_Component) Prototype() string {
	return pt.prototype
}

// InstanceRT 组件实例反射类型
func (pt *_Component) InstanceRT() reflect.Type {
	return reflect.PointerTo(pt.instanceRT)
}

// Construct 创建组件
func (pt *_Component) Construct() ec.Component {
	compRV := reflect.New(pt.instanceRT)

	comp := compRV.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetBuiltin(pt.builtin)
	ec.UnsafeComponent(comp).SetReflected(compRV)

	return comp
}

// String implements fmt.Stringer
func (pt *_Component) String() string {
	pt.stringerOnce.Do(func() {
		data, err := json.Marshal(pt)
		if err != nil {
			exception.Panicf("%w: unexpected failure marshaling component: %s", ErrPt, err)
		}
		pt.stringerCache = string(data)
	})
	return pt.stringerCache
}

type _ComponentJSON struct {
	Prototype string `json:"prototype"`
	Instance  string `json:"instance"`
}

// MarshalJSON implements json.Marshaler
func (pt _Component) MarshalJSON() ([]byte, error) {
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
