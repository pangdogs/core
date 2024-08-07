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

// Package core Golaxy分布式服务开发框架的内核，主要提供Actor线程模型（Actor Model）和EC组件框架（Entity-Component）。
/*
   - 使用EC组件框架（Entity-Component）组织代码结构。
   - 并发模式基于Actor编程模型，实体（Entity）就是Actor，其中组件（Component）用于实现状态（state）与行为（behavior），运行时（Runtime）中的任务处理流水线就是邮箱（mailbox），
     实体的Id就是邮箱地址(mailbox address)，服务上下文（Service context）提供的全局实体管理功能，可以用于投递邮件（mail）给Actor。不同于传统Actor编程模型的地方是
     多个Actor可以使用同个邮箱，也就是多个实体可以加入同个运行时，在同个线程中运行。
   - 一些分布式系统常用的依赖项，例如服务发现（Service Discovery）、消息队列与事件驱动架构（MQ and Broker）、分布式锁（Distributed Sync）等分布式服务特性，这些特性将会以官方插件形式提供。
*/
package core
