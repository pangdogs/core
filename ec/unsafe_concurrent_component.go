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

package ec

// UnsafeConcurrentComponent 暴露 ConcurrentComponent 的框架内部能力。
//
// Deprecated: 仅供框架内部使用。
func UnsafeConcurrentComponent(component ConcurrentComponent) _UnsafeConcurrentComponent {
	return _UnsafeConcurrentComponent{ConcurrentComponent: component}
}

type _UnsafeConcurrentComponent struct {
	ConcurrentComponent
}

// Instance 返回实际组件实例；调用者必须自行保证 Runtime goroutine 约束。
func (u _UnsafeConcurrentComponent) Instance() Component {
	return u.getInstance()
}
