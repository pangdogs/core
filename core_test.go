/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package core_test

import (
	"context"
	"log"
	"testing"
	"time"

	"git.golaxy.org/core/utils/assertion"

	"git.golaxy.org/core"
	"git.golaxy.org/core/define"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

func Test_StartService(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type EntityTest1 struct {
	ec.EntityBehavior
}

func (e *EntityTest1) Awake() {
	log.Printf("EntityTest1 %s Awake", e.GetId())
}

func (e *EntityTest1) Start() {
	log.Printf("EntityTest1 %s Start", e.GetId())
}

func (e *EntityTest1) Shut() {
	log.Printf("EntityTest1 %s Shut", e.GetId())
}

func (e *EntityTest1) Dispose() {
	log.Printf("EntityTest1 %s Dispose", e.GetId())
}

type EntityTest2 struct {
	ec.EntityBehavior
}

func (e *EntityTest2) Awake() {
	log.Printf("EntityTest2 %s Awake", e.GetId())
}

func (e *EntityTest2) Start() {
	log.Printf("EntityTest2 %s Start", e.GetId())
}

func (e *EntityTest2) Shut() {
	log.Printf("EntityTest2 %s Shut", e.GetId())
}

func (e *EntityTest2) Dispose() {
	log.Printf("EntityTest2 %s Dispose", e.GetId())
}

type ComponentTest1 struct {
	ec.ComponentBehavior
}

func (c *ComponentTest1) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest1) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest1) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest1) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

type ComponentTest2 struct {
	ec.ComponentBehavior
}

func (c *ComponentTest2) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest2) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest2) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest2) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

type ComponentTest3 struct {
	ec.ComponentBehavior
}

func (c *ComponentTest3) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest3) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest3) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTest3) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

func Test_ServiceRegisterEntityPT(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				ctx.GetEntityLib().Declare(
					pt.NewEntityDescriptor("Test1").SetInstance(EntityTest1{}),
					ComponentTest1{},
				)
				ctx.GetEntityLib().Declare(
					pt.EntityDescriptor{
						Prototype: "Test2",
						Instance:  EntityTest2{},
					},
					ComponentTest1{},
					ComponentTest2{},
				)
				ctx.GetEntityLib().Declare(
					"Test3",
					ComponentTest1{},
					ComponentTest2{},
					ComponentTest3{},
				)
			case service.RunningEvent_ComponentPTDeclared:
				compPT := args[0].(ec.ComponentPT)
				log.Println("component pt declared:", compPT.Prototype(), compPT.InstanceRT())
				return
			case service.RunningEvent_EntityPTDeclared:
				entityPT := args[0].(ec.EntityPT)
				log.Println("+ entity pt declared:", entityPT.Prototype(), entityPT.InstanceRT())
				for _, comp := range entityPT.ListComponents() {
					log.Println("- builtin component:", comp.Name, comp.PT.Prototype())
				}
				return
			}
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ServiceAddIn1 struct{}

func (ServiceAddIn1) Init(ctx service.Context) {
	log.Println("ServiceAddIn1 Init")
}

func (ServiceAddIn1) Shut(ctx service.Context) {
	log.Println("ServiceAddIn1 Shut")
}

func NewServiceAddIn1(...any) *ServiceAddIn1 {
	return &ServiceAddIn1{}
}

var (
	serviceAddIn1Define    = define.AddIn(NewServiceAddIn1)
	serviceAddIn1Name      = serviceAddIn1Define.Name
	serviceAddIn1Install   = serviceAddIn1Define.Install
	serviceAddIn1Uninstall = serviceAddIn1Define.Uninstall
	serviceAddIn1Using     = serviceAddIn1Define.Using
)

type IServiceAddIn2 interface{}

type ServiceAddIn2 struct{}

func (ServiceAddIn2) Init(ctx service.Context) {
	log.Println("ServiceAddIn2 Init")
}

func (ServiceAddIn2) Shut(ctx service.Context) {
	log.Println("ServiceAddIn2 Shut")
}

func NewServiceAddIn2(...any) IServiceAddIn2 {
	return &ServiceAddIn2{}
}

var (
	serviceAddIn2Define    = define.AddIn(NewServiceAddIn2)
	serviceAddIn2Name      = serviceAddIn2Define.Name
	serviceAddIn2Install   = serviceAddIn2Define.Install
	serviceAddIn2Uninstall = serviceAddIn2Define.Uninstall
	serviceAddIn2Using     = serviceAddIn2Define.Using
)

func Test_ServiceAddIn(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				serviceAddIn1Install(ctx)
				serviceAddIn2Install(ctx)
			}
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

