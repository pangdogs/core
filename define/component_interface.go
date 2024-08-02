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
	"fmt"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// ComponentInterface 定义组件接口
func ComponentInterface[COMP_IFACE any](compLib ...pt.ComponentLib) ComponentInterfaceDefinition {
	return defineComponentInterface[COMP_IFACE](getCompLib(compLib...))
}

// ComponentInterfaceDefinition 组件接口定义
type ComponentInterfaceDefinition struct {
	Name    string                                         // 组件接口名称
	Declare generic.PairFunc1[any, pt.ComponentPT, string] // 声明组件原型
}

func defineComponentInterface[COMP_IFACE any](compLib pt.ComponentLib) ComponentInterfaceDefinition {
	if compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	name := types.FullNameT[COMP_IFACE]()

	return ComponentInterfaceDefinition{
		Name: name,
		Declare: func(comp any) (pt.ComponentPT, string) {
			return compLib.Declare(comp, name), name
		},
	}
}
