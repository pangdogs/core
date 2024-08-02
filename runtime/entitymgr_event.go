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

//go:generate go run git.golaxy.org/core/event/eventc event
package runtime

import "git.golaxy.org/core/ec"

// EventEntityMgrAddEntity 事件：实体管理器添加实体
// +event-gen:export=0
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(entityMgr EntityMgr, entity ec.Entity)
}

// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
// +event-gen:export=0
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(entityMgr EntityMgr, entity ec.Entity)
}

// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
// +event-gen:export=0
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(entityMgr EntityMgr, entity ec.Entity, components []ec.Component)
}

// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
// +event-gen:export=0
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(entityMgr EntityMgr, entity ec.Entity, component ec.Component)
}

// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
// +event-gen:export=0
type EventEntityMgrEntityFirstAccessComponent interface {
	OnEntityMgrEntityFirstAccessComponent(entityMgr EntityMgr, entity ec.Entity, component ec.Component)
}