func Test_CreateEntity(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					SetInstance(EntityTest1{}).
					AddComponent(ComponentTest1{}).
					Declare()
				core.BuildEntityPT(ctx, "Test2").
					SetInstance(EntityTest2{}).
					AddComponent(ComponentTest1{}).
					AddComponent(ComponentTest2{}).
					Declare()
				core.BuildEntityPT(ctx, "Test3").
					AddComponent(ComponentTest1{}).
					AddComponent(ComponentTest2{}).
					AddComponent(ComponentTest3{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.Context.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								core.BuildEntity(ctx, "Test1").New()
								core.BuildEntity(ctx, "Test2").New()
								core.BuildEntity(ctx, "Test3").New()
							case runtime.RunningEvent_EntityActivating:
								entity := args[0].(ec.Entity)
								log.Println("entity activating:", entity.GetId(), entity.GetPT().Prototype())
								return
							case runtime.RunningEvent_EntityActivatingDone:
								entity := args[0].(ec.Entity)
								log.Println("entity activated:", entity.GetId(), entity.GetPT().Prototype())
								return
							case runtime.RunningEvent_EntityDeactivating:
								entity := args[0].(ec.Entity)
								log.Println("entity deactivating:", entity.GetId(), entity.GetPT().Prototype())
								return
							case runtime.RunningEvent_EntityDeactivatingDone:
								entity := args[0].(ec.Entity)
								log.Println("entity deactivated:", entity.GetId(), entity.GetPT().Prototype())
								return
							}
							log.Println("runtime event:", runningEvent)
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			case service.RunningEvent_EntityRegistered:
				entity := args[0].(ec.ConcurrentEntity)
				log.Println("entity registered:", entity.GetId(), entity.GetPT().Prototype())
				return
			case service.RunningEvent_EntityUnregistered:
				entity := args[0].(ec.ConcurrentEntity)
				log.Println("entity unregistered:", entity.GetId(), entity.GetPT().Prototype())
				return
			}
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestEnable1 struct {
	ec.ComponentBehavior
}

func (c *ComponentTestEnable1) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
	c.SetEnable(false)
}

func (c *ComponentTestEnable1) OnEnable() {
	log.Printf("Component %s.%s Enable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable1) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable1) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable1) OnDisable() {
	log.Printf("Component %s.%s Disable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable1) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

type ComponentTestEnable2 struct {
	ec.ComponentBehavior
}

func (c *ComponentTestEnable2) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable2) OnEnable() {
	log.Printf("Component %s.%s Enable", c.GetEntity().GetId(), c.GetName())
	c.SetEnable(false)
}

func (c *ComponentTestEnable2) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable2) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable2) OnDisable() {
	log.Printf("Component %s.%s Disable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable2) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

type ComponentTestEnable3 struct {
	ec.ComponentBehavior
}

func (c *ComponentTestEnable3) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable3) OnEnable() {
	log.Printf("Component %s.%s Enable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable3) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
	c.SetEnable(false)
}

func (c *ComponentTestEnable3) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable3) OnDisable() {
	log.Printf("Component %s.%s Disable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable3) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

type ComponentTestEnable4 struct {
	ec.ComponentBehavior
}

func (c *ComponentTestEnable4) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable4) OnEnable() {
	log.Printf("Component %s.%s Enable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable4) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable4) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
	c.SetEnable(false)
}

func (c *ComponentTestEnable4) OnDisable() {
	log.Printf("Component %s.%s Disable", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestEnable4) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

func Test_EntityComponentEnable(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestEnable1{}).
					AddComponent(ComponentTestEnable2{}).
					AddComponent(ComponentTestEnable3{}).
					AddComponent(ComponentTestEnable4{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.Context.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								core.BuildEntity(ctx, "Test1").New()
							}
							log.Println("runtime event:", runningEvent)
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestDynamic struct {
	ec.ComponentBehavior

	test2 *ComponentTest2
	test3 *ComponentTest3
}

func (c *ComponentTestDynamic) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestDynamic) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())

	if err := assertion.Inject(c.GetEntity(), c); err != nil {
		log.Panicln("Inject error:", err)
	}

	log.Println("Inject:", c.test2, c.test3)
}

func (c *ComponentTestDynamic) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestDynamic) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

func Test_EntityDynamicComponent(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				ctx.GetEntityLib().GetComponentLib().Declare(ComponentTest3{})

				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestDynamic{}).
					AddComponent(ComponentTest2{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.Context.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								core.BuildEntity(ctx, "Test1").New()
							}
							log.Println("runtime event:", runningEvent)
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type RuntimeAddIn1 struct{}

func (RuntimeAddIn1) Init(ctx runtime.Context) {
	log.Println("RuntimeAddIn1 Init")
}

func (RuntimeAddIn1) Shut(ctx runtime.Context) {
	log.Println("RuntimeAddIn1 Shut")
}

func (RuntimeAddIn1) OnContextRunningEvent(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
	log.Println("RuntimeAddIn1 OnContextRunningEvent:", runningEvent)
}

func NewRuntimeAddIn1(...any) *RuntimeAddIn1 {
	return &RuntimeAddIn1{}
}

var (
	runtimeAddIn1Define    = define.AddIn(NewRuntimeAddIn1)
	runtimeAddIn1Name      = runtimeAddIn1Define.Name
	runtimeAddIn1Install   = runtimeAddIn1Define.Install
	runtimeAddIn1Uninstall = runtimeAddIn1Define.Uninstall
	runtimeAddIn1Using     = runtimeAddIn1Define.Using
)

func Test_RuntimeAddIn(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.Context.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Birth:
								runtimeAddIn1Install(ctx)
							}
							log.Println("runtime event:", runningEvent)
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
			log.Println("service event:", runningEvent)
		}),
	)

	<-core.NewService(svcCtx).Run()
}
