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

// Code generated by eventc eventtab --name=entityComponentManagerEventTab; DO NOT EDIT.

package ec

import (
	event "git.golaxy.org/core/event"
)

type IEntityComponentManagerEventTab interface {
	EventComponentManagerAddComponents() event.IEvent
	EventComponentManagerRemoveComponent() event.IEvent
	EventComponentManagerFirstTouchComponent() event.IEvent
}

var (
	_entityComponentManagerEventTabId = event.DeclareEventTabIdT[entityComponentManagerEventTab]()
	EventComponentManagerAddComponentsId = _entityComponentManagerEventTabId + 0
	EventComponentManagerRemoveComponentId = _entityComponentManagerEventTabId + 1
	EventComponentManagerFirstTouchComponentId = _entityComponentManagerEventTabId + 2
)

type entityComponentManagerEventTab [3]event.Event

func (eventTab *entityComponentManagerEventTab) Init(autoRecover bool, reportError chan error, recursion event.EventRecursion) {
	(*eventTab)[0].Init(autoRecover, reportError, recursion)
	(*eventTab)[1].Init(autoRecover, reportError, recursion)
	(*eventTab)[2].Init(autoRecover, reportError, recursion)
}

func (eventTab *entityComponentManagerEventTab) Open() {
	for i := range *eventTab {
		(*eventTab)[i].Open()
	}
}

func (eventTab *entityComponentManagerEventTab) Close() {
	for i := range *eventTab {
		(*eventTab)[i].Close()
	}
}

func (eventTab *entityComponentManagerEventTab) Clean() {
	for i := range *eventTab {
		(*eventTab)[i].Clean()
	}
}

func (eventTab *entityComponentManagerEventTab) Ctrl() event.IEventCtrl {
	return eventTab
}

func (eventTab *entityComponentManagerEventTab) Event(id uint64) event.IEvent {
	if _entityComponentManagerEventTabId != id & 0xFFFFFFFF00000000 {
		return nil
	}
	pos := id & 0xFFFFFFFF
	if pos >= uint64(len(*eventTab)) {
		return nil
	}
	return &(*eventTab)[pos]
}

func (eventTab *entityComponentManagerEventTab) EventComponentManagerAddComponents() event.IEvent {
	return &(*eventTab)[0]
}

func (eventTab *entityComponentManagerEventTab) EventComponentManagerRemoveComponent() event.IEvent {
	return &(*eventTab)[1]
}

func (eventTab *entityComponentManagerEventTab) EventComponentManagerFirstTouchComponent() event.IEvent {
	return &(*eventTab)[2]
}