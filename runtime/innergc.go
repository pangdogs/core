package runtime

import "github.com/pangdogs/galaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}
