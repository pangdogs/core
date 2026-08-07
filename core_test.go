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
	"fmt"
	"log"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"git.golaxy.org/core/utils/assertion"
	"git.golaxy.org/core/utils/uid"
	"github.com/elliotchance/pie/v2"

	"git.golaxy.org/core"
	"git.golaxy.org/core/define"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

type coreTestScenario struct {
	ctx     context.Context
	cancel  context.CancelFunc
	timeout time.Duration
	once    sync.Once
	result  chan error
}

func newCoreTestScenario(timeout time.Duration) *coreTestScenario {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return &coreTestScenario{
		ctx:     ctx,
		cancel:  cancel,
		timeout: timeout,
		result:  make(chan error, 1),
	}
}

func (scenario *coreTestScenario) complete(err error) {
	scenario.once.Do(func() {
		scenario.result <- err
		scenario.cancel()
	})
}

func (scenario *coreTestScenario) run(t *testing.T, svcCtx service.Context) {
	t.Helper()
	defer scenario.cancel()

	shutdownTimer := time.NewTimer(scenario.timeout + time.Second)
	defer shutdownTimer.Stop()
	select {
	case <-core.NewService(svcCtx).Run().Done():
	case <-shutdownTimer.C:
		t.Fatalf("service did not stop within %s", scenario.timeout+time.Second)
	}

	select {
	case err := <-scenario.result:
		if err != nil {
			t.Fatal(err)
		}
	default:
		t.Fatalf("test scenario stopped before completion: %v", scenario.ctx.Err())
	}
}

func requireOrdered[T comparable](t *testing.T, got []T, want ...T) {
	t.Helper()

	next := 0
	for _, value := range got {
		if next < len(want) && value == want[next] {
			next++
		}
	}
	if next != len(want) {
		t.Fatalf("unexpected event order:\n got: %v\nwant: %v", got, want)
	}
}

func requireExact[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("unexpected values:\n got: %v\nwant: %v", got, want)
	}
}

func Test_StartService(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)
	var (
		eventsMutex sync.Mutex
		events      []service.RunningEvent
	)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			eventsMutex.Lock()
			events = append(events, runningEvent)
			eventsMutex.Unlock()
			if runningEvent == service.RunningEvent_Started {
				scenario.complete(nil)
			}
		}),
	)

	scenario.run(t, svcCtx)
	eventsMutex.Lock()
	gotEvents := slices.Clone(events)
	eventsMutex.Unlock()
	requireOrdered(t, gotEvents,
		service.RunningEvent_Birth,
		service.RunningEvent_Starting,
		service.RunningEvent_Started,
		service.RunningEvent_Terminating,
		service.RunningEvent_Terminated,
	)
}

type EntityTest1 struct {
	ec.EntityBehavior
}

func (e *EntityTest1) Awake() {
	log.Printf("EntityTest1 %s Awake", e.Id())
}

func (e *EntityTest1) Start() {
	log.Printf("EntityTest1 %s Start", e.Id())
}

func (e *EntityTest1) Shut() {
	log.Printf("EntityTest1 %s Shut", e.Id())
}

func (e *EntityTest1) Dispose() {
	log.Printf("EntityTest1 %s Dispose", e.Id())
}

type EntityTest2 struct {
	ec.EntityBehavior
}

func (e *EntityTest2) Awake() {
	log.Printf("EntityTest2 %s Awake", e.Id())
}

func (e *EntityTest2) Start() {
	log.Printf("EntityTest2 %s Start", e.Id())
}

func (e *EntityTest2) Shut() {
	log.Printf("EntityTest2 %s Shut", e.Id())
}

func (e *EntityTest2) Dispose() {
	log.Printf("EntityTest2 %s Dispose", e.Id())
}

type ComponentTest1 struct {
	ec.ComponentBehavior
}

func (c *ComponentTest1) Awake() {
	log.Printf("Component %s.%s Awake", c.Entity().Id(), c.Name())
}

func (c *ComponentTest1) Start() {
	log.Printf("Component %s.%s Start", c.Entity().Id(), c.Name())
}

func (c *ComponentTest1) Shut() {
	log.Printf("Component %s.%s Shut", c.Entity().Id(), c.Name())
}

func (c *ComponentTest1) Dispose() {
	log.Printf("Component %s.%s Dispose", c.Entity().Id(), c.Name())
}

type ComponentTest2 struct {
	ec.ComponentBehavior
}

func (c *ComponentTest2) Awake() {
	log.Printf("Component %s.%s Awake", c.Entity().Id(), c.Name())
}

func (c *ComponentTest2) Start() {
	log.Printf("Component %s.%s Start", c.Entity().Id(), c.Name())
}

func (c *ComponentTest2) Shut() {
	log.Printf("Component %s.%s Shut", c.Entity().Id(), c.Name())
}

func (c *ComponentTest2) Dispose() {
	log.Printf("Component %s.%s Dispose", c.Entity().Id(), c.Name())
}

type ComponentTest3 struct {
	ec.ComponentBehavior
}

func (c *ComponentTest3) Awake() {
	log.Printf("Component %s.%s Awake", c.Entity().Id(), c.Name())
}

func (c *ComponentTest3) Start() {
	log.Printf("Component %s.%s Start", c.Entity().Id(), c.Name())
}

