package container

type GC interface {
	GC()
	NeedGC() bool
}

type GCCollector interface {
	CollectGC(gc GC)
}
