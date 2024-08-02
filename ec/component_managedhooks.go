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

// ManagedHooks 托管hook，在组件销毁时自动解绑定
func (comp *ComponentBehavior) ManagedHooks(hooks ...event.Hook) {
	comp.managedHooks = slices.DeleteFunc(comp.managedHooks, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})
	comp.managedHooks = append(comp.managedHooks, hooks...)
}

func (comp *ComponentBehavior) cleanManagedHooks() {
	event.Clean(comp.managedHooks)
	comp.managedHooks = nil
}
