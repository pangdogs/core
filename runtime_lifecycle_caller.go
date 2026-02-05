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

func newEntityLifecycleCaller(entity ec.Entity) _EntityLifecycleCaller {
	return _EntityLifecycleCaller{entity: entity, state: entity.State()}
}

type _EntityLifecycleCaller struct {
	entity ec.Entity
	state  ec.EntityState
}

func (c _EntityLifecycleCaller) Call(fun func()) bool {
	if c.entity.State() != c.state {
		return false
	}

	fun()

	return c.entity.State() == c.state
}

func (c _EntityLifecycleCaller) IsProcessed(state ec.EntityState) bool {
	return ec.UnsafeEntity(c.entity).ProcessedStateBits().Is(int(state))
}

func (c _EntityLifecycleCaller) SetProcessed(state ec.EntityState) bool {
	bits := ec.UnsafeEntity(c.entity).ProcessedStateBits()
	if bits.Is(int(state)) {
		return false
	}
	bits.Set(int(state), true)
	return true
}

func (c _EntityLifecycleCaller) MarkProcessed() bool {
	return c.SetProcessed(c.state)
}

func newComponentLifecycleCaller(comp ec.Component) _ComponentLifecycleCaller {
	return _ComponentLifecycleCaller{component: comp, state: comp.State()}
}

type _ComponentLifecycleCaller struct {
	component ec.Component
	state     ec.ComponentState
}

func (c _ComponentLifecycleCaller) Call(fun func()) bool {
	if c.component.State() != c.state {
		return false
	}

	fun()

	return c.component.State() == c.state
}

func (c _ComponentLifecycleCaller) IsProcessed(state ec.ComponentState) bool {
	return ec.UnsafeComponent(c.component).ProcessedStateBits().Is(int(state))
}

func (c _ComponentLifecycleCaller) SetProcessed(state ec.ComponentState) bool {
	bits := ec.UnsafeComponent(c.component).ProcessedStateBits()
	if bits.Is(int(state)) {
		return false
	}
	bits.Set(int(state), true)
	return true
}

func (c _ComponentLifecycleCaller) MarkProcessed() bool {
	return c.SetProcessed(c.state)
}