func (c *ComponentTest3) Shut() {
	log.Printf("Component %s.%s Shut", c.Entity().Id(), c.Name())
}

func (c *ComponentTest3) Dispose() {
	log.Printf("Component %s.%s Dispose", c.Entity().Id(), c.Name())
}

func Test_ServiceRegisterEntityPT(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				ctx.EntityLib().Declare(
					pt.NewEntityDescriptor("Test1").SetInstance(EntityTest1{}),
					ComponentTest1{},
				)
				ctx.EntityLib().Declare(
					pt.EntityDescriptor{
						Prototype: "Test2",
						Instance:  EntityTest2{},
					},
					ComponentTest1{},
					ComponentTest2{},
				)
				ctx.EntityLib().Declare(
					"Test3",
					ComponentTest1{},
					ComponentTest2{},
					ComponentTest3{},
				)
			case service.RunningEvent_Started:
				if got := len(ctx.EntityLib().List()); got != 3 {
					scenario.complete(fmt.Errorf("unexpected entity prototype count: got %d, want 3", got))
					return
				}
				for prototype, componentCount := range map[string]int{"Test1": 1, "Test2": 2, "Test3": 3} {
					entityPT, ok := ctx.EntityLib().Get(prototype)
					if !ok {
						scenario.complete(fmt.Errorf("entity prototype %q was not registered", prototype))
						return
					}
					if got := entityPT.CountComponents(); got != componentCount {
						scenario.complete(fmt.Errorf("entity prototype %q component count: got %d, want %d", prototype, got, componentCount))
						return
					}
				}
				scenario.complete(nil)
			}
		}),
	)

	scenario.run(t, svcCtx)
}

func Test_CreateEntity(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)
	var entities []ec.Entity

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
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
							if runningEvent != runtime.RunningEvent_Started {
								return
							}

							for prototype, componentCount := range map[string]int{"Test1": 1, "Test2": 2, "Test3": 3} {
								entity, err := core.BuildEntity(ctx, prototype).New()
								if err != nil {
									scenario.complete(fmt.Errorf("create entity %q: %w", prototype, err))
									return
								}
								if entity.State() != ec.EntityState_Alive {
									scenario.complete(fmt.Errorf("entity %q state: got %s, want %s", prototype, entity.State(), ec.EntityState_Alive))
									return
								}
								if got := entity.CountComponents(); got != componentCount {
									scenario.complete(fmt.Errorf("entity %q component count: got %d, want %d", prototype, got, componentCount))
									return
								}
								entities = append(entities, entity)
							}
							if got := ctx.EntityManager().CountEntities(); got != len(entities) {
								scenario.complete(fmt.Errorf("runtime entity count: got %d, want %d", got, len(entities)))
								return
							}
							scenario.complete(nil)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)
	for _, entity := range entities {
		if entity.State() != ec.EntityState_Destroyed {
			t.Errorf("entity %s state after shutdown: got %s, want %s", entity.Id(), entity.State(), ec.EntityState_Destroyed)
		}
	}
}

type ComponentTestEnable1 struct {
	ec.ComponentBehavior
	events []string
}

func (c *ComponentTestEnable1) Awake() {
	c.events = append(c.events, "Awake")
	c.SetEnabled(false)
}

func (c *ComponentTestEnable1) OnEnable() {
	c.events = append(c.events, "OnEnable")
}

func (c *ComponentTestEnable1) Start() {
	c.events = append(c.events, "Start")
}

func (c *ComponentTestEnable1) Shut() {
	c.events = append(c.events, "Shut")
}

func (c *ComponentTestEnable1) OnDisable() {
	c.events = append(c.events, "OnDisable")
}

func (c *ComponentTestEnable1) Dispose() {
	c.events = append(c.events, "Dispose")
}

type ComponentTestEnable2 struct {
	ec.ComponentBehavior
	events []string
}

func (c *ComponentTestEnable2) Awake() {
	c.events = append(c.events, "Awake")
}

func (c *ComponentTestEnable2) OnEnable() {
	c.events = append(c.events, "OnEnable")
	c.SetEnabled(false)
}

func (c *ComponentTestEnable2) Start() {
	c.events = append(c.events, "Start")
}

func (c *ComponentTestEnable2) Shut() {
	c.events = append(c.events, "Shut")
}

func (c *ComponentTestEnable2) OnDisable() {
	c.events = append(c.events, "OnDisable")
}

func (c *ComponentTestEnable2) Dispose() {
	c.events = append(c.events, "Dispose")
}

type ComponentTestEnable3 struct {
	ec.ComponentBehavior
	events []string
}

func (c *ComponentTestEnable3) Awake() {
	c.events = append(c.events, "Awake")
}

func (c *ComponentTestEnable3) OnEnable() {
	c.events = append(c.events, "OnEnable")
}

func (c *ComponentTestEnable3) Start() {
	c.events = append(c.events, "Start")
	c.SetEnabled(false)
}

func (c *ComponentTestEnable3) Shut() {
	c.events = append(c.events, "Shut")
}

func (c *ComponentTestEnable3) OnDisable() {
	c.events = append(c.events, "OnDisable")
}

func (c *ComponentTestEnable3) Dispose() {
	c.events = append(c.events, "Dispose")
}

type ComponentTestEnable4 struct {
	ec.ComponentBehavior
	events []string
}

func (c *ComponentTestEnable4) Awake() {
	c.events = append(c.events, "Awake")
}

