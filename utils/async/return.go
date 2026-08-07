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

// Return 向 future 写入最后一项结果，关闭结果与完成频道，并返回消费者视图。
//
// future 为零值、已结束或缓冲区已满且没有消费者时，调用会分别 panic 或阻塞。
func Return(future FutureChan, ret Result) Future {
	if future.ch == nil || future.done == nil {
		exception.Panic("future is void result, cannot return")
	}
	future.ch <- ret
	close(future.ch)
	close(future.done)
	return future.Out()
}

// ReturnVoid 关闭无结果 Future 并返回消费者视图；重复关闭会 panic。
func ReturnVoid(future FutureVoid) Future {
	close(future)
	return future.Out()
}
