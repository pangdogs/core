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

// TaskQueueStats 描述一类运行时任务的累计队列指标。
type TaskQueueStats struct {
	Enqueued  int64 // 成功入队的任务总数。
	Pending   int64 // 当前等待执行的任务数。
	Rejected  int64 // 因队列关闭或已满而拒绝的任务总数。
	Completed int64 // 已完成的任务总数。
}

// RuntimeStats 描述运行时当前的等待组和任务队列状态。
type RuntimeStats struct {
	WaitGroupCount  int64             // 尚未完成的等待组任务数。
	WaitGroupClosed bool              // 等待组是否已经关闭并拒绝新任务。
	TaskQueue       [2]TaskQueueStats // 按 TaskType 索引的任务队列统计信息。
}

type iRuntimeStats interface {
	// Stats 返回运行时统计信息的当前快照。
	Stats() RuntimeStats
}

// Stats 返回运行时统计信息的当前快照。
func (rt *RuntimeBehavior) Stats() RuntimeStats {
	return RuntimeStats{
		WaitGroupCount:  rt.ctx.WaitGroup().Count(),
		WaitGroupClosed: rt.ctx.WaitGroup().Closed(),
		TaskQueue: [2]TaskQueueStats{
			{
				Enqueued:  rt.taskQueue.stats[0].enqueued.Load(),
				Pending:   rt.taskQueue.stats[0].pending.Load(),
				Rejected:  rt.taskQueue.stats[0].rejected.Load(),
				Completed: rt.taskQueue.stats[0].completed.Load(),
			},
			{
				Enqueued:  rt.taskQueue.stats[1].enqueued.Load(),
				Pending:   rt.taskQueue.stats[1].pending.Load(),
				Rejected:  rt.taskQueue.stats[1].rejected.Load(),
				Completed: rt.taskQueue.stats[1].completed.Load(),
			},
		},
	}
}
