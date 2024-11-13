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

package event

// IEventTab 事件表接口，方便访问多个事件
/*
使用方式：
	1.在定义事件的源码文件（.go）头部添加以下注释，在编译前自动化生成代码：
	//go:generate go run git.golaxy.org/core/event/eventc eventtab --name={事件表名称}

定义事件的选项（添加到定义事件的注释里）：
	1.事件表初始化时，该事件使用的递归处理方式，不填表示使用事件表初始化参数值
		+event-tab-gen:recursion=[allow,disallow,discard,truncate,deepest]
*/
type IEventTab interface {
	// Event 获取事件
	Event(id uint64) IEvent
}

// IEventCtrlTab 事件控制表接口，方便管理多个事件
type IEventCtrlTab interface {
	IEventCtrl
	// Event 获取事件
	Event(id uint64) IEvent
}
