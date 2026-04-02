# CORE
[English](./README.md) | [简体中文](./README.zh_CN.md)

## 简介
`core` 是
[**Golaxy 分布式服务开发框架**](https://github.com/pangdogs/framework)
的执行内核。它把服务级作用域、Actor 风格运行时、实体组件模型、原型库、
本地事件、插件系统和异步辅助统一成一套服务端编程模型。

这个仓库主要面向实时后端场景，例如游戏服务器、仿真/控制系统，以及其他需要
明确运行时归属和消息式执行的服务程序。

## 本模块提供什么
- **服务作用域**：负责启动/停止编排、原型注册、全局实体索引和 service 插件。
- **运行时作用域**：负责 Actor 风格执行、任务队列、可选帧循环、实体树和 runtime 插件。
- **实体组件模型**：支持生命周期驱动的激活、启停、动态增删组件和树形关系管理。
- **实体原型系统**：可在服务启动阶段预先声明可复用的实体/组件组合。
- **本地事件系统**：通过代码生成提供轻量事件绑定、派发和事件表抽象。
- **Async/Future 辅助**：用于把工作调度回所属 runtime，并协调并发结果。

## 运行模型
- `service.Context` 是外层全局作用域，持有实体原型库、service 插件管理器和服务运行事件。
- `runtime.Context` 是 Actor 风格的执行作用域，持有任务队列、可选帧循环、本地实体管理器和实体树。
- 实体和组件都运行在 runtime 中，它们的生命周期由 `core.Runtime` 自动推进，而不是由任意 goroutine 直接修改。
- 跨 goroutine 或跨 runtime 的工作，通常应通过 `CallAsync`、`CallVoidAsync`、`Await`，或者 `service.Context` / `runtime.Context` 提供的调用接口回到目标 runtime。

## 快速开始
[`core_test.go`](./core_test.go) 里的测试基本都是场景化示例。下面这段代码展示了从原型声明到运行时创建实体的最小路径：

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

## 本地事件与代码生成
事件系统推荐通过 `go generate` 使用：

1. 在 `*_event.go` 中定义事件接口。
2. 添加 `//go:generate go run git.golaxy.org/core/event/eventc event`。
3. 如果需要事件表，再添加
   `//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=...`。
4. 执行 `go generate ./...`。

仓库内可直接参考这些文件：
- [`ec/component_event.go`](./ec/component_event.go)
- [`ec/entity_event.go`](./ec/entity_event.go)
- [`runtime/context_event.go`](./runtime/context_event.go)
- [`runtime_event.go`](./runtime_event.go)

## 包说明
| 包 | 说明 |
| --- | --- |
| [`/`](https://github.com/pangdogs/core) | 对外公共入口、生命周期接口、异步辅助，以及 service/runtime 封装。 |
| [`/service`](https://github.com/pangdogs/core/tree/main/service) | 服务上下文、原型访问、全局实体调用和服务运行事件。 |
| [`/runtime`](https://github.com/pangdogs/core/tree/main/runtime) | 运行时上下文、任务调度、帧循环、实体管理器和实体树。 |
| [`/ec`](https://github.com/pangdogs/core/tree/main/ec) | 实体/组件模型、状态机、事件和树节点行为。 |
| [`/ec/pt`](https://github.com/pangdogs/core/tree/main/ec/pt) | 实体/组件原型描述、原型库和构造支持。 |
| [`/define`](https://github.com/pangdogs/core/tree/main/define) | 类型安全的插件定义，统一暴露 `Install`、`Require`、`Lookup` 和 `Uninstall`。 |
| [`/extension`](https://github.com/pangdogs/core/tree/main/extension) | 底层插件管理器、状态对象和插件生命周期事件。 |
| [`/event`](https://github.com/pangdogs/core/tree/main/event) | 本地事件基础设施、句柄、递归策略控制和事件表抽象。 |
| [`/event/eventc`](https://github.com/pangdogs/core/tree/main/event/eventc) | 通过 `go:generate` 使用的事件代码生成器。 |
| [`/utils`](https://github.com/pangdogs/core/tree/main/utils) | 通用工具包集合，例如 `assertion`、`async`、`corectx`、`generic`、`iface`、`meta`、`option`、`types` 和 `uid`。 |

## 安装
当前模块以 [`go.mod`](./go.mod) 中声明的 Go 版本为准。

```bash
go get git.golaxy.org/core@latest
```

## 更多示例
- 仓库内端到端示例：[`core_test.go`](./core_test.go)
- 外部示例工程：[Examples](https://github.com/pangdogs/examples)

## 相关项目
- [Golaxy 分布式服务开发框架](https://github.com/pangdogs/framework)
- [Golaxy 游戏服务器脚手架](https://github.com/pangdogs/scaffold)
