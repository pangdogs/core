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

// Package extension 定义插件的公共协议与辅助函数。
/*
Package extension 提供 AddInProvider、AddInManager、AddInStatus 和 AddInState 等
service 与 runtime 共用的插件协议，以及安装、查询和依赖插件的辅助函数。

具体管理策略由所属上下文实现：service 插件在启动前注册并随服务同步启停，
runtime 插件支持在所属运行时 goroutine 中热插拔。业务代码通常通过 define 包声明
插件，再调用 Install、Require 或 Lookup。Require 只接受正在运行的插件，Lookup
则可查询管理器仍持有但尚未激活的插件。
*/
package extension
