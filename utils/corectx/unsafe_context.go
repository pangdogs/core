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

package corectx

import (
	"context"
)

// UnsafeContext 暴露上下文初始化与完成信号等框架内部能力。
//
// Deprecated: 仅供框架内部使用。
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// Init 使用父上下文及 panic 处理策略初始化上下文。
func (u _UnsafeContext) Init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	u.init(parentCtx, autoRecover, reportError)
}

// CloseWaitGroup 关闭任务屏障，使其不再接受新任务。
func (u _UnsafeContext) CloseWaitGroup() {
	u.closeWaitGroup()
}

// ReturnTerminated 兑现宿主已完成清理的 Future。
func (u _UnsafeContext) ReturnTerminated() {
	u.returnTerminated()
}
