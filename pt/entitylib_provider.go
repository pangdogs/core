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
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

// EntityPTProvider 实体原型提供者
type EntityPTProvider interface {
	// GetEntityLib 获取实体原型库
	GetEntityLib() EntityLib
}

// For 查询实体原型
func For(provider EntityPTProvider, prototype string) EntityPT {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPt, exception.ErrArgs))
	}

	entity, ok := provider.GetEntityLib().Get(prototype)
	if !ok {
		panic(fmt.Errorf("%w: entity %q was not declared", ErrPt, prototype))
	}

	return entity
}
