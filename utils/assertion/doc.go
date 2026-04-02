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

// Package assertion 提供基于实体组件模型的断言与注入辅助。
/*
Package assertion 通过反射从 ec.Entity 中提取组件，并按字段类型、组件名或原型名
注入到结构体中。

它适合把多个组件组合成临时视图，或者在组件启动时把依赖字段补齐。字段可通过
`ec:"name,prototype"` tag 指定组件名或组件原型；如果目标原型已在服务的组件库中
注册，但实体上尚未存在对应组件，Inject 还可以按原型动态创建组件。

这类操作依赖反射，适合启动期、装配期或测试场景，不建议在高频热点路径中反复执行。
*/
package assertion
