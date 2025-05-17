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

package ictx

import (
	"context"
	"sync/atomic"
)

// Deprecated: UnsafeContext 访问上下文内部方法
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// Init 初始化
func (u _UnsafeContext) Init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	u.init(parentCtx, autoRecover, reportError)
}

// GetPaired 获取配对标记
func (u _UnsafeContext) GetPaired() *atomic.Bool {
	return u.getPaired()
}

// ReturnTerminated 返回已停止
func (u _UnsafeContext) ReturnTerminated() {
	u.returnTerminated()
}