func (c *ComponentTestEnable4) OnEnable() {
	c.events = append(c.events, "OnEnable")
}

func (c *ComponentTestEnable4) Start() {
	c.events = append(c.events, "Start")
}

func (c *ComponentTestEnable4) Shut() {
	c.events = append(c.events, "Shut")
	c.SetEnabled(false)
}

func (c *ComponentTestEnable4) OnDisable() {
	c.events = append(c.events, "OnDisable")
}

func (c *ComponentTestEnable4) Dispose() {
	c.events = append(c.events, "Dispose")
}

func Test_EntityComponentEnable(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)
	var (
		component1 *ComponentTestEnable1
		component2 *ComponentTestEnable2
		component3 *ComponentTestEnable3
		component4 *ComponentTestEnable4
	)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
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
							if runningEvent != runtime.RunningEvent_Started {
								return
							}
							entity, err := core.BuildEntity(ctx, "Test1").New()
							if err != nil {
								scenario.complete(fmt.Errorf("create entity: %w", err))
								return
							}

							var ok bool
							component1, ok = entity.GetComponent("ComponentTestEnable1").(*ComponentTestEnable1)
							if !ok {
								scenario.complete(fmt.Errorf("ComponentTestEnable1 was not created"))
								return
							}
							component2, ok = entity.GetComponent("ComponentTestEnable2").(*ComponentTestEnable2)
							if !ok {
								scenario.complete(fmt.Errorf("ComponentTestEnable2 was not created"))
								return
							}
							component3, ok = entity.GetComponent("ComponentTestEnable3").(*ComponentTestEnable3)
							if !ok {
								scenario.complete(fmt.Errorf("ComponentTestEnable3 was not created"))
								return
							}
							component4, ok = entity.GetComponent("ComponentTestEnable4").(*ComponentTestEnable4)
							if !ok {
								scenario.complete(fmt.Errorf("ComponentTestEnable4 was not created"))
								return
							}
							scenario.complete(nil)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)
	requireExact(t, component1.events, []string{"Awake", "Dispose"})
	requireExact(t, component2.events, []string{"Awake", "OnEnable", "OnDisable", "Dispose"})
	requireExact(t, component3.events, []string{"Awake", "OnEnable", "Start", "OnDisable", "Shut", "Dispose"})
	requireExact(t, component4.events, []string{"Awake", "OnEnable", "Start", "Shut", "OnDisable", "Dispose"})
}

type ComponentTestDynamic1 struct {
	ec.ComponentBehavior

	test2     *ComponentTest2
	test3     *ComponentTest3
	injectErr error
}

func (c *ComponentTestDynamic1) Awake() {
	log.Printf("Component %s.%s Awake", c.Entity().Id(), c.Name())
}

func (c *ComponentTestDynamic1) Start() {
	log.Printf("Component %s.%s Start", c.Entity().Id(), c.Name())
	c.injectErr = assertion.Inject(c.Entity(), c)
}

func (c *ComponentTestDynamic1) Shut() {
	log.Printf("Component %s.%s Shut", c.Entity().Id(), c.Name())
}

func (c *ComponentTestDynamic1) Dispose() {
	log.Printf("Component %s.%s Dispose", c.Entity().Id(), c.Name())
}

type ComponentTestDynamic2 struct {
	ec.ComponentBehavior

	test2     *ComponentTest2
	test3     *ComponentTest3
	injectErr error
}

func (c *ComponentTestDynamic2) Awake() {
	log.Printf("Component %s.%s Awake", c.Entity().Id(), c.Name())
	c.injectErr = assertion.Inject(c.Entity(), c)
}

func (c *ComponentTestDynamic2) Start() {
	log.Printf("Component %s.%s Start", c.Entity().Id(), c.Name())
}

func (c *ComponentTestDynamic2) Shut() {
	log.Printf("Component %s.%s Shut", c.Entity().Id(), c.Name())
}

func (c *ComponentTestDynamic2) Dispose() {
	log.Printf("Component %s.%s Dispose", c.Entity().Id(), c.Name())
}

func Test_EntityDynamicComponent(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				ctx.EntityLib().ComponentLib().Declare(ComponentTest2{})
				ctx.EntityLib().ComponentLib().Declare(ComponentTest3{})

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
							if runningEvent != runtime.RunningEvent_Started {
								return
							}

							entity1, err := core.BuildEntity(ctx, "Test1").New()
							if err != nil {
								scenario.complete(fmt.Errorf("create Test1: %w", err))
								return
							}
							dynamic1, ok := entity1.GetComponent("ComponentTestDynamic1").(*ComponentTestDynamic1)
							if !ok {
								scenario.complete(fmt.Errorf("ComponentTestDynamic1 was not created"))
								return
							}
							if dynamic1.injectErr != nil {
								scenario.complete(fmt.Errorf("inject Test1 components: %w", dynamic1.injectErr))
								return
							}
							if dynamic1.test2 == nil || dynamic1.test3 == nil {
								scenario.complete(fmt.Errorf("Test1 component injection incomplete: test2=%v, test3=%v", dynamic1.test2, dynamic1.test3))
								return
							}

							entity2, err := core.BuildEntity(ctx, "Test2").New()
							if err != nil {
								scenario.complete(fmt.Errorf("create Test2: %w", err))
								return
							}
							dynamic2, ok := entity2.GetComponent("ComponentTestDynamic2").(*ComponentTestDynamic2)
							if !ok {
								scenario.complete(fmt.Errorf("ComponentTestDynamic2 was not created"))
								return
							}
							if dynamic2.injectErr != nil {
								scenario.complete(fmt.Errorf("inject Test2 components: %w", dynamic2.injectErr))
								return
							}
							if dynamic2.test2 == nil || dynamic2.test3 == nil {
								scenario.complete(fmt.Errorf("Test2 component injection incomplete: test2=%v, test3=%v", dynamic2.test2, dynamic2.test3))
								return
							}

							if got := entity1.CountComponents(); got != 3 {
								scenario.complete(fmt.Errorf("Test1 component count: got %d, want 3", got))
								return
							}
							if got := entity2.CountComponents(); got != 3 {
								scenario.complete(fmt.Errorf("Test2 component count: got %d, want 3", got))
								return
							}
							scenario.complete(nil)
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)
}

