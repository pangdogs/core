package ec

import "kit.golaxy.org/golaxy/event"

// AutoHooks 保存绑定事件的hook，在实体销毁时自动解绑定
func (entity *EntityBehavior) AutoHooks(hooks ...event.Hook) {
	for i := len(entity.hooks) - 1; i >= 0; i-- {
		if !entity.hooks[i].IsBound() {
			entity.hooks = append(entity.hooks[:i], entity.hooks[i+1:]...)
		}
	}
	entity.hooks = append(entity.hooks, hooks...)
}

func (entity *EntityBehavior) cleanHooks() {
	for i := range entity.hooks {
		entity.hooks[i].Unbind()
	}
	entity.hooks = nil
}
