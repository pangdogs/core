// Package galaxy Galaxy服务器应用开发框架，基于Actor编程模型（Actor Model），使用EC组件框架（Entity Component）组织代码结构，可以像积木一样快速搭建游戏服务器应用。
//   - 框架使用EC组件框架组织代码结构。
//   - 框架基于Actor编程模型，实体就是Actor，其中组件用于实现状态（state）与行为（behavior），运行时中的任务处理流水线就是邮箱（mailbox），
//     实体的ID就是邮箱地址(mailbox address)，服务上下文提供的全局实体管理功能可以用于投递邮件（mail）给Actor。不同于传统Actor编程模型的地方是
//     多个Actor可以使用同个邮箱，也就是多个实体可以加入同个运行时，在同个线程中运行。
//   - 服务上下文，主要有全局实体管理功能、获取分布式工具链几项功能。
//   - 提供服务
//   - 提供运行时，实现了串行化的任务处理流水线，为逻辑层代码提供单线程运行环境。
package galaxy
