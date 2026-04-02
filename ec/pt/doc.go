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

// Package pt 提供实体与组件原型库。
/*
Package pt 用于声明、查询和构造实体原型与组件原型。

核心对象包括：

  - EntityDescriptor / ComponentDescriptor：原型声明描述；
  - EntityLib / ComponentLib：原型注册表与订阅接口；
  - EntityPT / ComponentPT：供 ec 与 service/runtime 使用的原型对象。

服务启动阶段通常会通过 service.Context.EntityLib() 或根包的 BuildEntityPT 声明
实体原型；运行时创建实体时，再根据原型生成实体与内建组件。
*/
package pt
