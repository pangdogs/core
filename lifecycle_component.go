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

// LifecycleComponentAwake 在组件进入 Awaking 状态时调用，每个组件最多调用一次。
// 启用 ComponentAwakeOnFirstTouch 后，正常激活期间的组件访问可以使目标组件优先调用 Awake。
type LifecycleComponentAwake interface {
	Awake()
}

// LifecycleComponentOnEnable 在组件进入 Enabling 状态后调用，可随启用状态切换而多次调用。
type LifecycleComponentOnEnable interface {
	OnEnable()
}

// LifecycleComponentStart 在已启用组件进入 Starting 状态后调用，每个组件最多调用一次。
type LifecycleComponentStart interface {
	Start()
}

// LifecycleComponentUpdate 在启用帧循环且组件处于 Alive 状态时接收每帧更新。
type LifecycleComponentUpdate = eventUpdate

// LifecycleComponentLateUpdate 在每帧普通更新结束后接收后置更新。
type LifecycleComponentLateUpdate = eventLateUpdate

// LifecycleComponentShut 在已进入过 Start 的组件处于 Shutting 状态时调用，与 LifecycleComponentStart 成对。
type LifecycleComponentShut interface {
	Shut()
}

// LifecycleComponentOnDisable 在已进入过 OnEnable 的组件被禁用时调用，与 LifecycleComponentOnEnable 成对。
type LifecycleComponentOnDisable interface {
	OnDisable()
}

// LifecycleComponentDispose 在已进入过 Awake 的组件处于 Dead 状态时调用，与 LifecycleComponentAwake 成对。
type LifecycleComponentDispose interface {
	Dispose()
}
