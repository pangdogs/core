package runtime

type _ContextInnerGC struct {
	*ContextBehavior
}

func (igc *_ContextInnerGC) Init(cb *ContextBehavior) {
	igc.ContextBehavior = cb
}

func (igc *_ContextInnerGC) GC() {
	for i := range igc.ContextBehavior.gcList {
		igc.ContextBehavior.gcList[i].GC()
	}
	igc.ContextBehavior.gcList = igc.ContextBehavior.gcList[:0]
}

func (igc *_ContextInnerGC) NeedGC() bool {
	return len(igc.ContextBehavior.gcList) > 0
}
