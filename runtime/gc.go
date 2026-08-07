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

// GC 表示可由运行时延迟清理的对象。
type GC interface {
	// GC 执行实际清理。
	GC()
	// NeedGC 报告对象当前是否需要清理。
	NeedGC() bool
}

// GCCollector 收集需要在后续运行时 GC 阶段清理的对象。
type GCCollector interface {
	// CollectGC 在 gc 需要清理时将其加入当前运行时的待清理队列。
	CollectGC(gc GC)
}
