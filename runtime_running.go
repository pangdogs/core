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

// Run 运行
func (rt *RuntimeBehavior) Run() <-chan struct{} {
	ctx := rt.ctx

	select {
	case <-ctx.Done():
		panic(fmt.Errorf("%w: %w", ErrRuntime, context.Canceled))
	case <-ctx.TerminatedChan():
		panic(fmt.Errorf("%w: terminated", ErrRuntime))
	default:
	}

	if parentCtx, ok := ctx.GetParentContext().(concurrent.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go rt.running()

	return concurrent.UnsafeContext(ctx).GetTerminatedChan()
}

// Terminate 停止
func (rt *RuntimeBehavior) Terminate() <-chan struct{} {
	return rt.ctx.Terminate()
}

// TerminatedChan 已停止chan
func (rt *RuntimeBehavior) TerminatedChan() <-chan struct{} {
	return rt.ctx.TerminatedChan()
}

func (rt *RuntimeBehavior) running() {
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

	close(concurrent.UnsafeContext(rt.ctx).GetTerminatedChan())
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
	pluginBundle := rt.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(rt.activatePlugin)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(rt.deactivatePlugin)

	pluginBundle.Range(func(pluginInfo plugin.PluginInfo) bool {
		rt.activatePlugin(pluginInfo)
		return true
	})
}

func (rt *RuntimeBehavior) shutPlugin() {
	pluginBundle := rt.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(nil)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(nil)

	pluginBundle.ReversedRange(func(pluginInfo plugin.PluginInfo) bool {
		rt.deactivatePlugin(pluginInfo)
		return true
	})
}

func (rt *RuntimeBehavior) activatePlugin(pluginInfo plugin.PluginInfo) {
	if pluginInit, ok := pluginInfo.Face.Iface.(LifecycleRuntimePluginInit); ok {
		generic.MakeAction1(pluginInit.InitRP).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
	}
	plugin.UnsafePluginBundle(rt.ctx.GetPluginBundle()).SetActive(pluginInfo.Name, true)
}

func (rt *RuntimeBehavior) deactivatePlugin(pluginInfo plugin.PluginInfo) {
	plugin.UnsafePluginBundle(rt.ctx.GetPluginBundle()).SetActive(pluginInfo.Name, false)
	if pluginShut, ok := pluginInfo.Face.Iface.(LifecycleRuntimePluginShut); ok {
		generic.MakeAction1(pluginShut.ShutRP).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
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

	event.Clean(hooks[:])

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningEnd()
	}
}

func (rt *RuntimeBehavior) mainLoop() {
	frame := rt.opts.Frame

	if frame == nil {
		rt.loopingNoFrame()
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
