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

// Package extension 插件系统，用于扩展服务或运行时能力，例如服务发现、消息队列与日志系统等。
/*
   - 插件主要以组合方式安装在服务或运行时上下文上，用于扩展服务或运行时能力。
   - 服务与运行时上下文均支持安装插件，注意服务上的插件需要支持多线程环境，运行时上的插件仅需支持单线程环境即可。
*/
package extension
