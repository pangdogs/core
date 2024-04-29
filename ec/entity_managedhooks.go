package ec

import (
	"git.golaxy.org/core/event"
	"slices"
)

// ManagedHooks 托管hook，在实体销毁时自动解绑定
func (entity *EntityBehavior) ManagedHooks(hooks ...event.Hook) {
	entity.managedHooks = slices.DeleteFunc(entity.managedHooks, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})
	entity.managedHooks = append(entity.managedHooks, hooks...)
}

func (entity *EntityBehavior) cleanManagedHooks() {
	event.Clean(entity.managedHooks)
	entity.managedHooks = nil
}
