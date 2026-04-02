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

// Package corectx 定义 service 与 runtime 共用的上下文契约。
/*
Package corectx 抽象了框架中通用的上下文能力，包括：

  - 父 context、终止与完成信号；
  - panic 自动恢复与错误上报策略；
  - 等待组/屏障，用于协调 service 与 runtime 的关闭顺序；
  - 当前上下文和并发安全上下文提供器接口。

service.Context、runtime.Context，以及 entity/component 对当前上下文的暴露方式，
都依赖这些基础接口。
*/
package corectx
