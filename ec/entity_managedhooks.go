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

package ec

import (
	"git.golaxy.org/core/event"
	"slices"
)

// ManagedAddHooks 托管事件钩子（event.Hook），在实体销毁时自动解绑定
func (entity *EntityBehavior) ManagedAddHooks(hooks ...event.Hook) {
	entity.managedHooks = slices.DeleteFunc(entity.managedHooks, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})
	entity.managedHooks = append(entity.managedHooks, hooks...)
}

// ManagedAddTagHooks 根据标签托管事件钩子（event.Hook），在实体销毁时自动解绑定
func (entity *EntityBehavior) ManagedAddTagHooks(tag string, hooks ...event.Hook) {
	exists, ok := entity.managedTagHooks.Get(tag)
	if !ok {
		entity.managedTagHooks.Add(tag, hooks)
		return
	}

	exists = slices.DeleteFunc(exists, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})

	managedHooks := append(exists, hooks...)

	if len(managedHooks) <= 0 {
		entity.managedTagHooks.Delete(tag)
		return
	}

	entity.managedTagHooks.Add(tag, managedHooks)
}

// ManagedGetTagHooks 根据标签获取托管事件钩子（event.Hook）
func (entity *EntityBehavior) ManagedGetTagHooks(tag string) []event.Hook {
	hooks, _ := entity.managedTagHooks.Get(tag)
	return hooks
}

// ManagedCleanTagHooks 清理根据标签托管的事件钩子（event.Hook）
func (entity *EntityBehavior) ManagedCleanTagHooks(tag string) {
	idx := entity.managedTagHooks.Index(tag)
	if idx < 0 {
		return
	}
	event.Clean(entity.managedTagHooks[idx].V)
	entity.managedTagHooks = slices.Delete(entity.managedTagHooks, idx, idx+1)
}

func (entity *EntityBehavior) managedCleanAllHooks() {
	event.Clean(entity.managedHooks)
	entity.managedHooks = nil

	entity.managedTagHooks.Each(func(tag string, hooks []event.Hook) { event.Clean(hooks) })
	entity.managedTagHooks = nil
}
