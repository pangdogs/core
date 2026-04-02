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

// Package utils 是核心工具包的命名空间根目录。
/*
Package utils 本身不承载业务逻辑，主要用于组织框架内部共享的工具子包。

常用子包包括：

  - assertion：基于反射的组件注入与组合视图提取；
  - async：Future/Result 抽象；
  - corectx：service 与 runtime 共用的上下文接口；
  - exception：错误与 panic 辅助；
  - generic：泛型函数类型、容器与并发基础设施；
  - iface / reinterpret：高性能接口缓存与重新解释工具；
  - meta、option、types、uid：元数据、选项、反射类型和唯一 Id 辅助。
*/
package utils
