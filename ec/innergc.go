package ec

import "kit.golaxy.org/golaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}

type _InnerGCCollector interface {
	getInnerGCCollector() container.GCCollector
}
