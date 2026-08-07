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

// ServiceStats 描述服务当前的等待组状态。
type ServiceStats struct {
	WaitGroupCount  int64 // 尚未完成的等待组任务数。
	WaitGroupClosed bool  // 等待组是否已经关闭并拒绝新任务。
}

type iServiceStats interface {
	// Stats 返回服务统计信息的当前快照。
	Stats() ServiceStats
}

// Stats 返回服务统计信息的当前快照。
func (svc *ServiceBehavior) Stats() ServiceStats {
	return ServiceStats{
		WaitGroupCount:  svc.ctx.WaitGroup().Count(),
		WaitGroupClosed: svc.ctx.WaitGroup().Closed(),
	}
}
