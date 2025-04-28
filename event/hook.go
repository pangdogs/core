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

package event

import (
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
)

// Hook 事件钩子，由BindEvent()创建并返回的绑定句柄，请勿自己创建
type Hook struct {
	subscriberFace iface.FaceAny
	priority       int32
	at             *generic.Node[Hook]
	received       int32
}

// Unbind 解绑定事件与订阅者
func (hook *Hook) Unbind() {
	if hook.at != nil {
		hook.at.Escape()
		hook.at = nil
	}
}

// IsBound 是否已绑定事件
func (hook *Hook) IsBound() bool {
	return hook.at != nil && !hook.at.Escaped()
}

// UnbindHooks 解绑定事件钩子（Hook）
func UnbindHooks(hooks []Hook) {
	for i := range hooks {
		hooks[i].Unbind()
	}
}
