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

// ManagedAddHooks 托管事件钩子（event.Hook），在组件销毁时自动解绑定
func (comp *ComponentBehavior) ManagedAddHooks(hooks ...event.Hook) {
	comp.managedHooks = slices.DeleteFunc(comp.managedHooks, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})
	comp.managedHooks = append(comp.managedHooks, hooks...)
}

// ManagedAddTagHooks 根据标签托管事件钩子（event.Hook），在组件销毁时自动解绑定
func (comp *ComponentBehavior) ManagedAddTagHooks(tag string, hooks ...event.Hook) {
	exists, ok := comp.managedTagHooks.Get(tag)
	if !ok {
		comp.managedTagHooks.Add(tag, hooks)
		return
	}

	exists = slices.DeleteFunc(exists, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})

	managedHooks := append(exists, hooks...)

	if len(managedHooks) <= 0 {
		comp.managedTagHooks.Delete(tag)
		return
	}

	comp.managedTagHooks.Add(tag, managedHooks)
}

// ManagedGetTagHooks 根据标签获取托管事件钩子（event.Hook）
func (comp *ComponentBehavior) ManagedGetTagHooks(tag string) []event.Hook {
	hooks, _ := comp.managedTagHooks.Get(tag)
	return hooks
}

// ManagedUnbindTagHooks 根据标签解绑定托管的事件钩子（event.Hook）
func (comp *ComponentBehavior) ManagedUnbindTagHooks(tag string) {
	idx, ok := comp.managedTagHooks.Index(tag)
	if !ok {
		return
	}
	event.UnbindHooks(comp.managedTagHooks[idx].V)
	comp.managedTagHooks = slices.Delete(comp.managedTagHooks, idx, idx+1)
}

func (comp *ComponentBehavior) managedUnbindAllHooks() {
	event.UnbindHooks(comp.managedHooks)
	comp.managedHooks = nil

	comp.managedTagHooks.Each(func(tag string, hooks []event.Hook) { event.UnbindHooks(hooks) })
	comp.managedTagHooks = nil
}
