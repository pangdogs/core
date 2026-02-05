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

	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
)

// Deprecated: UnsafeContext 访问运行时上下文内部方法
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// Options 获取运行时上下文所有选项
func (u _UnsafeContext) Options() *ContextOptions {
	return u.getOptions()
}

// EmitEventRunningEvent 发送运行事件
func (u _UnsafeContext) EmitEventRunningEvent(runningEvent RunningEvent, args ...any) {
	u.emitEventRunningEvent(runningEvent, args...)
}

// SetFrame 设置帧
func (u _UnsafeContext) SetFrame(frame Frame) {
	u.setFrame(frame)
}

// SetCallee 设置调用接受者
func (u _UnsafeContext) SetCallee(callee async.Callee) {
	u.setCallee(callee)
}

// ServiceCtx 获取服务上下文
func (u _UnsafeContext) ServiceCtx() service.Context {
	return u.getServiceCtx()
}

// AddInManager 获取插件管理器
func (u _UnsafeContext) AddInManager() extension.RuntimeAddInManager {
	return u.getAddInManager()
}

// Scoped 获取作用域状态
func (u _UnsafeContext) Scoped() *atomic.Bool {
	return u.getScoped()
}

// GC GC
func (u _UnsafeContext) GC() {
	u.gc()
}
