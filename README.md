# core

Golaxy分布式服务开发框架的内核，本包仅提供线程模型与代码组织框架，不提供例如服务发现（Service Discovery）、消息队列与事件驱动架构（MQ and Broker）、分布式锁（Distributed Sync）等分布式服务特性，这些特性将会以官方插件形式提供。

本包提供的主要特性如下:

- 实体组件框架（Entity Component）
- Actor线程模型（Actor Model）
- 运行时环境（Runtime and Context）
- 服务环境（Service and Context）
- 实体原型系统（Entity Prototype）
- 插件系统（Plugin Support）
- 本地事件系统（Local Event）
- 异步调用方案（Async/Await）

## Install
```
go get -u git.golaxy.org/core
```
