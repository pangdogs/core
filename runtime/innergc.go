package runtime

import "github.com/golaxy-kit/golaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}
