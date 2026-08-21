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

// Package async 提供一次性结果、完成信号、连续流和结构化异步作用域。
/*
Package async 将不同异步语义拆分为独立类型：

  - Promise/Future：非泛型、一次性、可重放的 Result；
  - Completer/Signal：不携带 Result 的生命周期完成通知；
  - Emitter/Stream：连续 Result 的单消费流；
  - Scope/Spawn：绑定宿主生命周期的后台任务取消、汇合与统计；
  - Race、FirstSuccess、All、AllSettled、Zip2、Map、FlatMap 和 Timeout：
    基于完成订阅的 Future 组合器。

Future 内部保存完成结果，并通过 OnComplete 在完成者 goroutine 中直接通知订阅者；
多个 Wait、TryGet 或 OnComplete 消费者读取的是同一个可重放结果，不会竞争消费，
也不会为每个订阅者启动等待 goroutine。

Stream 采用单消费语义；多个消费者会竞争元素。需要广播时应使用 event 或上层消息
设施。Scope.Close 默认以 context.Canceled 取消 Context，也可携带指定原因；Scope 关闭后
以 ErrScopeClosed 拒绝新任务。nil Scope 同样视为关闭，其 Context 以 ErrScopeClosed 为
原因取消且 Completion 已完成。Scope 不会等待或强制终止 goroutine，需要汇合时应等待
Scope.Completion 返回的 Signal。

根包 core 的 Submit、Post、Spawn 和 ContinueOn 建立在这些能力之上。
*/
package async
