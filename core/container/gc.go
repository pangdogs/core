package container

// GC GC接口
type GC interface {
	GC()
	NeedGC() bool
}

// GCCollector GC收集器接口
type GCCollector interface {
	CollectGC(gc GC)
}
