package runtime

import (
	"fmt"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// Current 获取当前运行时上下文
func Current(provider gctx.CurrentContextProvider) Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetCurrentContext())
}

func getRuntimeContext(provider gctx.CurrentContextProvider) Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetCurrentContext())
}
