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

// CombineEventTab 将多张事件表组合成一个 IEventTab 与 IEventCtrl。
//
// 查询按切片顺序进行，控制操作会依次作用于每张表。零值可用，但元素不能为 nil。
type CombineEventTab []IEventTab

// SetPanicHandling 为全部事件表设置订阅者 panic 的恢复与上报方式。
func (c *CombineEventTab) SetPanicHandling(autoRecover bool, reportError chan error) {
	for _, tab := range *c {
		tab.Ctrl().SetPanicHandling(autoRecover, reportError)
	}
}

// SetRecursion 为全部事件表设置默认递归派发策略。
func (c *CombineEventTab) SetRecursion(recursion EventRecursion) {
	for _, tab := range *c {
		tab.Ctrl().SetRecursion(recursion)
	}
}

// SetEnabled 设置全部事件表是否启用；禁用会解绑全部订阅者。
func (c *CombineEventTab) SetEnabled(b bool) {
	for _, tab := range *c {
		tab.Ctrl().SetEnabled(b)
	}
}

// UnbindAll 解绑全部事件表的所有订阅者。
func (c *CombineEventTab) UnbindAll() {
	for _, tab := range *c {
		tab.Ctrl().UnbindAll()
	}
}

// Ctrl 返回组合表自身作为控制器。
func (c *CombineEventTab) Ctrl() IEventCtrl {
	return c
}

// Event 按顺序查询首个匹配 ID 的事件；不存在时返回 nil。
func (c *CombineEventTab) Event(id uint64) IEvent {
	for _, tab := range *c {
		event := tab.Event(id)
		if event != nil {
			return event
		}
	}
	return nil
}
