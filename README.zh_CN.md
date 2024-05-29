# CORE
[English](./README.md) | [简体中文](./README.zh_CN.md)

## 简介
[**Golaxy分布式服务开发框架**](https://github.com/pangdogs/framework) 旨在为实时通信应用程序提供一个全面的服务端解决方案。框架基于EC系统与Actor线程模型，设计简洁、易于使用，特别适合用于开发游戏和远程控制系统。

本项目是框架的 [**内核**](https://github.com/pangdogs/core) 部分，主要功能特性包括：

- 实体组件框架（`Entity Component`）：提供灵活的实体和组件管理，支持复杂对象的创建与维护。
- 实体原型系统（`Entity Prototype`）：支持实体的原型定义和复用，简化实体的创建过程。
- Actor线程模型（`Actor Model`）：基于Actor模型的线程处理机制，每个Actor在独立的计算单元，实现并行任务处理，提升系统的并发性能和稳定性。
- 运行时环境（`Runtime and Context`）：实现Actor独立运行线程环境，并提供实体管理与通信调用机制。
- 服务环境（`Service and Context`）：支持服务的启动、停止和管理，提供全局的实体管理与通信调用机制。
- 插件系统（`Plugin Support`）：提供扩展框架功能的机制，支持在运行时环境或服务环境中扩展实现新功能。
- 本地事件系统（`Local Event`）：基于代码生成器，在Actor独立运行线程环境中，提供高效的本地事件机制。
- 异步调用方案（`Async/Await`）：支持异步操作，简化异步代码的编写，提升系统的响应能力。

## 目录
| Directory | Description |
| --------- | ----------- |
| [/](https://github.com/pangdogs/core) | 主要实现服务与运行时相关功能。|
| [/define](https://github.com/pangdogs/core/tree/main/define) | 利用泛型特性，支持定义插件或组件，简化代码编写。 |
| [/ec](https://github.com/pangdogs/core/tree/main/ec) | 实体组件框架。 |
| [/event](https://github.com/pangdogs/core/tree/main/event) | 本地事件系统。 |
| [/event/eventcode](https://github.com/pangdogs/core/tree/main/event/eventcode) | 本地事件代码生成器。 |
| [/plugin](https://github.com/pangdogs/core/tree/main/plugin) | 插件系统。 |
| [/pt](https://github.com/pangdogs/core/tree/main/pt) | 实体原型系统。 |
| [/runtime](https://github.com/pangdogs/core/tree/main/runtime) | 运行时上下文。 |
| [/service](https://github.com/pangdogs/core/tree/main/service) | 服务上下文。 |
| [/utils](https://github.com/pangdogs/core/tree/main/utils) | 一些工具类与函数。 |

## 示例

详见： [Examples](https://github.com/pangdogs/examples)

## 安装
```
go get -u git.golaxy.org/core
```
