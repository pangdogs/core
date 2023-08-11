package golaxy

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util"
	"time"
)

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (_runtime *RuntimeBehavior) Run() <-chan struct{} {
	ctx := _runtime.ctx

	if !runtime.UnsafeContext(ctx).MarkRunning(true) {
		panic("runtime already running")
	}

	shutChan := make(chan struct{}, 1)

	runtime.UnsafeContext(ctx).SetFrame(_runtime.opts.Frame)

	_runtime.processQueue = make(chan func(), _runtime.opts.ProcessQueueCapacity)
	runtime.UnsafeContext(ctx).SetCallee(_runtime)

	if parentCtx, ok := ctx.GetParentContext().(internal.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go _runtime.running(shutChan)

	return shutChan
}

// Terminate 停止
func (_runtime *RuntimeBehavior) Terminate() {
	_runtime.ctx.GetCancelFunc()()
}

func (_runtime *RuntimeBehavior) running(shutChan chan struct{}) {
	ctx := _runtime.ctx
	frame := _runtime.opts.Frame

	_runtime.changeRunningState(runtime.RunningState_Starting)

	_runtime.initPlugin()
	hooks := _runtime.loopStart()

	_runtime.changeRunningState(runtime.RunningState_Started)

	if frame == nil {
		_runtime.loopingNoFrame()
		_runtime.loopingNoFrameEnd()
	} else if frame.Blink() {
		_runtime.loopingBlinkFrame()
		_runtime.loopingBlinkFrameEnd()
	} else {
		_runtime.loopingFrame()
		_runtime.loopingFrameEnd()
	}

	_runtime.changeRunningState(runtime.RunningState_Terminating)

	_runtime.loopStop(hooks)
	ctx.GetWaitGroup().Wait()
	_runtime.shutPlugin()

	_runtime.changeRunningState(runtime.RunningState_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(internal.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	runtime.UnsafeContext(ctx).MarkRunning(false)
	shutChan <- struct{}{}
}

func (_runtime *RuntimeBehavior) loopStart() (hooks [5]localevent.Hook) {
	ctx := _runtime.ctx
	frame := _runtime.opts.Frame

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningBegin()
	}

	hooks[0] = localevent.BindEvent[runtime.EventEntityMgrAddEntity](ctx.GetEntityMgr().EventEntityMgrAddEntity(), _runtime)
	hooks[1] = localevent.BindEvent[runtime.EventEntityMgrRemoveEntity](ctx.GetEntityMgr().EventEntityMgrRemoveEntity(), _runtime)
	hooks[2] = localevent.BindEvent[runtime.EventEntityMgrEntityAddComponents](ctx.GetEntityMgr().EventEntityMgrEntityAddComponents(), _runtime)
	hooks[3] = localevent.BindEvent[runtime.EventEntityMgrEntityRemoveComponent](ctx.GetEntityMgr().EventEntityMgrEntityRemoveComponent(), _runtime)
	hooks[4] = localevent.BindEvent[runtime.EventEntityMgrEntityFirstAccessComponent](ctx.GetEntityMgr().EventEntityMgrEntityFirstAccessComponent(), _runtime)

	ctx.GetEntityMgr().RangeEntities(func(entity ec.Entity) bool {
		internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), func() {
			_runtime.OnEntityMgrAddEntity(ctx.GetEntityMgr(), entity)
		})
		return true
	})

	return
}

func (_runtime *RuntimeBehavior) loopStop(hooks [5]localevent.Hook) {
	ctx := _runtime.ctx
	frame := _runtime.opts.Frame

	ctx.GetEntityMgr().ReverseRangeEntities(func(entity ec.Entity) bool {
		internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), func() {
			_runtime.OnEntityMgrRemoveEntity(ctx.GetEntityMgr(), entity)
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

func (_runtime *RuntimeBehavior) loopingNoFrame() {
	ctx := _runtime.ctx

	defer close(_runtime.processQueue)

	gcTicker := time.NewTicker(_runtime.opts.GCInterval)
	defer gcTicker.Stop()

	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), process)

		case <-gcTicker.C:
			_runtime.gc()

		case <-ctx.Done():
			return
		}
	}
}

func (_runtime *RuntimeBehavior) loopingNoFrameEnd() {
	ctx := _runtime.ctx

	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), process)

		default:
			return
		}
	}
}

func (_runtime *RuntimeBehavior) loopingFrame() {
	ctx := _runtime.ctx
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	go func(curFrames, totalFrames uint64, targetFPS float32) {
		updateTicker := time.NewTicker(time.Duration(float64(time.Second) / float64(targetFPS)))
		defer updateTicker.Stop()

		for {
			if totalFrames > 0 && curFrames >= totalFrames {
				ctx.GetCancelFunc()()
				return
			}

			select {
			case <-updateTicker.C:
				func() {
					defer func() {
						recover()
					}()
					select {
					case _runtime.processQueue <- _runtime.frameLoop:
					default:
					}
				}()
			case <-ctx.Done():
				return
			default:
			}
		}
	}(frame.GetCurFrames()+1, frame.GetTotalFrames(), frame.GetTargetFPS())

	defer close(_runtime.processQueue)

	gcTicker := time.NewTicker(_runtime.opts.GCInterval)
	defer gcTicker.Stop()

	for _runtime.frameLoopBegin(); ; {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				return
			}
			internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), process)

		case <-gcTicker.C:
			_runtime.gc()

		case <-ctx.Done():
			return
		}
	}
}

