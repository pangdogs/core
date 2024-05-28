package event

import (
	"git.golaxy.org/core/utils/container"
	"git.golaxy.org/core/utils/iface"
)

// Hook 事件钩子，主要用于重新绑定或解除绑定事件，由BindEvent()创建并返回，请勿自己创建
type Hook struct {
	subscriberFace iface.FaceAny
	priority       int32
	element        *container.Element[Hook]
	received       int32
}

// Unbind 解绑定事件与订阅者
func (hook *Hook) Unbind() {
	if hook.element != nil {
		hook.element.Escape()
		hook.element = nil
	}
}

// IsBound 是否已绑定事件
func (hook *Hook) IsBound() bool {
	return hook.element != nil && !hook.element.Escaped()
}

// Clean 清理Hooks
func Clean(hooks []Hook) {
	for i := range hooks {
		hooks[i].Unbind()
	}
}
