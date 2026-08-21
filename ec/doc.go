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

// Package ec 定义实体—组件模型及其生命周期状态。
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

Entity 的组件容器独立于 Runtime 生命周期保持可用。Entity 进入 Leaving 后新增的
Component 仍会进入 Attached，但 Runtime 不再推进其 Awake、OnEnable 或 Start；Entity
进入 Dead 后组件管理器事件表关闭，后续组件操作只影响本地组件表。

除 ConcurrentEntity 与 ConcurrentComponent 明确暴露的能力外，实体、组件及实体树
操作都应在所属 Runtime 的运行 goroutine 中执行。Entity 成功加入 Runtime、Component
完成所属 Runtime 的身份初始化后，才能把对应并发视图发布给其他 goroutine。

两个并发视图均提供 Lifetime AsyncScope。Entity Scope 在绑定 Runtime Context 时创建，
Component Scope 在首次访问时懒创建；Entity 销毁或 Component 移除时会关闭对应 Scope，
取消其中的协作式后台任务并禁止新任务。Entity 和 Component 的销毁流程不会阻塞等待
这些任务退出，需要汇合时应在其他 goroutine 等待 Scope.Completion 返回的 Signal。
*/
package ec
