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

// UnsafeConcurrentEntity 暴露 ConcurrentEntity 的框架内部能力。
//
// Deprecated: 仅供框架内部使用。
func UnsafeConcurrentEntity(entity ConcurrentEntity) _UnsafeConcurrentEntity {
	return _UnsafeConcurrentEntity{
		ConcurrentEntity: entity,
	}
}

type _UnsafeConcurrentEntity struct {
	ConcurrentEntity
}

// Instance 返回实际实体实例；调用者必须自行保证运行协程约束。
func (u _UnsafeConcurrentEntity) Instance() Entity {
	return u.getInstance()
}
