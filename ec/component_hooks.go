package ec

import (
	"git.golaxy.org/core/event"
	"slices"
)

// ManagedHooks 托管hook，在组件销毁时自动解绑定
func (comp *ComponentBehavior) ManagedHooks(hooks ...event.Hook) {
	comp.hooks = slices.DeleteFunc(comp.hooks, func(hook event.Hook) bool {
		return !hook.IsBound()
	})
}

func (comp *ComponentBehavior) cleanHooks() {
	for i := range comp.hooks {
		comp.hooks[i].Unbind()
	}
	comp.hooks = nil
}
