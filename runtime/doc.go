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

// Package runtime 定义运行时级上下文。
/*
Package runtime 表示 Actor 风格的单运行时执行作用域。一个运行时拥有自己的任务
队列、可选帧循环、本地实体管理器、实体树、运行事件和 runtime add-in。

runtime add-in 由 runtime 包内的管理器维护，可在运行协程中安装和卸载。

所有会直接读写实体或组件状态的逻辑，通常都应回到所属 runtime 中执行。需要结果
时使用 Context.Submit，需要无 Future 投递时使用 Context.Post；后台 Future 完成后，
使用根包 core.ContinueOn 把续体调度回运行协程。Delegate 与 DelegateVoid 变体仍然
可用。

Runtime Context 会拒绝对 pending Future 的阻塞等待。同一执行器产生的 Future 返回
ErrRuntimeSelfWait，其他 pending Future 返回 ErrBlockingWaitInRuntime。应使用
ContinueOn 表达非阻塞的 Actor 续体。

用 NewContext 创建上下文后，再交给 core.NewRuntime 绑定和运行。实体实例通常在
runtime.RunningEvent_Started 阶段通过 core.BuildEntity 创建。
*/
package runtime
