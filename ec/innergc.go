package ec

import "github.com/pangdogs/galaxy/util/container"

type _InnerGC interface {
	getInnerGC() container.GC
}

type _InnerGCCollector interface {
	getInnerGCCollector() container.GCCollector
}
