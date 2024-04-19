package core

import (
	"context"
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/util/generic"
)

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (rt *RuntimeBehavior) Run() <-chan struct{} {
	ctx := rt.ctx

	select {
	case <-ctx.Done():
		panic(fmt.Errorf("%w: %w", ErrRuntime, context.Canceled))
	default:
	}

	if !rt.started.CompareAndSwap(false, true) {
		panic(fmt.Errorf("%w: running already started", ErrRuntime))
	}

	if parentCtx, ok := ctx.GetParentContext().(concurrent.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	shutChan := make(chan struct{}, 1)
	go rt.running(shutChan)

	return shutChan
}

// Terminate 停止
func (rt *RuntimeBehavior) Terminate() {
	rt.ctx.GetCancelFunc()()
}

func (rt *RuntimeBehavior) running(shutChan chan struct{}) {
	ctx := rt.ctx

	rt.changeRunningState(runtime.RunningState_Starting)

	hooks := rt.loopStart()

	rt.changeRunningState(runtime.RunningState_Started)

	rt.mainLoop()

	rt.changeRunningState(runtime.RunningState_Terminating)

	rt.loopStop(hooks)
	ctx.GetWaitGroup().Wait()

	rt.changeRunningState(runtime.RunningState_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(concurrent.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	shutChan <- struct{}{}
}

func (rt *RuntimeBehavior) changeRunningState(state runtime.RunningState) {
	switch state {
	case runtime.RunningState_Starting:
		rt.initPlugin()
	case runtime.RunningState_FrameLoopBegin:
		runtime.UnsafeFrame(rt.opts.Frame).LoopBegin()
	case runtime.RunningState_FrameUpdateBegin:
		runtime.UnsafeFrame(rt.opts.Frame).UpdateBegin()
	case runtime.RunningState_FrameUpdateEnd:
		runtime.UnsafeFrame(rt.opts.Frame).UpdateEnd()
	case runtime.RunningState_FrameLoopEnd:
		runtime.UnsafeFrame(rt.opts.Frame).LoopEnd()
	case runtime.RunningState_Terminated:
		rt.shutPlugin()
	}

	runtime.UnsafeContext(rt.ctx).ChangeRunningState(state)
}

func (rt *RuntimeBehavior) initPlugin() {
	if pluginBundle := runtime.UnsafeContext(rt.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.Range(func(info plugin.PluginInfo) bool {
			if pluginInit, ok := info.Face.Iface.(LifecycleRuntimePluginInit); ok {
				generic.MakeAction1(pluginInit.InitRP).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
			}
			plugin.UnsafePluginBundle(pluginBundle).Activate(info.Name, true)
			return true
		})
	}
}

func (rt *RuntimeBehavior) shutPlugin() {
	if pluginBundle := runtime.UnsafeContext(rt.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.ReverseRange(func(info plugin.PluginInfo) bool {
			plugin.UnsafePluginBundle(pluginBundle).Activate(info.Name, false)
			if pluginShut, ok := info.Face.Iface.(LifecycleRuntimePluginShut); ok {
				generic.MakeAction1(pluginShut.ShutRP).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
			}
			return true
		})
	}
}

func (rt *RuntimeBehavior) loopStart() (hooks [5]event.Hook) {
	ctx := rt.ctx
	frame := rt.opts.Frame

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningBegin()
	}

	hooks[0] = runtime.BindEventEntityMgrAddEntity(ctx.GetEntityMgr(), rt)
	hooks[1] = runtime.BindEventEntityMgrRemoveEntity(ctx.GetEntityMgr(), rt)
	hooks[2] = runtime.BindEventEntityMgrEntityAddComponents(ctx.GetEntityMgr(), rt)
	hooks[3] = runtime.BindEventEntityMgrEntityRemoveComponent(ctx.GetEntityMgr(), rt)
	hooks[4] = runtime.BindEventEntityMgrEntityFirstAccessComponent(ctx.GetEntityMgr(), rt)

	return
}

func (rt *RuntimeBehavior) loopStop(hooks [5]event.Hook) {
	frame := rt.opts.Frame

	for i := range hooks {
		hooks[i].Unbind()
	}

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningEnd()
	}
}

func (rt *RuntimeBehavior) mainLoop() {
	frame := rt.opts.Frame

	if frame == nil {
		rt.loopingNoFrame()
	} else if frame.GetBlink() {
		rt.loopingBlinkFrame()
	} else {
		rt.loopingWithFrame()
	}
}

func (rt *RuntimeBehavior) runTask(task _Task) {
	switch task.typ {
	case _TaskType_Call:
		rt.changeRunningState(runtime.RunningState_RunCallBegin)
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		rt.changeRunningState(runtime.RunningState_RunCallEnd)
	case _TaskType_Frame:
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}
}

func (rt *RuntimeBehavior) runGC() {
	rt.changeRunningState(runtime.RunningState_RunGCBegin)
	rt.gc()
	rt.changeRunningState(runtime.RunningState_RunGCEnd)
}
