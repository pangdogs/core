package ec

import "git.golaxy.org/core/event"

// ManagedHooks 托管hook，在组件销毁时自动解绑定
func (comp *ComponentBehavior) ManagedHooks(hooks ...event.Hook) {
	for i := len(comp.hooks) - 1; i >= 0; i-- {
		if !comp.hooks[i].IsBound() {
			comp.hooks = append(comp.hooks[:i], comp.hooks[i+1:]...)
		}
	}
	comp.hooks = append(comp.hooks, hooks...)
}

func (comp *ComponentBehavior) cleanHooks() {
	for i := range comp.hooks {
		comp.hooks[i].Unbind()
	}
	comp.hooks = nil
}
