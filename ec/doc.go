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

// Package ec 提供了一个EC（Entity-Component）组件框架，用于帮助开发者组织代码架构。
/*
在框架中，有以下几个核心概念：
	- Entity（实体）：代表应用程序中的一个实际对象。以游戏为例，可以是场景、玩家、NPC或物品等。一般情况下，实体只提供生命周期和组件管理两项核心能力，行为和属性将由不同的组件提供。
	- Component（组件）：为实体提供行为和属性，不同的组件可以为实体提供不同的行为和状态。以游戏为例，我们为实体添加兔子外观组件、飞行组件和挖洞组件，那么将会得到一个既能飞行又能挖洞并且拥有兔子外观的实体。
	- Context（上下文）：实体的运行环境，通常用于提供全局属性和方法。因为 Golaxy 分布式服务开发框架主要用于服务器逻辑开发，需要支持多线程并发环境，所以提供了运行时上下文（Runtime Context）和服务上下文（Service Context）两个上下文，分别对应本地单线程环境和并发多线程环境。
EC 组件框架基于组合模式，在大多数情况下优于继承模式，开发者可以更好地组织代码，实现高内聚和低耦合的代码架构。
*/
package ec
