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

package service

// RetainedAddIn 标记在 Service 终止后仍保留的插件。
//
// RetainedAddIn 不执行停服 Shut，也不会从插件管理器移除，其状态继续保持 Running。
// 此类插件必须能够在 Service Context 已取消后继续使用，且不得依赖需要主动关闭的
// 后台任务或外部资源；插件最终随 Service Context 一同由 GC 回收。
type RetainedAddIn interface {
	RetainAfterTermination()
}
