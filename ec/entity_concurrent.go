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

package ec

import (
	"context"
	"fmt"

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/uid"
)

// ConcurrentEntity 暴露可跨协程安全读取的实体信息与上下文。
//
// 组件管理、实体树和 Destroy 等操作仍须通过所属 Runtime 的运行协程执行。
type ConcurrentEntity interface {
	iConcurrentEntity
	iContext
	corectx.ConcurrentContextProvider
	fmt.Stringer

	// Id 返回实体 ID。
	Id() uid.Id
	// PT 返回实体原型。
	PT() EntityPT
}

type iContext interface {
	context.Context

	// Terminated 返回实体完成销毁时兑现的 Future。
	Terminated() async.Future
}

type iConcurrentEntity interface {
	getEntity() Entity
}

// Terminated 返回实体进入 Destroyed 状态时兑现的 Future。
func (entity *EntityBehavior) Terminated() async.Future {
	return entity.terminated.Out()
}

func (entity *EntityBehavior) getEntity() Entity {
	return entity.options.InstanceFace.Iface
}
