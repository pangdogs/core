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

// Package define 利用泛型特性，简化代码编写。
/*
   - 支持定义组件。
   - 支持定义插件与插件接口，按安装位置分类，共有通用插件、运行时插件、服务插件三种。
   - 使用 GoLand 作为 IDE 时，需要更新至 2023.2 版本以上，否则可能会有误报错。
   - 可以参考官方示例：https://git.golaxy.org/examples，学习如何使用。
*/
package define
