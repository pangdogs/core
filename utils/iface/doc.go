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

// Package iface 提供接口缓存与 Face 封装。
/*
Package iface 使用基于 unsafe 的缓存表示来降低高频接口重新解释的开销。

Face 用于同时保存“接口值”和“缓存后的底层表示”，便于：

  - 记录实例的默认对外接口；
  - 在不同接口视图之间快速重新解释；
  - 为 define、extension 与 reinterpret 等包提供统一的实例封装。

该包建立在 unsafe 之上，适合框架底层和明确理解其约束的场景。
*/
package iface
