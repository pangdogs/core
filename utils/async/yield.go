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

	"git.golaxy.org/core/utils/exception"
)

// YieldReturn 尝试向 future 产出一项结果。
//
// nil ctx 按 context.Background 处理。写入成功返回 true；ctx 先结束时返回 false。
// future 为零值或非结果 Future 时 panic。该函数不会结束 Future。
func YieldReturn(ctx context.Context, future FutureChan, ret Result) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	if future.ch == nil || future.done == nil {
		exception.Panic("future is void result, cannot yield return")
	}

	select {
	case future.ch <- ret:
		return true
	case <-ctx.Done():
		return false
	}
}

// YieldBreak 结束产出，关闭结果与完成频道，并返回消费者视图。
// future 为零值或已经结束时 panic。
func YieldBreak(future FutureChan) Future {
	if future.ch == nil || future.done == nil {
		exception.Panic("future is void result, cannot yield break")
	}
	close(future.ch)
	close(future.done)
	return future.Out()
}
