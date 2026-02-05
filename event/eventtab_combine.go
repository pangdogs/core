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

// CombineEventTab 联合事件表，可以将多个事件表联合在一起，方便管理多个事件表
type CombineEventTab []IEventTab

// SetPanicHandling 设置panic时的处理方式
func (c *CombineEventTab) SetPanicHandling(autoRecover bool, reportError chan error) {
	for _, tab := range *c {
		tab.Ctrl().SetPanicHandling(autoRecover, reportError)
	}
}

// SetRecursion 设置发生事件递归时的处理方式
func (c *CombineEventTab) SetRecursion(recursion EventRecursion) {
	for _, tab := range *c {
		tab.Ctrl().SetRecursion(recursion)
	}
}

// SetEnabled 设置事件是否启用
func (c *CombineEventTab) SetEnabled(b bool) {
	for _, tab := range *c {
		tab.Ctrl().SetEnabled(b)
	}
}

// UnbindAll 解绑定所有订阅者
func (c *CombineEventTab) UnbindAll() {
	for _, tab := range *c {
		tab.Ctrl().UnbindAll()
	}
}

// Ctrl 事件控制器
func (c *CombineEventTab) Ctrl() IEventCtrl {
	return c
}

// Event 获取事件
func (c *CombineEventTab) Event(id uint64) IEvent {
	for _, tab := range *c {
		event := tab.Event(id)
		if event != nil {
			return event
		}
	}
	return nil
}
