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

func YieldReturn(ctx context.Context, future FutureStream, ret Result) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	if cap(future) <= 0 {
		exception.Panic("future is void result, cannot yield return")
	}

	select {
	case future <- ret:
		return true
	case <-ctx.Done():
		return false
	}
}

func YieldBreak(future FutureStream) Future {
	if cap(future) <= 0 {
		exception.Panic("future is void result, cannot yield break")
	}
	close(future)
	return future.Out()
}
