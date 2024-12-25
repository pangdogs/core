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

// LifecycleAddInOnRuntimeRunningStatusChanged 运行时运行状态变化，当插件安装在运行时上时，插件实现此接口即可使用
type LifecycleAddInOnRuntimeRunningStatusChanged = eventRuntimeRunningStatusChanged
