package ec

import "github.com/galaxy-kit/galaxy-go/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}

type _InnerGCCollector interface {
	getInnerGCCollector() container.GCCollector
}
