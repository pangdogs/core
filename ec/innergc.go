package ec

import "github.com/golaxy-kit/golaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}

type _InnerGCCollector interface {
	getInnerGCCollector() container.GCCollector
}
