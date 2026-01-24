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

package extension

//go:generate go run git.golaxy.org/core/event/eventc event
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=runtimeAddInManagerEventTab

// EventRuntimeInstallAddIn 事件：运行时安装插件
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventRuntimeInstallAddIn interface {
	OnRuntimeInstallAddIn(status AddInStatus)
}

// EventRuntimeUninstallAddIn 事件：运行时卸载插件
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventRuntimeUninstallAddIn interface {
	OnRuntimeUninstallAddIn(status AddInStatus)
}

// EventRuntimeAddInStateChanged 事件：运行时插件状态改变
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventRuntimeAddInStateChanged interface {
	OnRuntimeAddInStateChanged(status AddInStatus, state AddInState)
}
