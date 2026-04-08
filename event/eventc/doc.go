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

// Command eventc 是 Golaxy 的事件代码生成工具。
/*
Command eventc 基于 go:generate 工作，用于从事件声明源文件中自动生成事件相关的辅助代码。
它面向两类生成目标：

1. 事件辅助代码

   由 `event` 子命令生成，主要包含事件绑定函数、事件触发函数，以及将普通函数适配为
   事件处理器接口的辅助类型。

2. 事件表代码

   由 `eventtab` 子命令生成，主要包含事件表接口、事件表结构，以及访问各个事件实例的
   入口方法。

基本用法

通常在声明事件的 Go 源文件顶部添加 go:generate 指令，例如：

	//go:generate go run git.golaxy.org/core/event/eventc event
	//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=componentEventTab

随后在该文件所在目录执行 `go generate`，即可生成对应代码；也可以通过 IDE 提供的
go generate 集成功能触发。

事件识别规则

eventc 会扫描 `--decl_file` 指定的 Go 源文件，并按以下规则识别事件：

- 默认使用正则 `^[eE]vent.+` 匹配类型名，只有名称匹配的类型声明才会继续处理。
- 事件定义必须是 `interface` 类型。
- 当前实现只读取接口中的第一个具名方法，并将该方法视为事件回调签名。
- 如果该方法存在返回值，则只支持单个 `bool` 返回值；其他返回形式不会生成代码。
- 带类型参数的泛型事件定义当前不会参与生成，eventc 会直接跳过这类声明。

event 子命令生成内容

`event` 子命令会为每个符合条件的事件生成辅助代码，通常包括：

- `BindXxx`：绑定订阅者。
- `_EmitXxx` 或 `EmitXxx`：触发事件。
- `_EmitXxxWithInterrupt` 或 `EmitXxxWithInterrupt`：支持中断判断的触发函数。
- `HandleXxx` 与 `XxxHandler`：将普通函数适配为事件接口实现。

其中 `Emit` 函数是否导出、是否生成 auto 风格的辅助函数，可以通过注释指令或命令行
默认参数共同控制。

eventtab 子命令生成内容

`eventtab` 子命令会生成事件表相关代码，用于集中管理一组事件实例。生成结果通常包括：

- 事件表接口。
- 事件表结构体。
- 各事件的访问方法。
- 事件递归策略的初始化逻辑。

注释指令

event 子命令支持在事件注释中声明：

	+event-gen:export_emit=[0,1]&auto=[0,1]

含义如下：

- `export_emit`：控制生成的事件触发函数是否导出。
- `auto`：控制是否生成基于宿主对象自动访问事件实例的辅助代码。

eventtab 子命令支持在事件注释中声明：

	+event-tab-gen:recursion=[allow,disallow,discard,skip_received,receive_once]

该选项用于控制事件表中对应事件的递归分发策略。

输出文件规则

`event` 子命令会在事件声明文件所在目录输出生成文件，命名规则为：

- 先取声明文件名并去掉 `.go` 后缀。
- 如果文件名不以 `_event` 结尾，则补上 `_event`。
- 最终追加 `.gen.go`。

例如 `component.go` 会生成 `component_event.gen.go`，`component_event.go` 也会生成
`component_event.gen.go`。

`eventtab` 子命令的命名规则与上面类似，但最终后缀为 `tab.gen.go`，并且输出目录由
`--dir` 参数决定。例如 `component.go` 默认会生成 `component_eventtab.gen.go`。

常用命令行参数

可以通过以下命令查看完整帮助：

	eventc --help
	eventc event --help
	eventc eventtab --help

常用参数包括：

- `--decl_file`：指定事件声明源文件。
- `--event_regexp`：指定匹配事件类型名的正则表达式。
- `--package_event_alias`：指定生成代码中导入 `event` 包时使用的别名。
- `event --default_export_emit`：设置默认的 Emit 函数导出策略。
- `event --default_auto`：设置是否默认生成 auto 风格辅助代码。
- `eventtab --package`、`--dir`、`--name`：分别控制事件表的包名、输出目录和类型名。

注意事项

- eventc 只负责读取源文件并生成辅助代码，不会改写事件声明本身。
- 若生成文件已存在，重新执行会覆盖同名输出文件，因此应将生成文件视为派生物。
- 包文档描述的是当前源码实现行为；若后续实现发生变化，应以实际源码为准。
*/
package main
