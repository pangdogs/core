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

// Package option 提供泛型选项构造器。
/*
Package option 定义了 `Setting[T]` 及其组合函数，用来实现本仓库大量使用的
`With.xxx(...)` 风格配置 API。

New 会先应用默认值，再按顺序应用额外配置；Append 和 Change 用于在现有值上继续叠加。
service、runtime、core 与 ec 包的选项体系都基于该包实现。
*/
package option
