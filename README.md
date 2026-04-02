# CORE
[English](./README.md) | [简体中文](./README.zh_CN.md)

## Overview
`core` is the execution kernel of the
[**Golaxy Distributed Service Development Framework**](https://github.com/pangdogs/framework).
It combines a service scope, actor-like runtime scopes, an entity-component
model, prototype libraries, local events, add-ins, and async helpers into one
coherent server-side programming model.

The repository is aimed at realtime backend scenarios such as game servers,
simulation/control systems, and other services that benefit from explicit
runtime ownership and message-style execution.

## What This Module Provides
- A **service scope** for startup/shutdown orchestration, prototype
  registration, global entity lookup, and service add-ins.
- A **runtime scope** for actor-style execution, task queues, optional frame
  loops, entity trees, and runtime add-ins.
- An **entity-component model** with lifecycle-driven activation, enable/disable
  flow, dynamic component changes, and tree relationships.
- An **entity prototype system** for declaring reusable entity/component
  compositions before runtimes start creating instances.
- A **local event system** driven by code generation, used heavily throughout
  the framework and available to application code.
- **Async/Future helpers** for scheduling work back onto the owning runtime and
  coordinating concurrent results.

## Working Model
- `service.Context` is the outer, global scope. It owns the entity prototype
  library, the service add-in manager, and service-level running events.
- `runtime.Context` is an actor-like execution scope. It owns a task queue, an
  optional frame loop, the local entity manager, the entity tree, and runtime
  running events.
- Entities and components live inside a runtime. Their lifecycle is advanced by
  `core.Runtime`, not by ad-hoc goroutine access.
- Cross-goroutine or cross-runtime work should usually be marshaled back through
  `CallAsync`, `CallVoidAsync`, `Await`, or the caller APIs on `service.Context`
  and `runtime.Context`.

## Quick Start
The tests in [`core_test.go`](./core_test.go) are scenario-style examples. The
snippet below shows the minimum path for declaring a prototype and creating an
entity inside a runtime:

```go
package main

import (
	"context"
	"log"
	"time"

	"git.golaxy.org/core"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

type Player struct {
	ec.ComponentBehavior
}

func (p *Player) Awake() {
	log.Printf("player %s awake", p.Entity().Id())
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, event service.RunningEvent, args ...any) {
			switch event {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "player").
					AddComponent(Player{}).
					Declare()

			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(
						ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, event runtime.RunningEvent, args ...any) {
							if event == runtime.RunningEvent_Started {
								_, _ = core.BuildEntity(ctx, "player").New()
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
					core.With.Runtime.Frame(core.With.Frame.Enabled(false)),
				)
			}
		}),
	)

	<-core.NewService(svcCtx).Run().Done()
}
```

## Local Events and Code Generation
The event system is meant to be used through `go generate`.

1. Define an event interface in a `*_event.go` file.
2. Add `//go:generate go run git.golaxy.org/core/event/eventc event`.
3. If you need an event table, also add
   `//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=...`.
4. Run `go generate ./...`.

In-repo references:
- [`ec/component_event.go`](./ec/component_event.go)
- [`ec/entity_event.go`](./ec/entity_event.go)
- [`runtime/context_event.go`](./runtime/context_event.go)
- [`runtime_event.go`](./runtime_event.go)

## Package Guide
| Package | Responsibility |
| --- | --- |
| [`/`](https://github.com/pangdogs/core) | Public entry points, lifecycle interfaces, async helpers, and runtime/service wrappers. |
| [`/service`](https://github.com/pangdogs/core/tree/main/service) | Service context, prototype access, global entity calls, and service running events. |
| [`/runtime`](https://github.com/pangdogs/core/tree/main/runtime) | Runtime context, task scheduling, frame loop, entity manager, and entity tree. |
| [`/ec`](https://github.com/pangdogs/core/tree/main/ec) | Entity/component model, state machines, events, and tree-node behavior. |
| [`/ec/pt`](https://github.com/pangdogs/core/tree/main/ec/pt) | Entity/component prototype descriptors, libraries, and construction support. |
| [`/define`](https://github.com/pangdogs/core/tree/main/define) | Typed add-in definitions with `Install`, `Require`, `Lookup`, and `Uninstall` helpers. |
| [`/extension`](https://github.com/pangdogs/core/tree/main/extension) | Low-level add-in managers, status objects, and add-in lifecycle events. |
| [`/event`](https://github.com/pangdogs/core/tree/main/event) | Local event primitives, handles, recursion control, and event-table abstractions. |
| [`/event/eventc`](https://github.com/pangdogs/core/tree/main/event/eventc) | Event code generator used through `go:generate`. |
| [`/utils`](https://github.com/pangdogs/core/tree/main/utils) | Shared helper packages such as `assertion`, `async`, `corectx`, `generic`, `iface`, `meta`, `option`, `types`, and `uid`. |

## Installation
The module currently targets the Go version declared in [`go.mod`](./go.mod).

```bash
go get git.golaxy.org/core@latest
```

## More Examples
- In-repo end-to-end scenarios: [`core_test.go`](./core_test.go)
- External samples: [Examples](https://github.com/pangdogs/examples)

## Related Repositories
- [Golaxy Distributed Service Development Framework](https://github.com/pangdogs/framework)
- [Golaxy Developing a Game Server Scaffold](https://github.com/pangdogs/scaffold)
