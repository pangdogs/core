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

package runtime

import (
	"sync/atomic"

	"git.golaxy.org/core/service"
)

// Deprecated: UnsafeContext 暴露运行时上下文内部能力，仅供框架集成代码使用。
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// Options 返回运行时上下文当前使用的选项。
func (u _UnsafeContext) Options() *ContextOptions {
	return u.getOptions()
}

// Instance 返回运行时上下文的实际实例。
func (u _UnsafeContext) Instance() Context {
	return u.getInstance()
}

// EmitEventRunningEvent 直接派发运行时运行事件。
func (u _UnsafeContext) EmitEventRunningEvent(runningEvent RunningEvent, args ...any) {
	u.emitEventRunningEvent(runningEvent, args...)
}

// SetFrame 设置运行时的帧统计接口。
func (u _UnsafeContext) SetFrame(frame Frame) {
	u.setFrame(frame)
}

// SetCaller 设置 Runtime 邮箱调用接口。
func (u _UnsafeContext) SetCaller(caller Caller) {
	u.setCaller(caller)
}

// ServiceContext 返回所属服务上下文。
func (u _UnsafeContext) ServiceContext() service.Context {
	return u.getServiceContext()
}

// AddInManager 返回运行时专用插件管理器。
func (u _UnsafeContext) AddInManager() AddInManager {
	return u.getAddInManager()
}

// Scoped 返回上下文是否已经绑定运行时作用域的原子标记。
func (u _UnsafeContext) Scoped() *atomic.Bool {
	return u.getScoped()
}

// GC 清理当前运行时收集的对象。
func (u _UnsafeContext) GC() {
	u.gc()
}
