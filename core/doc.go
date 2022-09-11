// Package core 银河（Galaxy）框架核心部分。
//	- 提供EC组件（Entity Component）框架，用于组织代码结构。
//	- 提供运行时上下文（Runtime Context）与运行时（Runtime），用于管理实体（Entity），提供串行化的任务处理流水线，为逻辑层提供单线程运行环境。
//	- 提供服务上下文（Service Context）与服务（Service），提供全局实体（Entity）管理功能，结合运行时（Runtime）提供的串行化的任务处理流水线，可以在多线程环境下安全的访问实体。
//	- 提供高效的本地单线程事件系统。
package core
