package core

import "time"

// Run ...
func (runtime *_RuntimeBehavior) Run() <-chan struct{} {
	if !runtime.ctx.markRunning() {
		panic("runtime already running")
	}

	shutChan := make(chan struct{}, 1)

	runtime.ctx.setFrame(runtime.opts.Frame)
	runtime.processQueue = make(chan func(), runtime.opts.ProcessQueueCapacity)
	runtime.ctx.setCallee(runtime)

	go runtime.running(shutChan)

	return shutChan
}

// Stop ...
func (runtime *_RuntimeBehavior) Stop() {
	runtime.ctx.GetCancelFunc()()
}

func (runtime *_RuntimeBehavior) running(shutChan chan struct{}) {
	if parentCtx, ok := runtime.ctx.GetParentCtx().(Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	hooks := runtime.loopStarted()

	defer func() {
		runtime.loopStopped(hooks)

		if parentCtx, ok := runtime.ctx.GetParentCtx().(Context); ok {
			parentCtx.GetWaitGroup().Done()
		}

		runtime.ctx.GetWaitGroup().Wait()

		runtime.ctx.markShutdown()
		shutChan <- struct{}{}
	}()

	frame := runtime.opts.Frame

	if frame == nil {
		defer runtime.loopNoFrameEnd()
		runtime.loopNoFrame()

	} else if frame.Blink() {
		defer runtime.loopWithBlinkFrameEnd()
		runtime.loopWithBlinkFrame()

	} else {
		defer runtime.loopWithFrameEnd()
		runtime.loopWithFrame()
	}
}

func (runtime *_RuntimeBehavior) loopStarted() (hooks [4]Hook) {
	runtimeCtx := runtime.ctx
	frame := runtime.opts.Frame

	if frame != nil {
		frame.runningBegin()
	}

	hooks[0] = BindEvent[EventEntityMgrAddEntity[RuntimeContext]](runtimeCtx.EventEntityMgrAddEntity(), runtime)
	hooks[1] = BindEvent[EventEntityMgrRemoveEntity[RuntimeContext]](runtimeCtx.EventEntityMgrRemoveEntity(), runtime)
	hooks[2] = BindEvent[EventEntityMgrEntityAddComponents[RuntimeContext]](runtimeCtx.EventEntityMgrEntityAddComponents(), runtime)
	hooks[3] = BindEvent[EventEntityMgrEntityRemoveComponent[RuntimeContext]](runtimeCtx.EventEntityMgrEntityRemoveComponent(), runtime)

	runtimeCtx.RangeEntities(func(entity Entity) bool {
		CallOuterNoRet(runtime.opts.EnableAutoRecover, runtimeCtx.GetReportError(), func() {
			runtime.OnEntityMgrAddEntity(runtimeCtx, entity)
		})
		return true
	})

	CallOuterNoRet(runtime.opts.EnableAutoRecover, runtimeCtx.GetReportError(), func() {
		if runtimeCtx.getOptions().StartedCallback != nil {
			runtimeCtx.getOptions().StartedCallback(runtime.opts.Inheritor.IFace)
		}
	})

	return
}

func (runtime *_RuntimeBehavior) loopStopped(hooks [4]Hook) {
	runtimeCtx := runtime.ctx
	frame := runtime.opts.Frame

	CallOuterNoRet(runtime.opts.EnableAutoRecover, runtimeCtx.GetReportError(), func() {
		if runtimeCtx.getOptions().StoppedCallback != nil {
			runtimeCtx.getOptions().StoppedCallback(runtime.opts.Inheritor.IFace)
		}
	})

	runtimeCtx.ReverseRangeEntities(func(entity Entity) bool {
		CallOuterNoRet(runtime.opts.EnableAutoRecover, runtimeCtx.GetReportError(), func() {
			runtime.OnEntityMgrRemoveEntity(runtimeCtx, entity)
		})
		return true
	})

	for i := range hooks {
		hooks[i].Unbind()
	}

	if frame != nil {
		frame.runningEnd()
	}
}

func (runtime *_RuntimeBehavior) loopNoFrame() {
	gcTicker := time.NewTicker(runtime.opts.GCInterval)
	defer gcTicker.Stop()

	for {
		select {
		case process, ok := <-runtime.processQueue:
			if !ok {
				return
			}
			CallOuterNoRet(runtime.opts.EnableAutoRecover, runtime.ctx.GetReportError(), process)

		case <-gcTicker.C:
			runtime.opts.Inheritor.IFace.GC()

		case <-runtime.ctx.Done():
			return
		}
	}
}

func (runtime *_RuntimeBehavior) loopNoFrameEnd() {
	close(runtime.processQueue)

	for {
		select {
		case process, ok := <-runtime.processQueue:
			if !ok {
				return
			}
			CallOuterNoRet(runtime.opts.EnableAutoRecover, runtime.ctx.GetReportError(), process)

		default:
			return
		}
	}
}

func (runtime *_RuntimeBehavior) loopWithFrame() {
	frame := runtime.opts.Frame

	go func() {
		updateTicker := time.NewTicker(time.Duration(float64(time.Second) / float64(frame.GetTargetFPS())))
		defer updateTicker.Stop()

		totalFrames := frame.GetTotalFrames()

		for curFrames := uint64(1); ; {
			if totalFrames > 0 && curFrames >= totalFrames {
				return
			}

			select {
			case <-updateTicker.C:
				CallOuterNoRet(runtime.opts.EnableAutoRecover, runtime.ctx.GetReportError(), func() {
					timeoutTimer := time.NewTimer(runtime.opts.ProcessQueueTimeout)
					defer timeoutTimer.Stop()

					select {
					case runtime.processQueue <- runtime.frameUpdate:
						curFrames++
						return
					case <-timeoutTimer.C:
						panic("process queue push frame update timeout")
					}
				})

			case <-runtime.ctx.Done():
				return
			}
		}
	}()

	frame.setCurFrames(0)
	runtime.firstFrameUpdate()

	gcTicker := time.NewTicker(runtime.opts.GCInterval)
	defer gcTicker.Stop()

	for {
		select {
		case process, ok := <-runtime.processQueue:
			if !ok {
				return
			}
			CallOuterNoRet(runtime.opts.EnableAutoRecover, runtime.ctx.GetReportError(), process)

		case <-gcTicker.C:
			runtime.opts.Inheritor.IFace.GC()

		case <-runtime.ctx.Done():
			return
		}
	}
}

func (runtime *_RuntimeBehavior) loopWithFrameEnd() {
	frame := runtime.opts.Frame

	close(runtime.processQueue)

	func() {
		for {
			select {
			case process, ok := <-runtime.processQueue:
				if !ok {
					return
				}
				CallOuterNoRet(runtime.opts.EnableAutoRecover, runtime.ctx.GetReportError(), process)

			default:
				return
			}
		}
	}()

	frame.frameEnd()
	frame.setCurFrames(frame.GetCurFrames() + 1)
}

func (runtime *_RuntimeBehavior) frameUpdate() {
	frame := runtime.opts.Frame

	frame.frameEnd()
	frame.setCurFrames(frame.GetCurFrames() + 1)

	runtime.firstFrameUpdate()
}

func (runtime *_RuntimeBehavior) firstFrameUpdate() {
	frame := runtime.opts.Frame

	frame.frameBegin()

	frame.updateBegin()
	defer frame.updateEnd()

	emitEventUpdate(&runtime.eventUpdate)
	emitEventLateUpdate(&runtime.eventLateUpdate)
}

func (runtime *_RuntimeBehavior) loopWithBlinkFrame() {
	frame := runtime.opts.Frame
	totalFrames := frame.GetTotalFrames()

	gcFrames := uint64(runtime.opts.GCInterval.Seconds() * float64(frame.GetTargetFPS()))

	for frame.setCurFrames(0); ; {
		curFrames := frame.GetCurFrames()

		if totalFrames > 0 && curFrames >= totalFrames {
			return
		}

		if !runtime.blinkFrameUpdate() {
			return
		}

		if curFrames%gcFrames == 0 {
			runtime.opts.Inheritor.IFace.GC()
		}

		frame.setCurFrames(curFrames + 1)
	}
}

func (runtime *_RuntimeBehavior) loopWithBlinkFrameEnd() {
	close(runtime.processQueue)

	for {
		select {
		case process, ok := <-runtime.processQueue:
			if !ok {
				return
			}
			CallOuterNoRet(runtime.opts.EnableAutoRecover, runtime.ctx.GetReportError(), process)

		default:
			break
		}
	}
}

func (runtime *_RuntimeBehavior) blinkFrameUpdate() bool {
	frame := runtime.opts.Frame

	frame.frameBegin()
	defer frame.frameEnd()

	for {
		select {
		case <-runtime.ctx.Done():
			return false

		default:
			func() {
				frame.updateBegin()
				defer frame.updateEnd()

				emitEventUpdate(&runtime.eventUpdate)
				emitEventLateUpdate(&runtime.eventLateUpdate)
			}()
			return true
		}
	}

	return true
}
