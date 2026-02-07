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

// TaskQueueStats 任务队列统计信息
type TaskQueueStats struct {
	Enqueued  int64
	Pending   int64
	Rejected  int64
	Completed int64
}

// RuntimeStats 运行时统计信息
type RuntimeStats struct {
	WaitGroupCount  int64             // 等待组任务数量
	WaitGroupClosed bool              // 等待组是否已关闭
	TaskQueue       [2]TaskQueueStats // 任务队列统计信息
}

type iRuntimeStats interface {
	// Stats 获取运行时统计信息
	Stats() RuntimeStats
}

// Stats 获取运行时统计信息
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
