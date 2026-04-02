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

// Package async 提供 Future 与异步结果类型。
/*
Package async 定义了 Result、Future、FutureVoid 以及一组 Return/Yield 辅助函数，
用于在 service/runtime 异步调用和独立 goroutine 任务之间传递结果。

Future 支持：

  - Wait：等待单个结果；
  - Chan：消费 yield 式多次产出；
  - Context：把 future 完成态转换为 context 取消信号。

根包 core 的 CallAsync、GoAsync、Await 等 API 都建立在该包之上。
*/
package async
