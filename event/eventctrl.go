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

// IEventCtrl 统一控制一个事件或一组事件。
type IEventCtrl interface {
	// SetPanicHandling 设置订阅者 panic 的恢复与上报方式。
	SetPanicHandling(autoRecover bool, reportError chan error)
	// SetRecursion 设置递归派发策略。
	SetRecursion(recursion EventRecursion)
	// SetEnabled 设置事件是否启用；禁用会解绑全部订阅者。
	SetEnabled(b bool)
	// UnbindAll 解绑全部订阅者。
	UnbindAll()
}
