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

// Package service 定义服务级上下文。
/*
Package service 表示应用的全局作用域，也是多个 runtime 的父上下文。

服务上下文负责：

  - 持有父 context、等待组与终止状态；
  - 管理实体原型库与组件原型库；
  - 提供全局实体索引，以及按实体 Id 发起异步调用；
  - 承载 service add-in，并派发服务运行事件。

通常先用 NewContext 创建上下文，再交给 core.NewService 绑定和运行。
原型声明、插件安装等启动逻辑，通常放在 service.With.RunningEventCB 处理
service.RunningEvent_Birth 或 service.RunningEvent_Started 时完成。
*/
package service
