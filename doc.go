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

// Package core 提供框架的公共入口与跨包编排能力。
/*
Package core 把 service、runtime、ec、event 和 extension 等子包串联成完整的
运行模型。

典型流程如下：

  1. 用 service.NewContext 创建服务上下文。
  2. 用 BuildEntityPT 或 service.Context.EntityLib() 声明实体原型。
  3. 用 runtime.NewContext 创建运行时上下文。
  4. 用 NewService 和 NewRuntime 启动服务与运行时。
  5. 用 BuildEntity 在运行时中创建实体，并让组件生命周期自动推进。

此外，根包还提供：

  - 实体、组件、插件的生命周期接口；
  - 异步调度与等待工具，例如 CallAsync、Await、GoAsync 和 TimeAfterAsync；
  - Service、Runtime、Frame 与 TaskQueue 的选项构造器；
  - 面向高级扩展场景的 unsafe 辅助入口。

业务模型与底层数据结构分别位于 ec、runtime、service、event、extension、
define 与 utils 子包中。
*/
package core
