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

// Package meta 提供按字符串键排序的元数据映射。
/*
Package meta 把 `map[string]any` 封装为基于 SliceMap 的元数据结构，并提供链式构造器，
便于在实体原型、实体实例和组件描述上附加额外信息。

Meta 始终按键升序保存条目，因此从 Go map 构造也能得到确定的遍历顺序。该类型不
提供并发保护。
*/
package meta
