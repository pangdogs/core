package runtime

func (ctx *ContextBehavior) gc() {
	for i := range ctx.gcList {
		ctx.gcList[i].GC()
	}
	ctx.gcList = ctx.gcList[:0]
}
