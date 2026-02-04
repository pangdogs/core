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
	"strings"
	"testing"
	"time"

	"git.golaxy.org/core/utils/assertion"
	"git.golaxy.org/core/utils/uid"
	"github.com/elliotchance/pie/v2"

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
			log.Println("service event:", runningEvent, args)
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
			}
			log.Println("service event:", runningEvent, args)
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
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								core.BuildEntity(ctx, "Test1").New()
								core.BuildEntity(ctx, "Test2").New()
								core.BuildEntity(ctx, "Test3").New()
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enable(false)),
				)
			}
			log.Println("service event:", runningEvent, args)
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
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								core.BuildEntity(ctx, "Test1").New()
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enable(false)),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestDynamic1 struct {
	ec.ComponentBehavior

	test2 *ComponentTest2
	test3 *ComponentTest3
}

func (c *ComponentTestDynamic1) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestDynamic1) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())

	if err := assertion.Inject(c.GetEntity(), c); err != nil {
		log.Panicln("Inject error:", err)
	}

	log.Println("Inject:", c.test2, c.test3)
}

func (c *ComponentTestDynamic1) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestDynamic1) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

type ComponentTestDynamic2 struct {
	ec.ComponentBehavior

	test2 *ComponentTest2
	test3 *ComponentTest3
}

func (c *ComponentTestDynamic2) Awake() {
	log.Printf("Component %s.%s Awake", c.GetEntity().GetId(), c.GetName())

	if err := assertion.Inject(c.GetEntity(), c); err != nil {
		log.Panicln("Inject error:", err)
	}

	log.Println("Inject:", c.test2, c.test3)
}

func (c *ComponentTestDynamic2) Start() {
	log.Printf("Component %s.%s Start", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestDynamic2) Shut() {
	log.Printf("Component %s.%s Shut", c.GetEntity().GetId(), c.GetName())
}

func (c *ComponentTestDynamic2) Dispose() {
	log.Printf("Component %s.%s Dispose", c.GetEntity().GetId(), c.GetName())
}

func Test_EntityDynamicComponent(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				ctx.GetEntityLib().GetComponentLib().Declare(ComponentTest2{})
				ctx.GetEntityLib().GetComponentLib().Declare(ComponentTest3{})

				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestDynamic1{}).
					AddComponent(ComponentTest2{}).
					Declare()

				core.BuildEntityPT(ctx, "Test2").
					AddComponent(ComponentTestDynamic2{}).
					Declare()

			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								core.BuildEntity(ctx, "Test1").New()
								core.BuildEntity(ctx, "Test2").New()
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enable(false)),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestParent struct {
	ec.ComponentBehavior
}

func (c *ComponentTestParent) Awake() {
	ec.BindEventTreeNodeAddChild(c.GetEntity(), c)
	ec.BindEventTreeNodeRemoveChild(c.GetEntity(), c)
}

func (c *ComponentTestParent) OnTreeNodeAddChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeAddChild %s <- %s", entity.GetId(), childId)
}

func (c *ComponentTestParent) OnTreeNodeRemoveChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeRemoveChild %s x- %s", entity.GetId(), childId)
}

type ComponentTestChild struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChild) Awake() {
	ec.BindEventTreeNodeAttachParent(c.GetEntity(), c)
	ec.BindEventTreeNodeDetachParent(c.GetEntity(), c)
}

func (c *ComponentTestChild) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.GetId(), parentId)
}

func (c *ComponentTestChild) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.GetId(), parentId)
}

func PrintEntityTreeForest(entityTree runtime.EntityTree) {
	entityTree.EachChildren(runtime.ForestNodeId, func(entity ec.Entity) {
		PrintEntityTree(entity)
	})
}

