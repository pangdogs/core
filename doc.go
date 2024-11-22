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

// Package core 是 Golaxy 分布式服务开发框架的内核，主要提供 Actor 线程模型（Actor Model）和 EC 组件框架（Entity-Component）。
/*
   - 使用 EC 组件框架（Entity-Component）来组织代码结构。
   - 并发模式基于 Actor 线程模型，实体（Entity）就是 Actor，组件（Component）用于实现状态（State）和行为（Behavior）。运行时（Runtime）中的任务处理流水线相当于邮箱（Mailbox），
     实体的 Id 就是邮箱地址（Mailbox Address）。服务上下文（Service Context）提供全局实体管理功能，可以用于向 Actor 投递邮件（Mail）。与传统 Actor 线程模型不同的是，
     多个 Actor 可以共享同一个邮箱，即多个实体可以加入同一个运行时，在同一个线程中运行。
   - 一些分布式系统常用的依赖项，例如服务发现（Service Discovery）、消息队列与事件驱动架构（MQ And Broker）、分布式锁（Distributed Sync）等分布式服务特性，将以官方插件形式提供。
*/
package core
