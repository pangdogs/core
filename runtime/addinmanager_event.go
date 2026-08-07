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

//go:generate go run git.golaxy.org/core/event/eventc event
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=addInManagerEventTab
package runtime

import "git.golaxy.org/core/extension"

// EventInstallAddIn 在插件进入 Loaded 状态后派发，用于执行激活流程。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventInstallAddIn interface {
	OnInstallAddIn(status AddInStatus)
}

// EventUninstallAddIn 在 Running 插件从管理器移除前派发，用于执行停用流程。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventUninstallAddIn interface {
	OnUninstallAddIn(status AddInStatus)
}

// EventAddInStateChanged 在插件状态完成变更后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventAddInStateChanged interface {
	OnAddInStateChanged(status AddInStatus, state extension.AddInState)
}