type ComponentTestParent struct {
	ec.ComponentBehavior
}

func (c *ComponentTestParent) Awake() {
	ec.BindEventTreeNodeAddChild(c.Entity(), c)
	ec.BindEventTreeNodeRemoveChild(c.Entity(), c)
}

func (c *ComponentTestParent) OnTreeNodeAddChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeAddChild %s <- %s", entity.Id(), childId)
}

func (c *ComponentTestParent) OnTreeNodeRemoveChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeRemoveChild %s x- %s", entity.Id(), childId)
}

type ComponentTestChild struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChild) Awake() {
	ec.BindEventTreeNodeAttachParent(c.Entity(), c)
	ec.BindEventTreeNodeDetachParent(c.Entity(), c)
}

func (c *ComponentTestChild) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.Id(), parentId)
}

func (c *ComponentTestChild) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.Id(), parentId)
}

func PrintEntityTreeForest(entityTree runtime.EntityTree) {
	entityTree.EachChildren(runtime.ForestNodeId, func(entity ec.Entity) {
		PrintEntityTree(entity)
	})
}

func PrintEntityTree(entity ec.Entity, depth ...int) {
	entityTree := runtime.Current(entity).EntityTree()
	if b, _ := entityTree.IsFree(entity.Id()); b {
		return
	}

	root := ""

	isRoot, _ := entityTree.IsRoot(entity.Id())
	if isRoot {
		root = "R"
	}

	leaf := ""

	isLeaf, _ := entityTree.IsLeaf(entity.Id())
	if isLeaf {
		leaf = "L"
	}

	_depth := pie.First(depth)

	if isLeaf {
		log.Printf("%s- [%s] %s%s", strings.Repeat(" ", _depth), entity.Id(), root, leaf)
	} else {
		log.Printf("%s+ [%s] %s%s", strings.Repeat(" ", _depth), entity.Id(), root, leaf)
	}

	entityTree.EachChildren(entity.Id(), func(entity ec.Entity) {
		PrintEntityTree(entity, _depth+1)
	})
}

type entityTreeEventCounts struct {
	added   int
	removed int
	moved   int
}

func newTreeEntity(ctx runtime.Context, prototype string, parentId uid.Id) (ec.Entity, error) {
	entity, err := core.BuildEntity(ctx, prototype).New()
	if err != nil {
		return nil, fmt.Errorf("create %s entity: %w", prototype, err)
	}
	if err := ctx.EntityTree().AddChild(parentId, entity.Id()); err != nil {
		return nil, fmt.Errorf("add %s entity %s under %s: %w", prototype, entity.Id(), parentId, err)
	}
	return entity, nil
}

func checkTreeRoot(tree runtime.EntityTree, entity ec.Entity) error {
	isRoot, err := tree.IsRoot(entity.Id())
	if err != nil {
		return fmt.Errorf("check root %s: %w", entity.Id(), err)
	}
	if !isRoot {
		return fmt.Errorf("entity %s is not a root", entity.Id())
	}
	return nil
}

func checkTreeParent(tree runtime.EntityTree, child, parent ec.Entity) error {
	actual, err := tree.GetParent(child.Id())
	if err != nil {
		return fmt.Errorf("get parent of %s: %w", child.Id(), err)
	}
	if actual.Id() != parent.Id() {
		return fmt.Errorf("parent of %s: got %s, want %s", child.Id(), actual.Id(), parent.Id())
	}
	return nil
}

func checkTreeFree(tree runtime.EntityTree, entity ec.Entity) error {
	isFree, err := tree.IsFree(entity.Id())
	if err != nil {
		return fmt.Errorf("check free entity %s: %w", entity.Id(), err)
	}
	if !isFree || entity.TreeNodeState() != ec.TreeNodeState_Free {
		return fmt.Errorf("entity %s tree state: free=%t, state=%s", entity.Id(), isFree, entity.TreeNodeState())
	}
	return nil
}

func checkTreeChildCount(tree runtime.EntityTree, parentId uid.Id, want int) error {
	got, err := tree.CountChildren(parentId)
	if err != nil {
		return fmt.Errorf("count children of %s: %w", parentId, err)
	}
	if got != want {
		return fmt.Errorf("children of %s: got %d, want %d", parentId, got, want)
	}
	return nil
}

