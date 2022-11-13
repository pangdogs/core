package runtime

import "github.com/galaxy-kit/galaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}
