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

import (
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

// LifecycleAddInInit 插件初始化回调，插件实现此接口即可使用，当插件安装在服务上时，rtCtx为nil
type LifecycleAddInInit interface {
	Init(svcCtx service.Context, rtCtx runtime.Context)
}

// LifecycleAddInShut 插件结束回调，插件实现此接口即可使用，当插件安装在服务上时，rtCtx为nil
type LifecycleAddInShut interface {
	Shut(svcCtx service.Context, rtCtx runtime.Context)
}

// LifecycleRuntimeAddInInit 运行时插件初始化回调，当插件安装在运行时上时，实现此接口即可使用
type LifecycleRuntimeAddInInit interface {
	Init(rtCtx runtime.Context)
}

// LifecycleRuntimeAddInShut 运行时插件结束回调，当插件安装在运行时上时，实现此接口即可使用
type LifecycleRuntimeAddInShut interface {
	Shut(rtCtx runtime.Context)
}

// LifecycleAddInOnRuntimeRunningStatusChanged 运行时运行状态变化，当插件安装在运行时上时，实现此接口即可使用
type LifecycleAddInOnRuntimeRunningStatusChanged = eventRuntimeRunningStatusChanged

// LifecycleServiceAddInInit 服务插件初始化回调，当插件安装在服务上时，实现此接口即可使用
type LifecycleServiceAddInInit interface {
	Init(svcCtx service.Context)
}

// LifecycleServiceAddInShut 服务插件结束回调，当插件安装在服务上时，实现此接口即可使用
type LifecycleServiceAddInShut interface {
	Shut(svcCtx service.Context)
}

// LifecycleAddInOnServiceRunningStatusChanged 服务运行状态变化，当插件安装在服务上时，实现此接口即可使用
type LifecycleAddInOnServiceRunningStatusChanged interface {
	OnServiceRunningStatusChanged(svcCtx service.Context, status service.RunningStatus, args ...any)
}
