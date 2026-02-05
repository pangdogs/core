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

// RuntimeStats 运行时统计信息
type RuntimeStats struct {
	WaitGroupCount          int64 // 等待组任务数量
	WaitGroupClosed         bool  // 等待组是否已关闭
	TaskQueueCallEnqueued   int64 // 任务队列调用任务入队数量
	TaskQueueCallCompleted  int64 // 任务队列调用任务完成数量
	TaskQueueFrameEnqueued  int64 // 任务队列帧任务入队数量
	TaskQueueFrameCompleted int64 // 任务队列帧任务完成数量
}

type iRuntimeStats interface {
	// Stats 获取运行时统计信息
	Stats() RuntimeStats
}

// Stats 获取运行时统计信息
func (rt *RuntimeBehavior) Stats() RuntimeStats {
	return RuntimeStats{
		WaitGroupCount:          rt.ctx.WaitGroup().Count(),
		WaitGroupClosed:         rt.ctx.WaitGroup().Closed(),
		TaskQueueCallEnqueued:   rt.taskQueue.callEnqueued.Load(),
		TaskQueueCallCompleted:  rt.taskQueue.callCompleted.Load(),
		TaskQueueFrameEnqueued:  rt.taskQueue.frameEnqueued.Load(),
		TaskQueueFrameCompleted: rt.taskQueue.frameCompleted.Load(),
	}
}
