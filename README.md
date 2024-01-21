# core

GOLAXY分布式服务开发框架内核，本层仅提供线程模型与代码组织框架，不提供例如服务发现（Service Registry）、消息队列与事件驱动架构（MQ and Broker）、分布式锁（Distributed Sync）等分布式服务特性，这些特性将会以官方插件形式提供。
本层提供的主要特性如下:

- 实体组件框架（Entity Component）
- Actor线程模型（Actor Model）
- 运行时环境
- 服务环境
- 实体原型系统
- 插件系统
- 本地事件系统

## Install
```
go get -u git.golaxy.org/core
```
