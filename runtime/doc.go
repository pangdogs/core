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

所有会直接读写实体或组件状态的逻辑，通常都应回到所属 runtime 中执行。可以通过
Context.CallAsync / CallVoidAsync，或者根包中的 CallAsync / Await 辅助函数，
把工作调度回运行时线程。

用 NewContext 创建上下文后，再交给 core.NewRuntime 绑定和运行。实体实例通常在
runtime.RunningEvent_Started 阶段通过 core.BuildEntity 创建。
*/
package runtime
