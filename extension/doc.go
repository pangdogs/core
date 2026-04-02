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

// Package extension 实现插件管理器与插件状态机。
/*
Package extension 提供 service 与 runtime 两种插件管理器，以及插件安装、运行、
卸载过程中的状态跟踪和事件通知。

大多数业务代码只需要通过 define 包声明插件，再调用 Install/Require 即可。
当需要直接观察插件生命周期、实现自定义管理器，或操作低层状态对象时，再使用
extension 包暴露的 AddInManager、Status 和事件类型。
*/
package extension
