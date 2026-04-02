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

// Package event 提供本地同步事件基础设施。
/*
Package event 是框架内部与业务代码共用的轻量事件系统，适合单线程或由调用方保证
串行访问的场景。

推荐工作流是：

  1. 定义事件接口，例如 `type EventFoo interface { OnFoo(...) }`。
  2. 在同一文件添加 `//go:generate go run git.golaxy.org/core/event/eventc event`。
  3. 如果需要统一管理多个事件，再添加 `eventtab --name=...` 生成事件表。
  4. 通过生成的 `BindEventFoo`、`HandleEventFoo` 和 emit 辅助函数进行绑定与派发。

底层的 Event、Handle、ManagedHandles、IEvent 和 IEventTab 也可以直接使用，
用于实现自定义事件容器、统一解绑和递归策略控制。递归行为由 EventRecursion 控制，
生成代码中的 `+event-gen:*` 与 `+event-tab-gen:*` 注释用于配置可见性、自动绑定
代码和事件表默认递归策略。

可参考 ec、runtime 和 core 包中的 `*_event.go` 与 `*_eventtab.gen.go` 文件了解
推荐用法。
*/
package event