func runEntityTreeScenario(ctx runtime.Context, counts *entityTreeEventCounts) error {
	tree := ctx.EntityTree()
	root, err := newTreeEntity(ctx, "Test1", runtime.ForestNodeId)
	if err != nil {
		return err
	}
	child1, err := newTreeEntity(ctx, "Test1", root.Id())
	if err != nil {
		return err
	}
	child2, err := newTreeEntity(ctx, "Test1", root.Id())
	if err != nil {
		return err
	}
	child3, err := newTreeEntity(ctx, "Test1", child1.Id())
	if err != nil {
		return err
	}
	child4, err := newTreeEntity(ctx, "Test1", child3.Id())
	if err != nil {
		return err
	}
	child5, err := newTreeEntity(ctx, "Test1", child3.Id())
	if err != nil {
		return err
	}
	child6, err := newTreeEntity(ctx, "Test1", child3.Id())
	if err != nil {
		return err
	}
	child7, err := newTreeEntity(ctx, "Test1", runtime.ForestNodeId)
	if err != nil {
		return err
	}
	child8, err := newTreeEntity(ctx, "Test1", child2.Id())
	if err != nil {
		return err
	}

	if err := tree.DetachNode(child2.Id()); err != nil {
		return fmt.Errorf("detach child2: %w", err)
	}
	if err := tree.RemoveNode(child3.Id()); err != nil {
		return fmt.Errorf("remove child3 subtree: %w", err)
	}
	if err := tree.MoveNode(child7.Id(), child2.Id()); err != nil {
		return fmt.Errorf("move child7 under child2: %w", err)
	}

	for _, check := range []func() error{
		func() error { return checkTreeRoot(tree, root) },
		func() error { return checkTreeParent(tree, child1, root) },
		func() error { return checkTreeRoot(tree, child2) },
		func() error { return checkTreeParent(tree, child7, child2) },
		func() error { return checkTreeParent(tree, child8, child2) },
		func() error { return checkTreeChildCount(tree, root.Id(), 1) },
		func() error { return checkTreeChildCount(tree, child1.Id(), 0) },
		func() error { return checkTreeChildCount(tree, child2.Id(), 2) },
		func() error { return checkTreeFree(tree, child3) },
		func() error { return checkTreeFree(tree, child4) },
		func() error { return checkTreeFree(tree, child5) },
		func() error { return checkTreeFree(tree, child6) },
	} {
		if err := check(); err != nil {
			return err
		}
	}

	if counts.added != 9 || counts.removed != 4 || counts.moved != 2 {
		return fmt.Errorf("unexpected entity-tree event counts: added=%d, removed=%d, moved=%d", counts.added, counts.removed, counts.moved)
	}
	return nil
}

func Test_EntityTree(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)
	counts := &entityTreeEventCounts{}

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
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
								runtime.BindEventEntityTreeAddNode(ctx.EntityTree(), runtime.HandleEventEntityTreeAddNode(func(runtime.EntityTree, uid.Id, uid.Id) {
									counts.added++
								}))
								runtime.BindEventEntityTreeRemoveNode(ctx.EntityTree(), runtime.HandleEventEntityTreeRemoveNode(func(runtime.EntityTree, uid.Id, uid.Id) {
									counts.removed++
								}))
								runtime.BindEventEntityTreeMoveNode(ctx.EntityTree(), runtime.HandleEventEntityTreeMoveNode(func(runtime.EntityTree, uid.Id, uid.Id, uid.Id) {
									counts.moved++
								}))
							case runtime.RunningEvent_Started:
								scenario.complete(runEntityTreeScenario(ctx, counts))
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)
}

type ComponentTestChildDetachInAttaching struct {
	ec.ComponentBehavior
	detachErr error
}

func (c *ComponentTestChildDetachInAttaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.Entity(), c)
	ec.BindEventTreeNodeDetachParent(c.Entity(), c)
}

func (c *ComponentTestChildDetachInAttaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.Id(), parentId)
	c.detachErr = runtime.Current(entity).EntityTree().DetachNode(entity.Id())
}

func (c *ComponentTestChildDetachInAttaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.Id(), parentId)
}

type ComponentTestChildRemoveInAttaching struct {
	ec.ComponentBehavior
	removeErr error
}

func (c *ComponentTestChildRemoveInAttaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.Entity(), c)
	ec.BindEventTreeNodeDetachParent(c.Entity(), c)
}

func (c *ComponentTestChildRemoveInAttaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.Id(), parentId)
	c.removeErr = runtime.Current(entity).EntityTree().RemoveNode(entity.Id())
}

func (c *ComponentTestChildRemoveInAttaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.Id(), parentId)
}

type ComponentTestChildDestroyInAttaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChildDestroyInAttaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.Entity(), c)
	ec.BindEventTreeNodeDetachParent(c.Entity(), c)
}

func (c *ComponentTestChildDestroyInAttaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.Id(), parentId)
	entity.Destroy()
}

func (c *ComponentTestChildDestroyInAttaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.Id(), parentId)
}

type ComponentTestChildDestroyInDetaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestChildDestroyInDetaching) Awake() {
	ec.BindEventTreeNodeAttachParent(c.Entity(), c)
	ec.BindEventTreeNodeDetachParent(c.Entity(), c)
}

func (c *ComponentTestChildDestroyInDetaching) OnTreeNodeAttachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeAttachParent %s -> %s", entity.Id(), parentId)
}

