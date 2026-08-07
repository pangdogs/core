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
)

// EntityPTProvider 提供实体原型库。
type EntityPTProvider interface {
	// EntityLib 返回实体原型库。
	EntityLib() EntityLib
}

// For 从 provider 查询指定实体原型；provider 为 nil 或原型未声明时 panic。
func For(provider EntityPTProvider, prototype string) ec.EntityPT {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrPt, exception.ErrArgs)
	}

	entity, ok := provider.EntityLib().Get(prototype)
	if !ok {
		exception.Panicf("%w: entity %q was not declared", ErrPt, prototype)
	}

	return entity
}
