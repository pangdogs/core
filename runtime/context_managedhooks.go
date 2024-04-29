package runtime

import (
	"git.golaxy.org/core/event"
	"slices"
)

// ManagedHooks 托管hook，在运行时停止时自动解绑定
func (ctx *ContextBehavior) ManagedHooks(hooks ...event.Hook) {
	ctx.managedHooks = slices.DeleteFunc(ctx.managedHooks, func(hook event.Hook) bool {
		return !hook.IsBound() || slices.Contains(hooks, hook)
	})
	ctx.managedHooks = append(ctx.managedHooks, hooks...)
}

func (ctx *ContextBehavior) cleanManagedHooks() {
	event.Clean(ctx.managedHooks)
	ctx.managedHooks = nil
}
