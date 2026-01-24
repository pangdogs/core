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

// IEventCtrl 事件控制器接口
type IEventCtrl interface {
	// SetPanicHandling 设置panic时的处理方式
	SetPanicHandling(autoRecover bool, reportError chan error)
	// SetRecursion 设置发生事件递归时的处理方式
	SetRecursion(recursion EventRecursion)
	// SetEnable 设置事件是否启用
	SetEnable(b bool)
	// UnbindAll 解绑定所有订阅者
	UnbindAll()
}