func (_runtime *RuntimeBehavior) loopingFrameEnd() {
	ctx := _runtime.ctx

loop:
	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				break loop
			}
			internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), process)

		default:
			break loop
		}
	}

	_runtime.frameLoopEnd()
}

func (_runtime *RuntimeBehavior) frameLoop() {
	_runtime.frameLoopEnd()
	_runtime.frameLoopBegin()
}

func (_runtime *RuntimeBehavior) frameLoopBegin() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	frame.LoopBegin()
	_runtime.changeRunningState(runtime.RunningState_FrameLoopBegin)

	frame.UpdateBegin()
	_runtime.changeRunningState(runtime.RunningState_FrameUpdateBegin)

	emitEventUpdate(&_runtime.eventUpdate)
	emitEventLateUpdate(&_runtime.eventLateUpdate)

	frame.UpdateEnd()
	_runtime.changeRunningState(runtime.RunningState_FrameUpdateEnd)

	_runtime.changeRunningState(runtime.RunningState_AsyncProcessingBegin)
}

func (_runtime *RuntimeBehavior) frameLoopEnd() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	_runtime.changeRunningState(runtime.RunningState_AsyncProcessingEnd)

	frame.LoopEnd()
	_runtime.changeRunningState(runtime.RunningState_FrameLoopEnd)

	frame.SetCurFrames(frame.GetCurFrames() + 1)
}

func (_runtime *RuntimeBehavior) loopingBlinkFrame() {
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	totalFrames := frame.GetTotalFrames()
	gcFrames := uint64(_runtime.opts.GCInterval.Seconds() * float64(frame.GetTargetFPS()))

	for {
		curFrames := frame.GetCurFrames()

		if totalFrames > 0 && curFrames >= totalFrames {
			return
		}

		if !_runtime.blinkFrameLoop() {
			return
		}

		if curFrames%gcFrames == 0 {
			_runtime.gc()
		}

		frame.SetCurFrames(curFrames + 1)
	}
}

func (_runtime *RuntimeBehavior) loopingBlinkFrameEnd() {
	ctx := _runtime.ctx

	close(_runtime.processQueue)

	_runtime.changeRunningState(runtime.RunningState_AsyncProcessingBegin)

loop:
	for {
		select {
		case process, ok := <-_runtime.processQueue:
			if !ok {
				break loop
			}
			internal.CallOuterVoid(ctx.GetAutoRecover(), ctx.GetReportError(), process)

		default:
			break loop
		}
	}

	_runtime.changeRunningState(runtime.RunningState_AsyncProcessingEnd)
}

func (_runtime *RuntimeBehavior) blinkFrameLoop() bool {
	ctx := _runtime.ctx
	frame := runtime.UnsafeFrame(_runtime.opts.Frame)

	select {
	case <-ctx.Done():
		return false
	default:
	}

	frame.LoopBegin()
	_runtime.changeRunningState(runtime.RunningState_FrameLoopBegin)

	frame.UpdateBegin()
	_runtime.changeRunningState(runtime.RunningState_FrameUpdateBegin)

	emitEventUpdate(&_runtime.eventUpdate)
	emitEventLateUpdate(&_runtime.eventLateUpdate)

	frame.UpdateEnd()
	_runtime.changeRunningState(runtime.RunningState_FrameUpdateEnd)

	frame.LoopEnd()
	_runtime.changeRunningState(runtime.RunningState_FrameLoopEnd)

	return true
}

func (_runtime *RuntimeBehavior) changeRunningState(state runtime.RunningState) {
	if handler := runtime.UnsafeContext(_runtime.ctx).GetOptions().RunningHandler; handler != nil {
		internal.CallOuterVoid(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			handler(_runtime.ctx, state)
		})
	}
}

func (_runtime *RuntimeBehavior) initPlugin() {
	if pluginBundle := runtime.UnsafeContext(_runtime.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.Range(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginInit, ok := pluginFace.Iface.(LifecycleRuntimePluginInit); ok {
				internal.CallOuterVoid(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
					pluginInit.InitRP(_runtime.ctx)
				})
			}
			return true
		})
	}
}

func (_runtime *RuntimeBehavior) shutPlugin() {
	if pluginBundle := runtime.UnsafeContext(_runtime.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.ReverseRange(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginShut, ok := pluginFace.Iface.(LifecycleRuntimePluginShut); ok {
				internal.CallOuterVoid(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
					pluginShut.ShutRP(_runtime.ctx)
				})
			}
			return true
		})
	}
}
