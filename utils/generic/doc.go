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

// Package generic 提供框架内部常用的泛型基础设施。
/*
Package generic 包含三类能力：

  - Func / Action / Delegate 等函数类型及其安全调用包装；
  - FreeList、SliceMap、Barrier、EventStream、UnboundedChannel 等通用数据结构；
  - Bits、ReentrancyGuard 等状态与并发辅助工具。

这些类型是 service、runtime、event、extension 等包的实现基础，也可以被业务代码
直接复用来表达统一的泛型回调协议。

Func/Action 的 UnsafeCall 不恢复 panic，SafeCall 会恢复并返回错误，Call 则允许调用方
选择策略；nil 函数按零值或空操作处理。Delegate 按切片顺序调用多个函数，并可通过
interrupt 回调提前终止。

除 Barrier、EventStream 和 UnboundedChannel 明确提供并发保护外，容器与辅助类型默认
要求调用方串行访问。
*/
package generic
