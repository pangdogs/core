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

package runtime

import (
	"git.golaxy.org/core/event"
	"slices"
)

// ManagedAddHooks 托管事件钩子（event.Hook），在运行时停止时自动解绑定
func (ctx *ContextBehavior) ManagedAddHooks(hooks ...event.Hook) {
	ctx.managedHooks = slices.DeleteFunc(ctx.managedHooks, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})
	ctx.managedHooks = append(ctx.managedHooks, hooks...)
}

// ManagedAddTagHooks 根据标签托管事件钩子（event.Hook），在运行时停止时自动解绑定
func (ctx *ContextBehavior) ManagedAddTagHooks(tag string, hooks ...event.Hook) {
	exists, ok := ctx.managedTagHooks.Get(tag)
	if !ok {
		ctx.managedTagHooks.Add(tag, hooks)
		return
	}

	exists = slices.DeleteFunc(exists, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})

	managedHooks := append(exists, hooks...)

	if len(managedHooks) <= 0 {
		ctx.managedTagHooks.Delete(tag)
		return
	}

	ctx.managedTagHooks.Add(tag, managedHooks)
}

// ManagedGetTagHooks 根据标签获取托管事件钩子（event.Hook）
func (ctx *ContextBehavior) ManagedGetTagHooks(tag string) []event.Hook {
	hooks, _ := ctx.managedTagHooks.Get(tag)
	return hooks
}

// ManagedCleanTagHooks 清理根据标签托管的事件钩子（event.Hook）
func (ctx *ContextBehavior) ManagedCleanTagHooks(tag string) {
	idx := ctx.managedTagHooks.Index(tag)
	if idx < 0 {
		return
	}
	event.Clean(ctx.managedTagHooks[idx].V)
	ctx.managedTagHooks = slices.Delete(ctx.managedTagHooks, idx, idx+1)
}

func (ctx *ContextBehavior) managedCleanAllHooks() {
	event.Clean(ctx.managedHooks)
	ctx.managedHooks = nil

	ctx.managedTagHooks.Each(func(tag string, hooks []event.Hook) { event.Clean(hooks) })
	ctx.managedTagHooks = nil
}