func PrintEntityTree(entity ec.Entity, depth ...int) {
	entityTree := runtime.Current(entity).GetEntityTree()
	if b, _ := entityTree.IsFreedom(entity.GetId()); b {
		return
	}

	root := ""

	isRoot, _ := entityTree.IsRoot(entity.GetId())
	if isRoot {
		root = "R"
	}

	leaf := ""

	isLeaf, _ := entityTree.IsLeaf(entity.GetId())
	if isLeaf {
		leaf = "L"
	}

	_depth := pie.First(depth)

	if isLeaf {
		log.Printf("%s- [%s] %s%s", strings.Repeat(" ", _depth), entity.GetId(), root, leaf)
	} else {
		log.Printf("%s+ [%s] %s%s", strings.Repeat(" ", _depth), entity.GetId(), root, leaf)
	}

	entityTree.EachChildren(entity.GetId(), func(entity ec.Entity) {
		PrintEntityTree(entity, _depth+1)
	})
}

func Test_EntityTree(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestParent{}).
					AddComponent(ComponentTestChild{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Starting:
								runtime.BindEventEntityTreeAddNode(ctx.GetEntityTree(), runtime.HandleEventEntityTreeAddNode(func(entityTree runtime.EntityTree, parentId, childId uid.Id) {
									var children []uid.Id

									entityTree.EachChildren(parentId, func(entity ec.Entity) {
										children = append(children, entity.GetId())
									})

									log.Printf("OnEntityTreeAddNode %s: %v + %s", parentId, children, childId)
								}))
								runtime.BindEventEntityTreeRemoveNode(ctx.GetEntityTree(), runtime.HandleEventEntityTreeRemoveNode(func(entityTree runtime.EntityTree, parentId, childId uid.Id) {
									var children []uid.Id

									entityTree.EachChildren(parentId, func(entity ec.Entity) {
										children = append(children, entity.GetId())
									})

									log.Printf("OnEntityTreeRemoveNode %s: %v - %s", parentId, children, childId)
								}))
								runtime.BindEventEntityTreeMoveNode(ctx.GetEntityTree(), runtime.HandleEventEntityTreeMoveNode(func(entityTree runtime.EntityTree, childId, fromParentId, toParentId uid.Id) {
									log.Printf("OnEntityTreeMoveNode %s: %s => %s", childId, fromParentId, toParentId)
								}))
							case runtime.RunningEvent_Started:
								root, err := core.BuildEntity(ctx, "Test1").New()
								if err != nil {
									log.Panicln("new root error:", err)
								}

								err = ctx.GetEntityTree().MakeRoot(root.GetId())
								if err != nil {
									log.Panicln("make root error:", err)
								}

								child1, err := core.BuildEntity(ctx, "Test1").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child1 error:", err)
								}

								child2, err := core.BuildEntity(ctx, "Test1").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child2 error:", err)
								}

								child3, err := core.BuildEntity(ctx, "Test1").SetParentId(child1.GetId()).New()
								if err != nil {
									log.Panicln("new child3 error:", err)
								}

								child4, err := core.BuildEntity(ctx, "Test1").SetParentId(child3.GetId()).New()
								if err != nil {
									log.Panicln("new child4 error:", err)
								}

								child5, err := core.BuildEntity(ctx, "Test1").SetParentId(child3.GetId()).New()
								if err != nil {
									log.Panicln("new child5 error:", err)
								}

								child6, err := core.BuildEntity(ctx, "Test1").SetParentId(child3.GetId()).New()
								if err != nil {
									log.Panicln("new child6 error:", err)
								}

								child7, err := core.BuildEntity(ctx, "Test1").SetParentId(runtime.ForestNodeId).New()
								if err != nil {
									log.Panicln("new child7 error:", err)
								}

								child8, err := core.BuildEntity(ctx, "Test1").SetParentId(child2.GetId()).New()
								if err != nil {
									log.Panicln("new child8 error:", err)
								}

								log.Println("1. testing detach node")

								PrintEntityTreeForest(ctx.GetEntityTree())

								err = ctx.GetEntityTree().DetachNode(child2.GetId())
								if err != nil {
									log.Panicln("detach child2 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Println("2. testing remove node")

								PrintEntityTreeForest(ctx.GetEntityTree())

								err = ctx.GetEntityTree().RemoveNode(child3.GetId())
								if err != nil {
									log.Panicln("remove child3 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Println("3. testing move node")

								PrintEntityTreeForest(ctx.GetEntityTree())

								err = ctx.GetEntityTree().MoveNode(child7.GetId(), child2.GetId())
								if err != nil {
									log.Panicln("move child7 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								_ = child1
								_ = child2
								_ = child3
								_ = child4
								_ = child5
								_ = child6
								_ = child7
								_ = child8
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enable(false)),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestChildDetachInAttaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChildDetachInAttaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.GetEntity(), c)
	ec.BindEventTreeNodeDetachParent(c.GetEntity(), c)
}

func (c *ComponentTestChildDetachInAttaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.GetId(), parentId)

	err := runtime.Current(entity).GetEntityTree().DetachNode(entity.GetId())
	if err != nil {
		log.Printf("OnTreeNodeAttachParent %s DetachNode failed, %s", entity.GetId(), err)
	}
}

func (c *ComponentTestChildDetachInAttaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.GetId(), parentId)
}

type ComponentTestChildRemoveInAttaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChildRemoveInAttaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.GetEntity(), c)
	ec.BindEventTreeNodeDetachParent(c.GetEntity(), c)
}

func (c *ComponentTestChildRemoveInAttaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.GetId(), parentId)

	err := runtime.Current(entity).GetEntityTree().RemoveNode(entity.GetId())
	if err != nil {
		log.Printf("OnTreeNodeAttachParent %s RemoveNode failed, %s", entity.GetId(), err)
	}
}

func (c *ComponentTestChildRemoveInAttaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.GetId(), parentId)
}

type ComponentTestChildDestroyInAttaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChildDestroyInAttaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.GetEntity(), c)
	ec.BindEventTreeNodeDetachParent(c.GetEntity(), c)
}

func (c *ComponentTestChildDestroyInAttaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.GetId(), parentId)
	entity.Destroy()
}

func (c *ComponentTestChildDestroyInAttaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.GetId(), parentId)
}

type ComponentTestChildDestroyInDetaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChildDestroyInDetaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.GetEntity(), c)
	ec.BindEventTreeNodeDetachParent(c.GetEntity(), c)
}

func (c *ComponentTestChildDestroyInDetaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.GetId(), parentId)
}

func (c *ComponentTestChildDestroyInDetaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.GetId(), parentId)
	entity.Destroy()
}

type ComponentTestParentDestroyInAttaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestParentDestroyInAttaching) Awake() {
	ec.BindEventTreeNodeAddChild(c.GetEntity(), c)
	ec.BindEventTreeNodeRemoveChild(c.GetEntity(), c)
}

func (c *ComponentTestParentDestroyInAttaching) OnTreeNodeAddChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeAddChild %s <- %s", entity.GetId(), childId)
	entity.Destroy()
}

func (c *ComponentTestParentDestroyInAttaching) OnTreeNodeRemoveChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeRemoveChild %s x- %s", entity.GetId(), childId)
}

type ComponentTestParentDestroyInDetaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestParentDestroyInDetaching) Awake() {
	ec.BindEventTreeNodeAddChild(c.GetEntity(), c)
	ec.BindEventTreeNodeRemoveChild(c.GetEntity(), c)
}

func (c *ComponentTestParentDestroyInDetaching) OnTreeNodeAddChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeAddChild %s <- %s", entity.GetId(), childId)
}

func (c *ComponentTestParentDestroyInDetaching) OnTreeNodeRemoveChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeRemoveChild %s x- %s", entity.GetId(), childId)
	entity.Destroy()
}

