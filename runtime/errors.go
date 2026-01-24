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

package runtime

import (
	"fmt"

	"git.golaxy.org/core/utils/exception"
)

var (
	ErrContext       = fmt.Errorf("%w: runtime-context", exception.ErrCore) // 运行时上下文错误
	ErrEntityTree    = fmt.Errorf("%w: entity-tree", ErrContext)            // 实体树错误
	ErrEntityManager = fmt.Errorf("%w: entity-manager", ErrContext)         // 实体管理器错误
	ErrFrame         = fmt.Errorf("%w: frame", ErrContext)                  // 帧错误
)
