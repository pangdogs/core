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

package async

import (
	"context"
	"fmt"

	"git.golaxy.org/core/utils/exception"
)

var (
	ErrAsync             = fmt.Errorf("%w: async", exception.ErrCore)                     // ErrAsync 是异步模块错误的共同根错误。
	ErrScopeClosed       = fmt.Errorf("%w: scope closed: %w", ErrAsync, context.Canceled) // ErrScopeClosed 表示异步作用域已因取消而关闭。
	ErrNoCandidates      = fmt.Errorf("%w: no future candidates", ErrAsync)               // ErrNoCandidates 表示组合器没有可用的候选 Future。
	ErrNoFutureSucceeded = fmt.Errorf("%w: no future succeeded", ErrAsync)                // ErrNoFutureSucceeded 表示所有候选 Future 均失败。
	ErrFutureTimeout     = fmt.Errorf("%w: future timeout", ErrAsync)                     // ErrFutureTimeout 表示 Future 等待超时。
)
