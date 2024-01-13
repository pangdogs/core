// package core GOLAXY分布式服务开发框架内核，基于Actor编程模型（Actor Model），使用EC组件框架（Entity-Component）组织代码结构，可以像积木一样快速搭建服务器应用。
/*
   - 使用EC组件框架（Entity-Component）组织代码结构。
   - 并发模式基于Actor编程模型，实体（Entity）就是Actor，其中组件（Component）用于实现状态（state）与行为（behavior），运行时（Runtime）中的任务处理流水线就是邮箱（mailbox），
     实体的Id就是邮箱地址(mailbox address)，服务上下文（Service context）提供的全局实体管理功能，可以用于投递邮件（mail）给Actor。不同于传统Actor编程模型的地方是
     多个Actor可以使用同个邮箱，也就是多个实体可以加入同个运行时，在同个线程中运行。
   - 一些分布式系统常用的依赖项，例如Service Registry、Message Queue、Sync Locker、Gate等，将以插件（plugin）形式提供，导入"kit.golaxy.org/plugins"包即可使用。
     同时也可以参考教程和代码，自己编写插件。
   - 框架对长连接、有状态、无状态和分布式特性支持比较完备，适合开发一些对实时性要求较高的APP服务器，例如游戏服务器、远程控制系统服务器。也可以接DHT网络，开发一些分布式应用，例如分布式文件存储、分布式聊天系统等等。
*/
package core
