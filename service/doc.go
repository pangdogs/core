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
  - 提供随服务关闭并汇合后台任务的 AsyncScope；
  - 管理实体原型库与组件原型库；
  - 提供全局实体索引，以及按实体 ID Submit 或 Post 到所属 Runtime；
  - 管理随服务同步启停的 service add-in，并派发服务运行事件。

通常先用 NewContext 创建上下文，再交给 core.NewService 绑定和运行。
service add-in 只能在启动前安装或卸载；启动前卸载只移除尚未激活的插件，不会调用
Shut。管理器会在 RunningEvent_Starting 回调前永久冻结，随后按安装顺序初始化插件。

停服时，Service 会先关闭并等待 AsyncScope 和 WaitGroup（包括已加入的 Runtime），
再按安装顺序的逆序调用普通插件 Shut。Shut 执行期间插件仍处于 Running 状态并
保留在管理器中；回调返回后才从管理器移除并转为 Unloaded。实现 RetainedAddIn 的
插件跳过 Shut 和移除，在 Service 终止后仍保持 Running；这类插件必须能在 Service
Context 已取消后使用，且不能持有需要主动关闭的任务或外部资源。未绑定 Service
Scope 或 WaitGroup 的私有任务与资源，应由普通插件在 Shut 中自行停止和汇合。

原型声明、插件安装等准备逻辑可在创建上下文后完成，也可放在
service.RunningEvent_Birth 中。
*/
package service
