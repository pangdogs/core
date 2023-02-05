package runtime

import "kit.golaxy.org/golaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}
