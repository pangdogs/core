package container

// GC GC接口
type GC interface {
	// GC GC
	GC()
	// NeedGC 是否需要GC
	NeedGC() bool
}

// GCCollector GC收集器接口
type GCCollector interface {
	// CollectGC 收集GC
	CollectGC(gc GC)
}