func (c *ComponentTestChildDestroyInDetaching) OnTreeNodeDetachParent(entity ec.Entity, parentId uid.Id) {
	log.Printf("OnTreeNodeDetachParent %s -x %s", entity.Id(), parentId)
	entity.Destroy()
}

type ComponentTestParentDestroyInAttaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestParentDestroyInAttaching) Awake() {
	ec.BindEventTreeNodeAddChild(c.Entity(), c)
	ec.BindEventTreeNodeRemoveChild(c.Entity(), c)
}

func (c *ComponentTestParentDestroyInAttaching) OnTreeNodeAddChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeAddChild %s <- %s", entity.Id(), childId)
	entity.Destroy()
}

func (c *ComponentTestParentDestroyInAttaching) OnTreeNodeRemoveChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeRemoveChild %s x- %s", entity.Id(), childId)
}

type ComponentTestParentDestroyInDetaching struct {
	ec.ComponentBehavior
}

func (c *ComponentTestParentDestroyInDetaching) Awake() {
	ec.BindEventTreeNodeAddChild(c.Entity(), c)
	ec.BindEventTreeNodeRemoveChild(c.Entity(), c)
}

func (c *ComponentTestParentDestroyInDetaching) OnTreeNodeAddChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeAddChild %s <- %s", entity.Id(), childId)
}

func (c *ComponentTestParentDestroyInDetaching) OnTreeNodeRemoveChild(entity ec.Entity, childId uid.Id) {
	log.Printf("OnTreeNodeRemoveChild %s x- %s", entity.Id(), childId)
	entity.Destroy()
}

func checkDestroyedTreeEntity(ctx runtime.Context, entity ec.Entity) error {
	if entity.State() != ec.EntityState_Destroyed {
		return fmt.Errorf("entity %s state: got %s, want %s", entity.Id(), entity.State(), ec.EntityState_Destroyed)
	}
	if entity.TreeNodeState() != ec.TreeNodeState_Free {
		return fmt.Errorf("destroyed entity %s tree state: got %s, want %s", entity.Id(), entity.TreeNodeState(), ec.TreeNodeState_Free)
	}
	if _, ok := ctx.EntityManager().GetEntity(entity.Id()); ok {
		return fmt.Errorf("destroyed entity %s is still managed by the runtime", entity.Id())
	}
	return nil
}

func checkAliveFreeTreeEntity(ctx runtime.Context, entity ec.Entity) error {
	if entity.State() != ec.EntityState_Alive {
		return fmt.Errorf("entity %s state: got %s, want %s", entity.Id(), entity.State(), ec.EntityState_Alive)
	}
	return checkTreeFree(ctx.EntityTree(), entity)
}

func runEntityTreeSequenceScenario(ctx runtime.Context) error {
	tree := ctx.EntityTree()
	root, err := newTreeEntity(ctx, "Test1", runtime.ForestNodeId)
	if err != nil {
		return err
	}

	child1, err := newTreeEntity(ctx, "Test2", root.Id())
	if err != nil {
		return err
	}
	detachComponent, ok := child1.GetComponent("ComponentTestChildDetachInAttaching").(*ComponentTestChildDetachInAttaching)
	if !ok {
		return fmt.Errorf("ComponentTestChildDetachInAttaching was not created")
	}
	if detachComponent.detachErr == nil {
		return fmt.Errorf("detaching a child while it is attaching unexpectedly succeeded")
	}
	if err := checkTreeParent(tree, child1, root); err != nil {
		return err
	}

	child2, err := newTreeEntity(ctx, "Test3", root.Id())
	if err != nil {
		return err
	}
	removeComponent, ok := child2.GetComponent("ComponentTestChildRemoveInAttaching").(*ComponentTestChildRemoveInAttaching)
	if !ok {
		return fmt.Errorf("ComponentTestChildRemoveInAttaching was not created")
	}
	if removeComponent.removeErr == nil {
		return fmt.Errorf("removing a child while it is attaching unexpectedly succeeded")
	}
	if err := checkTreeParent(tree, child2, root); err != nil {
		return err
	}

	child3, err := newTreeEntity(ctx, "Test4", root.Id())
	if err != nil {
		return err
	}
	if err := checkDestroyedTreeEntity(ctx, child3); err != nil {
		return fmt.Errorf("destroy child while attaching: %w", err)
	}

	child4, err := newTreeEntity(ctx, "Test5", root.Id())
	if err != nil {
		return err
	}
	if err := tree.DetachNode(child4.Id()); err != nil {
		return fmt.Errorf("detach child that destroys itself: %w", err)
	}
	if err := checkDestroyedTreeEntity(ctx, child4); err != nil {
		return fmt.Errorf("destroy child while detaching: %w", err)
	}

	child5, err := newTreeEntity(ctx, "Test6", root.Id())
	if err != nil {
		return err
	}
	child6, err := newTreeEntity(ctx, "Test1", child5.Id())
	if err != nil {
		return err
	}
	if err := checkDestroyedTreeEntity(ctx, child5); err != nil {
		return fmt.Errorf("destroy parent while attaching child: %w", err)
	}
	if err := checkAliveFreeTreeEntity(ctx, child6); err != nil {
		return fmt.Errorf("child of parent destroyed while attaching: %w", err)
	}

	child7, err := newTreeEntity(ctx, "Test7", root.Id())
	if err != nil {
		return err
	}
	child8, err := newTreeEntity(ctx, "Test1", child7.Id())
	if err != nil {
		return err
	}
	if err := tree.DetachNode(child8.Id()); err != nil {
		return fmt.Errorf("detach child whose parent destroys itself: %w", err)
	}
	if err := checkDestroyedTreeEntity(ctx, child7); err != nil {
		return fmt.Errorf("destroy parent while detaching child: %w", err)
	}
	if err := checkAliveFreeTreeEntity(ctx, child8); err != nil {
		return fmt.Errorf("child of parent destroyed while detaching: %w", err)
	}

	if err := checkTreeChildCount(tree, root.Id(), 2); err != nil {
		return err
	}
	return nil
}

