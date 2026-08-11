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

// Deprecated: UnsafeConcurrentContext 暴露并发上下文的内部能力，仅供框架集成代码使用。
func UnsafeConcurrentContext(context ConcurrentContext) _UnsafeConcurrentContext {
	return _UnsafeConcurrentContext{
		ConcurrentContext: context,
	}
}

type _UnsafeConcurrentContext struct {
	ConcurrentContext
}

// Instance 返回实际运行时上下文实例；调用者必须自行保证运行协程约束。
func (u _UnsafeConcurrentContext) Instance() Context {
	return u.getInstance()
}
