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
	"git.golaxy.org/core/utils/types"
	"github.com/elliotchance/pie/v2"
)

// Component 定义组件
func Component[COMP any](compLib ...pt.ComponentLib) ComponentDefinition {
	return defineComponent[COMP](getCompLib(compLib...), "")
}

// ComponentWithInterface 定义有接口的组件，接口名称将作为组件名
func ComponentWithInterface[COMP, COMP_IFACE any](compLib ...pt.ComponentLib) ComponentDefinition {
	return defineComponent[COMP](getCompLib(compLib...), ComponentInterface[COMP_IFACE](getCompLib(compLib...)).Name)
}

// ComponentDefinition 组件定义
type ComponentDefinition struct {
	Name          string // 组件名称
	InterfaceName string // 组件接口名称
}

func defineComponent[COMP any](compLib pt.ComponentLib, ifaceName string) ComponentDefinition {
	if compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	return ComponentDefinition{
		Name:          compLib.Declare(types.ZeroT[COMP]()).Name,
		InterfaceName: ifaceName,
	}
}

func getCompLib(compLib ...pt.ComponentLib) pt.ComponentLib {
	if l := pie.First(compLib); l != nil {
		return l
	}
	return pt.DefaultComponentLib()
}
