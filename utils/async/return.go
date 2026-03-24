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

import "git.golaxy.org/core/utils/exception"

func Return(future FutureChan, ret Result) Future {
	if future.ch == nil || future.done == nil {
		exception.Panic("future is void result, cannot return")
	}
	future.ch <- ret
	close(future.ch)
	close(future.done)
	return future.Out()
}

func ReturnVoid(future FutureVoid) Future {
	close(future)
	return future.Out()
}
