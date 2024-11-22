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

// LifecycleEntityAwake 实体的生命周期进入唤醒（Awake）时的回调，实体实现此接口即可使用
type LifecycleEntityAwake interface {
	Awake()
}

// LifecycleEntityStart 实体的生命周期进入开始（Start）时的回调，实体实现此接口即可使用
type LifecycleEntityStart interface {
	Start()
}

// LifecycleEntityUpdate 如果开启运行时的帧更新特性，那么实体状态为活跃（Alive）时，将会收到这个帧更新（Update）回调，实体实现此接口即可使用
type LifecycleEntityUpdate = eventUpdate

// LifecycleEntityLateUpdate 如果开启运行时的帧更新特性，那么实体状态为活跃（Alive）时，将会收到这个帧迟滞更新（Late Update）回调，实体实现此接口即可使用
type LifecycleEntityLateUpdate = eventLateUpdate

// LifecycleEntityShut 实体的生命周期进入结束（Shut）时的回调，实体实现此接口即可使用
type LifecycleEntityShut interface {
	Shut()
}

// LifecycleEntityDispose 实体的生命周期进入死亡（Death）时的回调，实体实现此接口即可使用
type LifecycleEntityDispose interface {
	Dispose()
}
