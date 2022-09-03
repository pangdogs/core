package core

import "github.com/pangdogs/galaxy/core/container"

type _InnerGC interface {
	getGC() container.GC
}

type _InnerGCCollector interface {
	getGCCollector() container.GCCollector
}
