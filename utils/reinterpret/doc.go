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

// Package reinterpret 提供基于实例缓存的接口重新解释。
/*
Package reinterpret 约定 InstanceProvider 暴露统一的实例缓存，然后通过 Cast 把同一个
对象重新解释为其他接口类型。

这个机制常用于扩展实体、组件、service 或 runtime 的对外接口，而无需在调用点做
重复的类型断言或显式保存多个接口引用。
*/
package reinterpret
