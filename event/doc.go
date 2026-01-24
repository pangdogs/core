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

// Package event 高效的事件系统，适用于单线程环境，需要使用 go:generate 功能来生成代码。
/*
定义事件：
	1.按以下格式编写一个接口，即完成事件的定义：
	type Event{事件名} interface {
		On{事件名}({参数列表})
	}

	2.在定义事件的源码文件（.go）头部添加以下注释，在编译前自动化生成代码：
	//go:generate go run git.golaxy.org/core/event/eventc event

使用事件：
	1.事件一般作为组件的成员，在组件 Awake 时初始化，组件 Dispose 时关闭，示例如下：
	type Comp struct {
		ec.ComponentBehavior
		event{事件名} event.Event
	}
	func (c *Comp) Awake() {
		runtime.Current(c).ActivateEvent(&c.event{事件名}, event.EventRecursion_Discard)
	}
	func (c *Comp) Dispose() {
		c.event{事件名}.Disable()
	}

订阅事件：
	1.在组件的成员函数，编写以下代码：
	func (c *Comp) On{事件名}({参数列表}) {
		...
	}

	2.在需要订阅事件时，编写以下代码：
	func (c *Comp) MethodXXX() {
		{事件定义包名}.Bind{事件名}({发布者}, c)
	}

	3.如果订阅者生命周期小于发布者，那么需要记录 handle 并且在 Dispose 时解除绑定，示例如下：
	type Comp struct {
		ec.ComponentBehavior
		handle event.Handle
	}
	func (c *Comp) MethodXXX() {
		c.handle = {事件定义包名}.Bind{事件名}({发布者}, c)
	}
	func (c *Comp) Dispose() {
		c.handle.Unbind()
	}

	4.如果不想写代码记录 handle，可以使用 ec.Component、ec.Entity 或 runtime.Context 的 ManagedAddEventHandles() 来记录 handle，在它们生命周期结束时，将会自动解除绑定

定义事件表：
	1.在定义事件的源码文件（.go）头部添加以下注释，在编译前自动化生成代码：
	//go:generate go run git.golaxy.org/core/event/eventc eventtab --name={事件表名称}

事件的选项（添加到定义事件的注释里）：
	1.发送事件的代码的可见性
		+event-gen:export_emit=[0,1]

	2.是否生成简化绑定事件的代码
		+event-gen:auto=[0,1]

	3.事件表初始化时，该事件使用的递归处理方式，不填表示使用事件表初始化参数值
		+event-tab-gen:recursion=[allow,disallow,discard,skip_received,receive_once]
*/
package event
