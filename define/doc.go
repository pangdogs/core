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

// Package define 为插件声明提供类型安全的辅助定义。
/*
Package define 把插件构造函数、接口类型和插件名称封装成可复用的定义对象，避免
业务代码反复编写名称、ID 计算和类型断言。

定义对象会暴露 Install、Uninstall、Require 和 Lookup 等操作，适合在包级变量中
声明后复用。常用入口包括：

  - AddIn：同时适用于 service 与 runtime 的通用插件定义；
  - ServiceAddIn / RuntimeAddIn：分别提供 service.Context 与 runtime.Context 类型的
    Require 和 Lookup 函数；
  - AddInInterface 等 Interface 变体：只声明依赖契约，不绑定构造函数。
*/
package define
