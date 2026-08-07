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

// LifecycleEntityAwake 在实体进入 Awakened 状态时调用，每个实体最多调用一次。
type LifecycleEntityAwake interface {
	Awake()
}

// LifecycleEntityStart 在实体进入 Starting 状态时调用，每个实体最多调用一次。
type LifecycleEntityStart interface {
	Start()
}

// LifecycleEntityUpdate 在启用帧循环且实体处于 Alive 状态时接收每帧更新。
type LifecycleEntityUpdate = eventUpdate

// LifecycleEntityLateUpdate 在每帧普通更新结束后接收后置更新。
type LifecycleEntityLateUpdate = eventLateUpdate

// LifecycleEntityShut 在已开始的实体进入 Shutting 状态时调用，与 LifecycleEntityStart 成对。
type LifecycleEntityShut interface {
	Shut()
}

// LifecycleEntityDispose 在已唤醒的实体进入 Dead 状态时调用，与 LifecycleEntityAwake 成对。
type LifecycleEntityDispose interface {
	Dispose()
}
