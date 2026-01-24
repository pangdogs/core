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

package service

import (
	"sync/atomic"

	"git.golaxy.org/core/extension"
)

// Deprecated: UnsafeContext 访问服务上下文内部方法
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// GetOptions 获取服务上下文所有选项
func (u _UnsafeContext) GetOptions() *ContextOptions {
	return u.getOptions()
}

// EmitEventRunningEvent 发送运行事件
func (u _UnsafeContext) EmitEventRunningEvent(runningEvent RunningEvent, args ...any) {
	u.emitEventRunningEvent(runningEvent, args...)
}

// GetAddInManager 获取插件管理器
func (u _UnsafeContext) GetAddInManager() extension.ServiceAddInManager {
	return u.getAddInManager()
}

// GetScoped 获取作用域状态
func (u _UnsafeContext) GetScoped() *atomic.Bool {
	return u.getScoped()
}
