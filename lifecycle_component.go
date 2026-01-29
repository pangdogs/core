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

// LifecycleComponentAwake 组件的生命周期进入唤醒（Awake）时的回调，与死亡（Death）成对，只会调用一次，组件实现此接口即可使用
type LifecycleComponentAwake interface {
	Awake()
}

// LifecycleComponentOnEnable 组件的生命周期进入启用（OnEnable）时的回调，与关闭（OnDisable）成对，组件实现此接口即可使用
type LifecycleComponentOnEnable interface {
	OnEnable()
}

// LifecycleComponentStart 组件的生命周期进入开始（Start）时的回调，与结束（Shut）成对，只会调用一次，组件实现此接口即可使用
type LifecycleComponentStart interface {
	Start()
}

// LifecycleComponentUpdate 如果开启运行时的帧更新特性，那么组件状态为活跃（Alive）时，将会收到这个帧更新（Update）回调，组件实现此接口即可使用
type LifecycleComponentUpdate = eventUpdate

// LifecycleComponentLateUpdate 如果开启运行时的帧更新特性，那么组件状态为活跃（Alive）时，将会收到这个帧迟滞更新（Late Update）回调，组件实现此接口即可使用
type LifecycleComponentLateUpdate = eventLateUpdate

// LifecycleComponentShut 组件的生命周期进入结束（Shut）时的回调，与开始（Start）成对，只会调用一次，组件实现此接口即可使用
type LifecycleComponentShut interface {
	Shut()
}

// LifecycleComponentOnDisable 组件的生命周期进入关闭（OnDisable）时的回调，与启用（OnEnable）成对，组件实现此接口即可使用
type LifecycleComponentOnDisable interface {
	OnDisable()
}

// LifecycleComponentDispose 组件的生命周期进入死亡（Death）时的回调，与唤醒（Awake）成对，只会调用一次，组件实现此接口即可使用
type LifecycleComponentDispose interface {
	Dispose()
}