func Test_EntityTreeSequence(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestParent{}).
					AddComponent(ComponentTestChild{}).
					Declare()
				core.BuildEntityPT(ctx, "Test2").
					AddComponent(ComponentTestParent{}).
					AddComponent(ComponentTestChildDetachInAttaching{}).
					Declare()
				core.BuildEntityPT(ctx, "Test3").
					AddComponent(ComponentTestParent{}).
					AddComponent(ComponentTestChildRemoveInAttaching{}).
					Declare()
				core.BuildEntityPT(ctx, "Test4").
					AddComponent(ComponentTestParent{}).
					AddComponent(ComponentTestChildDestroyInAttaching{}).
					Declare()
				core.BuildEntityPT(ctx, "Test5").
					AddComponent(ComponentTestParent{}).
					AddComponent(ComponentTestChildDestroyInDetaching{}).
					Declare()
				core.BuildEntityPT(ctx, "Test6").
					AddComponent(ComponentTestParentDestroyInAttaching{}).
					AddComponent(ComponentTestChild{}).
					Declare()
				core.BuildEntityPT(ctx, "Test7").
					AddComponent(ComponentTestParentDestroyInDetaching{}).
					AddComponent(ComponentTestChild{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Starting:
								runtime.BindEventEntityTreeAddNode(ctx.GetEntityTree(), runtime.HandleEventEntityTreeAddNode(func(entityTree runtime.EntityTree, parentId, childId uid.Id) {
									var children []uid.Id

									entityTree.EachChildren(parentId, func(entity ec.Entity) {
										children = append(children, entity.GetId())
									})

									log.Printf("OnEntityTreeAddNode %s: %v + %s", parentId, children, childId)
								}))
								runtime.BindEventEntityTreeRemoveNode(ctx.GetEntityTree(), runtime.HandleEventEntityTreeRemoveNode(func(entityTree runtime.EntityTree, parentId, childId uid.Id) {
									var children []uid.Id

									entityTree.EachChildren(parentId, func(entity ec.Entity) {
										children = append(children, entity.GetId())
									})

									log.Printf("OnEntityTreeRemoveNode %s: %v - %s", parentId, children, childId)
								}))
								runtime.BindEventEntityTreeMoveNode(ctx.GetEntityTree(), runtime.HandleEventEntityTreeMoveNode(func(entityTree runtime.EntityTree, childId, fromParentId, toParentId uid.Id) {
									log.Printf("OnEntityTreeMoveNode %s: %s => %s", childId, fromParentId, toParentId)
								}))
							case runtime.RunningEvent_Started:
								root, err := core.BuildEntity(ctx, "Test1").New()
								if err != nil {
									log.Panicln("new root error:", err)
								}

								err = ctx.GetEntityTree().MakeRoot(root.GetId())
								if err != nil {
									log.Panicln("make root error:", err)
								}

								log.Println("1. testing child detach in attaching")

								PrintEntityTreeForest(ctx.GetEntityTree())

								child1, err := core.BuildEntity(ctx, "Test2").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child1 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Println("2. testing child remove in attaching")

								PrintEntityTreeForest(ctx.GetEntityTree())

								child2, err := core.BuildEntity(ctx, "Test3").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child2 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Println("3. testing child destroy in attaching")

								PrintEntityTreeForest(ctx.GetEntityTree())

								child3, err := core.BuildEntity(ctx, "Test4").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child3 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Println("4. testing child destroy in detaching")

								PrintEntityTreeForest(ctx.GetEntityTree())

								child4, err := core.BuildEntity(ctx, "Test5").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child4 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								ctx.GetEntityTree().DetachNode(child4.GetId())
								log.Printf("%s: state=%s, tree_node_state=%s", child4.GetId(), child4.GetState(), child4.GetTreeNodeState())

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Println("4. testing parent destroy in attaching")

								PrintEntityTreeForest(ctx.GetEntityTree())

								child5, err := core.BuildEntity(ctx, "Test6").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child5 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								child6, err := core.BuildEntity(ctx, "Test1").SetParentId(child5.GetId()).New()
								if err != nil {
									log.Panicln("new child6 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Printf("%s: state=%s, tree_node_state=%s", child5.GetId(), child5.GetState(), child5.GetTreeNodeState())
								log.Printf("%s: state=%s, tree_node_state=%s", child6.GetId(), child6.GetState(), child6.GetTreeNodeState())

								log.Println("5. testing parent destroy in detaching")

								PrintEntityTreeForest(ctx.GetEntityTree())

								child7, err := core.BuildEntity(ctx, "Test7").SetParentId(root.GetId()).New()
								if err != nil {
									log.Panicln("new child7 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								child8, err := core.BuildEntity(ctx, "Test1").SetParentId(child7.GetId()).New()
								if err != nil {
									log.Panicln("new child8 error:", err)
								}

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Printf("%s: state=%s, tree_node_state=%s", child7.GetId(), child7.GetState(), child7.GetTreeNodeState())
								log.Printf("%s: state=%s, tree_node_state=%s", child8.GetId(), child8.GetState(), child8.GetTreeNodeState())

								ctx.GetEntityTree().DetachNode(child8.GetId())

								PrintEntityTreeForest(ctx.GetEntityTree())

								log.Printf("%s: state=%s, tree_node_state=%s", child7.GetId(), child7.GetState(), child7.GetTreeNodeState())
								log.Printf("%s: state=%s, tree_node_state=%s", child8.GetId(), child8.GetState(), child8.GetTreeNodeState())

								_ = child1
								_ = child2
								_ = child3
								_ = child4
								_ = child5
								_ = child6
								_ = child7
								_ = child8
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enable(false)),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestFrameUpdate struct {
	ec.ComponentBehavior
}

func (c *ComponentTestFrameUpdate) Update() {
	frame := runtime.Current(c).GetFrame()
	log.Printf("Component %s.%s Update, fps: %.2f", c.GetEntity().GetId(), c.GetName(), frame.GetCurFPS())
}

func (c *ComponentTestFrameUpdate) LateUpdate() {
	log.Printf("Component %s.%s LateUpdate", c.GetEntity().GetId(), c.GetName())
}

func Test_CreateEntityFrameUpdate(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestFrameUpdate{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								for range 10 {
									core.BuildEntity(ctx, "Test1").New()
								}
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run()
}

type ComponentTestStressFrameUpdate struct {
	ec.ComponentBehavior
	count int
}

func (c *ComponentTestStressFrameUpdate) Update() {
	c.count++
}

func (c *ComponentTestStressFrameUpdate) LateUpdate() {
	c.count++
}

func Test_CreateEntityStressFrameUpdate(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 120*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestStressFrameUpdate{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_FrameLoopBegin:
								for range 200 {
									core.BuildEntity(ctx, "Test1").New()
								}
							case runtime.RunningEvent_RunGCBegin:
								log.Printf("fps: %.2f, running_elapse_time: %.3f, last_loop_elapse_time: %.3f, entities: %d",
									ctx.GetFrame().GetCurFPS(),
									ctx.GetFrame().GetRunningElapseTime().Seconds(),
									ctx.GetFrame().GetLastLoopElapseTime().Seconds(),
									ctx.GetEntityManager().CountEntities())
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
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
	serviceAddIn1 = define.AddIn(NewServiceAddIn1)
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
	serviceAddIn2 = define.AddIn(NewServiceAddIn2)
)

func Test_ServiceAddIn(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				serviceAddIn1.Install(ctx)
				serviceAddIn2.Install(ctx)
			}
			log.Println("service event:", runningEvent, args)
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
	runtimeAddIn1 = define.AddIn(NewRuntimeAddIn1)
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
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Birth:
								runtimeAddIn1.Install(ctx)
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enable(false)),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run()
}