func Test_EntityTreeSequence(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
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
							if runningEvent == runtime.RunningEvent_Started {
								scenario.complete(runEntityTreeSequenceScenario(ctx))
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)
}

type ComponentTestFrameUpdate struct {
	ec.ComponentBehavior
	updates     int
	lateUpdates int
}

func (c *ComponentTestFrameUpdate) Update() {
	c.updates++
}

func (c *ComponentTestFrameUpdate) LateUpdate() {
	c.lateUpdates++
}

func Test_CreateEntityFrameUpdate(t *testing.T) {
	const (
		entityCount = 10
		frameCount  = 3
	)

	scenario := newCoreTestScenario(3 * time.Second)
	components := make([]*ComponentTestFrameUpdate, 0, entityCount)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
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
								for range entityCount {
									entity, err := core.BuildEntity(ctx, "Test1").New()
									if err != nil {
										scenario.complete(fmt.Errorf("create frame-update entity: %w", err))
										return
									}
									component, ok := entity.GetComponent("ComponentTestFrameUpdate").(*ComponentTestFrameUpdate)
									if !ok {
										scenario.complete(fmt.Errorf("ComponentTestFrameUpdate was not created"))
										return
									}
									components = append(components, component)
								}
							case runtime.RunningEvent_Terminated:
								scenario.complete(nil)
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(
						core.With.Frame.TargetFPS(120),
						core.With.Frame.TotalFrames(frameCount),
					),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)
	if len(components) != entityCount {
		t.Fatalf("frame-update component count: got %d, want %d", len(components), entityCount)
	}
	for i, component := range components {
		if component.updates != frameCount || component.lateUpdates != frameCount {
			t.Errorf("component %d callbacks: Update=%d, LateUpdate=%d, want %d each",
				i, component.updates, component.lateUpdates, frameCount)
		}
	}
}

type testEventRecorder struct {
	mutex  sync.Mutex
	events []string
}

func (recorder *testEventRecorder) record(event string) {
	if recorder == nil {
		return
	}
	recorder.mutex.Lock()
	recorder.events = append(recorder.events, event)
	recorder.mutex.Unlock()
}

func (recorder *testEventRecorder) snapshot() []string {
	recorder.mutex.Lock()
	defer recorder.mutex.Unlock()
	return slices.Clone(recorder.events)
}

type ServiceAddIn1 struct {
	recorder *testEventRecorder
}

type IServiceAddIn1 interface {
	Hello()
}

func (addIn *ServiceAddIn1) Init(service.Context) {
	addIn.recorder.record("ServiceAddIn1.Init")
}

func (addIn *ServiceAddIn1) Shut(service.Context) {
	addIn.recorder.record("ServiceAddIn1.Shut")
}

func (addIn *ServiceAddIn1) Hello() {
	addIn.recorder.record("ServiceAddIn1.Hello")
}

func NewServiceAddIn1(settings ...*testEventRecorder) IServiceAddIn1 {
	return &ServiceAddIn1{recorder: pie.First(settings)}
}

var serviceAddIn1 = define.AddIn(NewServiceAddIn1)

type IServiceAddIn2 interface {
	Hello()
}

type ServiceAddIn2 struct {
	recorder *testEventRecorder
}

func (addIn *ServiceAddIn2) Init(service.Context) {
	addIn.recorder.record("ServiceAddIn2.Init")
}

func (addIn *ServiceAddIn2) Shut(service.Context) {
	addIn.recorder.record("ServiceAddIn2.Shut")
}

func (addIn *ServiceAddIn2) Hello() {
	addIn.recorder.record("ServiceAddIn2.Hello")
}

func NewServiceAddIn2(settings ...*testEventRecorder) IServiceAddIn2 {
	return &ServiceAddIn2{recorder: pie.First(settings)}
}

var serviceAddIn2 = define.AddIn(NewServiceAddIn2)

