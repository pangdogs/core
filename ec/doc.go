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

// Package ec 定义实体组件模型。
/*
Package ec 提供框架核心的数据模型：Entity、Component、实体树节点状态、组件启停
事件、作用域与原型描述接口。

常见用法是：

  - 为实体嵌入 EntityBehavior；
  - 为组件嵌入 ComponentBehavior；
  - 通过生命周期接口在 core.Runtime 驱动下接收 Awake/Start/Update/Shut/Dispose；
  - 通过 Entity 的组件管理与树节点接口组合复杂对象结构。

ec 包负责实体与组件本身的状态机与事件表，不负责原型库的声明与注册。原型系统
位于 ec/pt 包，运行时调度与生命周期推进则由根包 core 和 runtime 包完成。
*/
package ec
