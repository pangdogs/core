package galaxy

import (
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/internal"
	"github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/runtime"
	"github.com/golaxy-kit/golaxy/util"
	"time"
)

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (_runtime *RuntimeBehavior) Run() <-chan struct{} {
	if !runtime.UnsafeContext(_runtime.ctx).MarkRunning() {
		panic("_runtime already running")
	}

	shutChan := make(chan struct{}, 1)

	runtime.UnsafeContext(_runtime.ctx).SetFrame(_runtime.opts.Frame)
	_runtime.processQueue = make(chan func(), _runtime.opts.ProcessQueueCapacity)
	runtime.UnsafeContext(_runtime.ctx).SetCallee(_runtime)

	if parentCtx, ok := _runtime.ctx.GetParentContext().(internal.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go _runtime.running(shutChan)

	return shutChan
}

// Stop 停止
func (_runtime *RuntimeBehavior) Stop() {
	_runtime.ctx.GetCancelFunc()()
}

func (_runtime *RuntimeBehavior) running(shutChan chan struct{}) {
	if pluginBundle := runtime.UnsafeContext(_runtime.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.Range(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginInit, ok := pluginFace.Iface.(_RuntimePluginInit); ok {
				internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
					pluginInit.Init(_runtime.ctx)
				})
			}
			return true
		})
	}

	hooks := _runtime.loopStarted()

	defer func() {
		if callback := runtime.UnsafeContext(_runtime.ctx).GetOptions().StoppingCallback; callback != nil {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				callback(_runtime.ctx)
			})
		}

		_runtime.loopStopped(hooks)

		_runtime.ctx.GetWaitGroup().Wait()

		if callback := runtime.UnsafeContext(_runtime.ctx).GetOptions().StoppedCallback; callback != nil {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				callback(_runtime.ctx)
			})
		}

		if parentCtx, ok := _runtime.ctx.GetParentContext().(internal.Context); ok {
			parentCtx.GetWaitGroup().Done()
		}

		if pluginBundle := runtime.UnsafeContext(_runtime.ctx).GetOptions().PluginBundle; pluginBundle != nil {
			pluginBundle.Range(func(pluginName string, pluginFace util.FaceAny) bool {
				if pluginShut, ok := pluginFace.Iface.(_PluginShut); ok {
					internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
						pluginShut.Shut()
					})
				}
				return true
			})
		}

		runtime.UnsafeContext(_runtime.ctx).MarkShutdown()
		shutChan <- struct{}{}
	}()

	frame := _runtime.opts.Frame

	if frame == nil {
		defer _runtime.loopNoFrameEnd()
		_runtime.loopNoFrame()

	} else if frame.Blink() {
		defer _runtime.loopWithBlinkFrameEnd()
		_runtime.loopWithBlinkFrame()

	} else {
		defer _runtime.loopWithFrameEnd()
		_runtime.loopWithFrame()
	}
}

func (_runtime *RuntimeBehavior) loopStarted() (hooks [5]localevent.Hook) {
	runtimeCtx := _runtime.ctx
	frame := _runtime.opts.Frame

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningBegin()
	}

	hooks[0] = localevent.BindEvent[runtime.EventEntityMgrAddEntity](runtimeCtx.GetEntityMgr().EventEntityMgrAddEntity(), _runtime)
	hooks[1] = localevent.BindEvent[runtime.EventEntityMgrRemoveEntity](runtimeCtx.GetEntityMgr().EventEntityMgrRemoveEntity(), _runtime)
	hooks[2] = localevent.BindEvent[runtime.EventEntityMgrEntityAddComponents](runtimeCtx.GetEntityMgr().EventEntityMgrEntityAddComponents(), _runtime)
	hooks[3] = localevent.BindEvent[runtime.EventEntityMgrEntityRemoveComponent](runtimeCtx.GetEntityMgr().EventEntityMgrEntityRemoveComponent(), _runtime)
	hooks[4] = localevent.BindEvent[runtime.EventEntityMgrEntityFirstAccessComponent](runtimeCtx.GetEntityMgr().EventEntityMgrEntityFirstAccessComponent(), _runtime)

	runtimeCtx.GetEntityMgr().RangeEntities(func(entity ec.Entity) bool {
		internal.CallOuterNoRet(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), func() {
			_runtime.OnEntityMgrAddEntity(runtimeCtx.GetEntityMgr(), entity)
		})
		return true
	})

	if callback := runtime.UnsafeContext(runtimeCtx).GetOptions().StartedCallback; callback != nil {
		internal.CallOuterNoRet(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), func() {
			callback(runtimeCtx)
		})
	}

	return
}

func (_runtime *RuntimeBehavior) loopStopped(hooks [5]localevent.Hook) {
	runtimeCtx := _runtime.ctx
	frame := _runtime.opts.Frame

	runtimeCtx.GetEntityMgr().ReverseRangeEntities(func(entity ec.Entity) bool {
		internal.CallOuterNoRet(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), func() {
			_runtime.OnEntityMgrRemoveEntity(runtimeCtx.GetEntityMgr(), entity)
		})
		return true
	})

	for i := range hooks {
		hooks[i].Unbind()
	}

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningEnd()
	}
}

