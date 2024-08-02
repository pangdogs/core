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

// LifecycleComponentAwake 组件的生命周期进入唤醒（awake）时的回调，组件实现此接口即可使用
type LifecycleComponentAwake interface {
	Awake()
}

// LifecycleComponentStart 组件的生命周期进入开始（start）时的回调，组件实现此接口即可使用
type LifecycleComponentStart interface {
	Start()
}

// LifecycleComponentUpdate 如果开启运行时的帧更新特性，那么组件状态为活跃（alive）时，将会收到这个帧更新（update）回调，组件实现此接口即可使用
type LifecycleComponentUpdate = eventUpdate

// LifecycleComponentLateUpdate 如果开启运行时的帧更新特性，那么组件状态为活跃（alive）时，将会收到这个帧迟滞更新（late update）回调，组件实现此接口即可使用
type LifecycleComponentLateUpdate = eventLateUpdate

// LifecycleComponentShut 组件的生命周期进入结束（shut）时的回调，组件实现此接口即可使用
type LifecycleComponentShut interface {
	Shut()
}

// LifecycleComponentDispose 组件的生命周期进入死亡（death）时的回调，组件实现此接口即可使用
type LifecycleComponentDispose interface {
	Dispose()
}
