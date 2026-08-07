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

package core

// Deprecated: UnsafeRuntime 暴露运行时内部能力，仅供框架集成代码使用。
func UnsafeRuntime(runtime Runtime) _UnsafeRuntime {
	return _UnsafeRuntime{
		Runtime: runtime,
	}
}

type _UnsafeRuntime struct {
	Runtime
}

// Options 返回运行时当前使用的选项。
func (u _UnsafeRuntime) Options() *RuntimeOptions {
	return u.getOptions()
}

// Instance 返回运行时的实际实例。
func (u _UnsafeRuntime) Instance() Runtime {
	return u.getInstance()
}
