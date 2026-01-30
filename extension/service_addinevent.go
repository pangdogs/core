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

// EventServiceInstallAddIn 服务安装插件事件
type EventServiceInstallAddIn struct {
	Status AddInStatus
}

// EventServiceUninstallAddIn 服务卸载插件事件
type EventServiceUninstallAddIn struct {
	Status AddInStatus
}

// EventServiceAddInSnapshot 服务插件快照事件
type EventServiceAddInSnapshot struct {
	Statuses []AddInStatus
}
