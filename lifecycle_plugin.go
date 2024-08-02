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

// LifecycleServicePluginInit 服务上的插件初始化回调，插件实现此接口即可使用
type LifecycleServicePluginInit interface {
	InitSP(ctx service.Context)
}

// LifecycleServicePluginShut 服务上的插件结束回调，插件实现此接口即可使用
type LifecycleServicePluginShut interface {
	ShutSP(ctx service.Context)
}

// LifecycleRuntimePluginInit 运行时上的插件初始化回调，插件实现此接口即可使用
type LifecycleRuntimePluginInit interface {
	InitRP(ctx runtime.Context)
}

// LifecycleRuntimePluginShut 运行时上的插件结束回调，插件实现此接口即可使用
type LifecycleRuntimePluginShut interface {
	ShutRP(ctx runtime.Context)
}
