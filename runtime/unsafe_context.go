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

// Init 初始化
func (u _UnsafeContext) Init(svcCtx service.Context, options ContextOptions) {
	u.Context.init(svcCtx, options)
}

// GetOptions 获取运行时上下文所有选项
func (u _UnsafeContext) GetOptions() *ContextOptions {
	return u.getOptions()
}

// SetFrame 设置帧
func (u _UnsafeContext) SetFrame(frame Frame) {
	u.setFrame(frame)
}

// SetCallee 设置调用接受者
func (u _UnsafeContext) SetCallee(callee async.Callee) {
	u.setCallee(callee)
}

// GetServiceCtx 获取服务上下文
func (u _UnsafeContext) GetServiceCtx() service.Context {
	return u.getServiceCtx()
}

// ChangeRunningStatus 修改运行状态
func (u _UnsafeContext) ChangeRunningStatus(status RunningStatus, args ...any) {
	u.changeRunningStatus(status, args...)
}

// GC GC
func (u _UnsafeContext) GC() {
	u.gc()
}
