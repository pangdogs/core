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
在这个框架中，我们有以下核心概念：
	- Entity（实体）：代表了应用程序中的一个实际对象。以游戏举例，可以是场景，也可以是玩家、NPC、物品等。实体本身只提供生命周期管理、组件管理两项核心能力，行为和属性将由不同的组件来提供。
	- Component（组件）：为实体提供行为与属性，不同的组件可以为实体提供不同的行为和状态。以游戏举例，为实体添加兔子外观组件、飞行组件、挖洞组件，那么我们将会得到一个即会飞行又会挖洞的拥有兔子外观的实体。
	- Context（上下文）：实体的运行环境，通常用于提供全局属性与方法。因为Golaxy主要用于服务器逻辑开发，需要支持多线程并发环境，所以提供了运行时上下文（runtime context）、服务上下文（service context）两个上下文，分别对应本地单线程环境与并发多线程环境。
EC组件框架基于组合模式，在大部分情况下优于继承模式，开发者可以更好地组织代码，实现高内聚和低耦合的代码架构。
*/
package ec