func Test_ServiceAddIn(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)
	recorder := &testEventRecorder{}
	var (
		status1  extension.AddInStatus
		status2  extension.AddInStatus
		setupErr error
	)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				serviceAddIn1.Install(ctx, recorder)
				serviceAddIn2.Install(ctx, recorder)

				var ok bool
				status1, ok = ctx.AddInManager().GetStatusById(serviceAddIn1.Id)
				if !ok {
					setupErr = fmt.Errorf("ServiceAddIn1 status was not registered")
					return
				}
				status2, ok = ctx.AddInManager().GetStatusById(serviceAddIn2.Id)
				if !ok {
					setupErr = fmt.Errorf("ServiceAddIn2 status was not registered")
					return
				}
				if status1.State() != extension.AddInState_Loaded || status2.State() != extension.AddInState_Loaded {
					setupErr = fmt.Errorf("service add-ins were not loaded after installation")
				}
			case service.RunningEvent_Started:
				if setupErr != nil {
					scenario.complete(setupErr)
					return
				}
				if status1.State() != extension.AddInState_Running || status2.State() != extension.AddInState_Running {
					scenario.complete(fmt.Errorf("service add-ins were not running at service start"))
					return
				}

				addIn1, ok := serviceAddIn1.Lookup(ctx)
				if !ok || addIn1 != serviceAddIn1.Require(ctx) {
					scenario.complete(fmt.Errorf("ServiceAddIn1 Lookup and Require disagree"))
					return
				}
				addIn2, ok := serviceAddIn2.Lookup(ctx)
				if !ok || addIn2 != serviceAddIn2.Require(ctx) {
					scenario.complete(fmt.Errorf("ServiceAddIn2 Lookup and Require disagree"))
					return
				}
				addIn1.Hello()
				addIn2.Hello()
				scenario.complete(nil)
			}
		}),
	)

	scenario.run(t, svcCtx)
	if status1.State() != extension.AddInState_Unloaded || status2.State() != extension.AddInState_Unloaded {
		t.Errorf("service add-ins were not unloaded during shutdown: first=%s, second=%s", status1.State(), status2.State())
	}
	if _, ok := serviceAddIn1.Lookup(svcCtx); ok {
		t.Error("ServiceAddIn1 remained available after service shutdown")
	}
	if _, ok := serviceAddIn2.Lookup(svcCtx); ok {
		t.Error("ServiceAddIn2 remained available after service shutdown")
	}
	requireExact(t, recorder.snapshot(), []string{
		"ServiceAddIn1.Init",
		"ServiceAddIn2.Init",
		"ServiceAddIn1.Hello",
		"ServiceAddIn2.Hello",
		"ServiceAddIn2.Shut",
		"ServiceAddIn1.Shut",
	})
}

type RuntimeAddIn1 struct {
	recorder *testEventRecorder
}

type IRuntimeAddIn1 interface {
	Hello()
}

func (addIn *RuntimeAddIn1) Init(runtime.Context) {
	addIn.recorder.record("RuntimeAddIn1.Init")
}

func (addIn *RuntimeAddIn1) Shut(runtime.Context) {
	addIn.recorder.record("RuntimeAddIn1.Shut")
}

func (*RuntimeAddIn1) OnContextRunningEvent(runtime.Context, runtime.RunningEvent, ...any) {}

func (addIn *RuntimeAddIn1) Hello() {
	addIn.recorder.record("RuntimeAddIn1.Hello")
}

func NewRuntimeAddIn1(settings ...*testEventRecorder) IRuntimeAddIn1 {
	return &RuntimeAddIn1{recorder: pie.First(settings)}
}

var runtimeAddIn1 = define.AddIn(NewRuntimeAddIn1)

func Test_RuntimeAddIn(t *testing.T) {
	scenario := newCoreTestScenario(3 * time.Second)
	recorder := &testEventRecorder{}
	var (
		rtCtx    runtime.Context
		status   extension.AddInStatus
		setupErr error
	)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			if runningEvent != service.RunningEvent_Started {
				return
			}

			rtCtx = runtime.NewContext(ctx,
				runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
					switch runningEvent {
					case runtime.RunningEvent_Birth:
						runtimeAddIn1.Install(ctx, recorder)
						var ok bool
						status, ok = ctx.AddInManager().GetStatusById(runtimeAddIn1.Id)
						if !ok {
							setupErr = fmt.Errorf("RuntimeAddIn1 status was not registered")
							return
						}
						if status.State() != extension.AddInState_Loaded {
							setupErr = fmt.Errorf("runtime add-in state after installation: got %s, want %s", status.State(), extension.AddInState_Loaded)
						}
					case runtime.RunningEvent_Started:
						if setupErr != nil {
							scenario.complete(setupErr)
							return
						}
						if status.State() != extension.AddInState_Running {
							scenario.complete(fmt.Errorf("runtime add-in state at runtime start: got %s, want %s", status.State(), extension.AddInState_Running))
							return
						}
						addIn, ok := runtimeAddIn1.Lookup(ctx)
						if !ok || addIn != runtimeAddIn1.Require(ctx) {
							scenario.complete(fmt.Errorf("RuntimeAddIn1 Lookup and Require disagree"))
							return
						}
						addIn.Hello()
						scenario.complete(nil)
					}
				}),
			)
			core.NewRuntime(
				rtCtx,
				core.With.Runtime.AutoRun(true),
				core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
			)
		}),
	)

	scenario.run(t, svcCtx)
	if status.State() != extension.AddInState_Unloaded {
		t.Errorf("runtime add-in state after shutdown: got %s, want %s", status.State(), extension.AddInState_Unloaded)
	}
	if _, ok := runtimeAddIn1.Lookup(rtCtx); ok {
		t.Error("RuntimeAddIn1 remained available after runtime shutdown")
	}
	requireExact(t, recorder.snapshot(), []string{
		"RuntimeAddIn1.Init",
		"RuntimeAddIn1.Hello",
		"RuntimeAddIn1.Shut",
	})
}
