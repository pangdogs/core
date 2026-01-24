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

package core

import "git.golaxy.org/core/ec"

func makeEntityLifecycleCaller(entity ec.Entity) _EntityLifecycleCaller {
	return _EntityLifecycleCaller{entity: entity, state: entity.GetState()}
}

type _EntityLifecycleCaller struct {
	entity ec.Entity
	state  ec.EntityState
}

func (c _EntityLifecycleCaller) CallOnce(fun func(state ec.EntityState)) bool {
	state := c.entity.GetState()
	if state != c.state {
		return false
	}

	bits := ec.UnsafeEntity(c.entity).GetCallingStateBits()
	if bits.Is(int(state)) {
		return true
	}

	bits.Set(int(state), true)
	defer bits.Set(int(state), false)

	fun(c.state)

	return c.entity.GetState() == c.state
}

func (c _EntityLifecycleCaller) Call(fun func(state ec.EntityState)) bool {
	if c.entity.GetState() != c.state {
		return false
	}

	fun(c.state)

	return c.entity.GetState() == c.state
}

func makeComponentLifecycleCaller(comp ec.Component) _ComponentLifecycleCaller {
	return _ComponentLifecycleCaller{component: comp, state: comp.GetState()}
}

type _ComponentLifecycleCaller struct {
	component ec.Component
	state     ec.ComponentState
}

func (c _ComponentLifecycleCaller) CallOnce(fun func(state ec.ComponentState)) bool {
	state := c.component.GetState()
	if state != c.state {
		return false
	}

	bits := ec.UnsafeComponent(c.component).GetCallingStateBits()
	if bits.Is(int(state)) {
		return true
	}

	bits.Set(int(state), true)
	defer bits.Set(int(state), false)

	fun(c.state)

	return c.component.GetState() == c.state
}

func (c _ComponentLifecycleCaller) Call(fun func(state ec.ComponentState)) bool {
	if c.component.GetState() != c.state {
		return false
	}

	fun(c.state)

	return c.component.GetState() == c.state
}
