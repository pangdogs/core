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

// Code generated by eventc eventtab --name=entityTreeNodeEventTab; DO NOT EDIT.

package ec

import (
	event "git.golaxy.org/core/event"
)

type IEntityTreeNodeEventTab interface {
	EventTreeNodeAddChild() event.IEvent
	EventTreeNodeRemoveChild() event.IEvent
	EventTreeNodeEnterParent() event.IEvent
	EventTreeNodeLeaveParent() event.IEvent
}

var (
	_entityTreeNodeEventTabId = event.DeclareEventTabIdT[entityTreeNodeEventTab]()
	EventTreeNodeAddChildId = _entityTreeNodeEventTabId + 0
	EventTreeNodeRemoveChildId = _entityTreeNodeEventTabId + 1
	EventTreeNodeEnterParentId = _entityTreeNodeEventTabId + 2
	EventTreeNodeLeaveParentId = _entityTreeNodeEventTabId + 3
)

type entityTreeNodeEventTab [4]event.Event

func (eventTab *entityTreeNodeEventTab) Init(autoRecover bool, reportError chan error, recursion event.EventRecursion) {
	(*eventTab)[0].Init(autoRecover, reportError, recursion)
	(*eventTab)[1].Init(autoRecover, reportError, recursion)
	(*eventTab)[2].Init(autoRecover, reportError, recursion)
	(*eventTab)[3].Init(autoRecover, reportError, recursion)
}

func (eventTab *entityTreeNodeEventTab) Enable() {
	for i := range *eventTab {
		(*eventTab)[i].Enable()
	}
}

func (eventTab *entityTreeNodeEventTab) Disable() {
	for i := range *eventTab {
		(*eventTab)[i].Disable()
	}
}

func (eventTab *entityTreeNodeEventTab) UnbindAll() {
	for i := range *eventTab {
		(*eventTab)[i].UnbindAll()
	}
}

func (eventTab *entityTreeNodeEventTab) Ctrl() event.IEventCtrl {
	return eventTab
}

func (eventTab *entityTreeNodeEventTab) Event(id uint64) event.IEvent {
	if _entityTreeNodeEventTabId != id & 0xFFFFFFFF00000000 {
		return nil
	}
	pos := id & 0xFFFFFFFF
	if pos >= uint64(len(*eventTab)) {
		return nil
	}
	return &(*eventTab)[pos]
}

func (eventTab *entityTreeNodeEventTab) EventTreeNodeAddChild() event.IEvent {
	return &(*eventTab)[0]
}

func (eventTab *entityTreeNodeEventTab) EventTreeNodeRemoveChild() event.IEvent {
	return &(*eventTab)[1]
}

func (eventTab *entityTreeNodeEventTab) EventTreeNodeEnterParent() event.IEvent {
	return &(*eventTab)[2]
}

func (eventTab *entityTreeNodeEventTab) EventTreeNodeLeaveParent() event.IEvent {
	return &(*eventTab)[3]
}