func (_runtime *RuntimeBehavior) loopNoFrame() {
	gcTicker := time.NewTicker(_runtime.opts.GCInterval)
	defer gcTicker.Stop()

	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), process)

		case <-gcTicker.C:
			_runtime.gc()

		case <-_runtime.ctx.Done():
			return
		}
	}
}

func (_runtime *RuntimeBehavior) loopNoFrameEnd() {
	close(_runtime.processQueue)

	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), process)

		default:
			return
		}
	}
}

func (_runtime *RuntimeBehavior) loopWithFrame() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	go func() {
		updateTicker := time.NewTicker(time.Duration(float64(time.Second) / float64(frame.GetTargetFPS())))
		defer updateTicker.Stop()

		totalFrames := frame.GetTotalFrames()

		for curFrames := uint64(1); ; {
			if totalFrames > 0 && curFrames >= totalFrames {
				_runtime.opts.Inheritor.Iface.Stop()
				return
			}

			select {
			case <-updateTicker.C:
				internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
					timeoutTimer := time.NewTimer(_runtime.opts.ProcessQueueTimeout)
					defer timeoutTimer.Stop()

					select {
					case _runtime.processQueue <- _runtime.frameUpdate:
						curFrames++
						return
					case <-timeoutTimer.C:
						panic("process queue push frame update timeout")
					}
				})

			case <-_runtime.ctx.Done():
				return
			}
		}
	}()

	frame.SetCurFrames(0)
	_runtime.firstFrameUpdate()

	gcTicker := time.NewTicker(_runtime.opts.GCInterval)
	defer gcTicker.Stop()

	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), process)

		case <-gcTicker.C:
			_runtime.gc()

		case <-_runtime.ctx.Done():
			return
		}
	}
}

func (_runtime *RuntimeBehavior) loopWithFrameEnd() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	close(_runtime.processQueue)

	func() {
		for {
			select {
			case process, ok := <-_runtime.processQueue:
				if !ok {
					return
				}
				internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), process)

			default:
				return
			}
		}
	}()

	frame.FrameEnd()
	frame.SetCurFrames(frame.GetCurFrames() + 1)
}

func (_runtime *RuntimeBehavior) frameUpdate() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	if callback := runtime.UnsafeContext(_runtime.ctx).GetOptions().FrameEndCallback; callback != nil {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			callback(_runtime.ctx)
		})
	}

	frame.FrameEnd()

	frame.SetCurFrames(frame.GetCurFrames() + 1)

	_runtime.firstFrameUpdate()
}

func (_runtime *RuntimeBehavior) firstFrameUpdate() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	frame.FrameBegin()

	if callback := runtime.UnsafeContext(_runtime.ctx).GetOptions().FrameBeginCallback; callback != nil {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			callback(_runtime.ctx)
		})
	}

	frame.UpdateBegin()
	defer frame.UpdateEnd()

	emitEventUpdate(&_runtime.eventUpdate)
	emitEventLateUpdate(&_runtime.eventLateUpdate)
}

func (_runtime *RuntimeBehavior) loopWithBlinkFrame() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)
	totalFrames := frame.GetTotalFrames()

	gcFrames := uint64(_runtime.opts.GCInterval.Seconds() * float64(frame.GetTargetFPS()))

	for frame.SetCurFrames(0); ; {
		curFrames := frame.GetCurFrames()

		if totalFrames > 0 && curFrames >= totalFrames {
			return
		}

		if !_runtime.blinkFrameUpdate() {
			return
		}

		if curFrames%gcFrames == 0 {
			_runtime.gc()
		}

		frame.SetCurFrames(curFrames + 1)
	}
}

func (_runtime *RuntimeBehavior) loopWithBlinkFrameEnd() {
	close(_runtime.processQueue)

	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), process)

		default:
			break
		}
	}
}

func (_runtime *RuntimeBehavior) blinkFrameUpdate() bool {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	frame.FrameBegin()

	if callback := runtime.UnsafeContext(_runtime.ctx).GetOptions().FrameBeginCallback; callback != nil {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			callback(_runtime.ctx)
		})
	}

	defer func() {
		if callback := runtime.UnsafeContext(_runtime.ctx).GetOptions().FrameEndCallback; callback != nil {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				callback(_runtime.ctx)
			})
		}

		frame.FrameEnd()
	}()

	for {
		select {
		case <-_runtime.ctx.Done():
			return false

		default:
			func() {
				frame.UpdateBegin()
				defer frame.UpdateEnd()

				emitEventUpdate(&_runtime.eventUpdate)
				emitEventLateUpdate(&_runtime.eventLateUpdate)
			}()
			return true
		}
	}
}
