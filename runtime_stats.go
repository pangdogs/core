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

import "git.golaxy.org/core/utils/async"

// TaskQueueStats 描述一种调度语义的 Runtime 邮箱统计。
type TaskQueueStats struct {
	Accepted       int64 // 成功进入队列的任务总数。
	Queued         int64 // 当前仅在队列中等待的任务数。
	Running        int64 // 当前正在执行的任务数。
	Completed      int64 // 已进入终态的任务总数。
	Canceled       int64 // Completed 中因调度上下文取消结束的数量。
	Panicked       int64 // Completed 中恢复过 panic 的数量。
	RejectedClosed int64 // 因队列关闭而拒绝的数量。
	RejectedFull   int64 // 因有界队列容量不足而拒绝的数量。
}

// RuntimeTaskStats 按 Submit、Post 和 Frame 三种调度语义分类。
type RuntimeTaskStats struct {
	Submit TaskQueueStats
	Post   TaskQueueStats
	Frame  TaskQueueStats
}

// RuntimeHealthStats 描述 Runtime 当前执行健康状态。
type RuntimeHealthStats struct {
	LastProgressTime int64          // 最近一次开始或完成任务的 UnixNano。
	BlockedFutureID  async.FutureID // 最近由 Runtime Context 尝试阻塞等待的 Future ID。
	LastWaitRejectID async.FutureID // 最近一次被自等待规则拒绝的 Future ID。
}

// RuntimeStats 描述 Runtime 的生命周期、邮箱、异步作用域和健康状态快照。
// 各字段通过独立原子读取获得，不保证是同一时刻的事务快照。
type RuntimeStats struct {
	WaitGroupCount  int64
	WaitGroupClosed bool
	Tasks           RuntimeTaskStats
	Scope           async.ScopeStats
	Health          RuntimeHealthStats
}

type iRuntimeStats interface {
	Stats() RuntimeStats
}

func snapshotTaskStats(stats *_TaskQueueStats) TaskQueueStats {
	ret := TaskQueueStats{
		Accepted:       stats.accepted.Load(),
		Queued:         stats.queued.Load(),
		Running:        stats.running.Load(),
		Completed:      stats.completed.Load(),
		Canceled:       stats.canceled.Load(),
		Panicked:       stats.panicked.Load(),
		RejectedClosed: stats.rejectedClosed.Load(),
		RejectedFull:   stats.rejectedFull.Load(),
	}
	return ret
}

// Stats 返回 Runtime 的并发安全、近似瞬时统计快照。
func (rt *RuntimeBehavior) Stats() RuntimeStats {
	return RuntimeStats{
		WaitGroupCount:  rt.ctx.WaitGroup().Count(),
		WaitGroupClosed: rt.ctx.WaitGroup().Closed(),
		Tasks: RuntimeTaskStats{
			Submit: snapshotTaskStats(&rt.taskQueue.stats[TaskType_Submit]),
			Post:   snapshotTaskStats(&rt.taskQueue.stats[TaskType_Post]),
			Frame:  snapshotTaskStats(&rt.taskQueue.stats[TaskType_Frame]),
		},
		Scope: rt.ctx.AsyncScope().Stats(),
		Health: RuntimeHealthStats{
			LastProgressTime: rt.lastProgressTime.Load(),
			BlockedFutureID:  rt.ctx.BlockedFutureID(),
			LastWaitRejectID: rt.ctx.LastWaitRejectID(),
		},
	}
}
