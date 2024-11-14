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

package define

import (
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/utils/types"
	"github.com/elliotchance/pie/v2"
)

// Component 定义组件
func Component[COMP any](name ...string) ComponentDefinition {
	return ComponentDefinition{
		Prototype: pt.DefaultComponentLib().Declare(types.ZeroT[COMP]()).Prototype(),
		Name:      pie.First(name),
	}
}

// ComponentDefinition 组件定义
type ComponentDefinition struct {
	Prototype string // 组件原型名称
	Name      string // 组件名称
}
