package ec

import (
	"git.golaxy.org/core/event"
	"slices"
)

// ManagedHooks 托管hook，在实体销毁时自动解绑定
func (entity *EntityBehavior) ManagedHooks(hooks ...event.Hook) {
	entity.hooks = slices.DeleteFunc(entity.hooks, func(hook event.Hook) bool {
		return !hook.IsBound()
	})
}

func (entity *EntityBehavior) cleanHooks() {
	for i := range entity.hooks {
		entity.hooks[i].Unbind()
	}
	entity.hooks = nil
}
