# CORE
[English](./README.md) | [简体中文](./README.zh_CN.md)

## Introduction
The [**Golaxy Distributed Service Development Framework**](https://github.com/pangdogs/framework) aims to provide a comprehensive server-side solution for real-time communication applications. Based on the EC system and Actor thread model, the framework is designed to be simple and easy to use, making it particularly suitable for developing games and remote control systems.

This project is the [**core**](https://github.com/pangdogs/core) part of the framework, with the main features including:

- **Entity Component Framework (`Entity Component`)**: Provides flexible entity and component management, supporting the creation and maintenance of complex objects.
- **Entity Prototype System (`Entity Prototype`)**: Supports the definition and reuse of entity prototypes, simplifying the creation process of entities.
- **Actor Thread Model (`Actor Model`)**: Based on the Actor model's thread processing mechanism, each Actor runs in an independent computing unit, achieving parallel task processing and enhancing the system's concurrency performance and stability.
- **Runtime Environment (`Runtime and Context`)**: Implements an independent runtime thread environment for Actors, providing mechanisms for entity management and communication calls.
- **Service Environment (`Service and Context`)**: Supports the startup, shutdown, and management of services, offering global entity management and communication call mechanisms.
- **AddIn System (`Add-In System`)**: Provides mechanisms for extending framework functions, supporting the implementation of new features in the runtime or service environments.
- **Local Event System (`Local Event`)**: Based on a code generator, it provides an efficient local event mechanism within the independent runtime thread environment of Actors.
- **Asynchronous Call Scheme (`Async/Await`)**: Supports asynchronous operations, simplifying the writing of asynchronous code, and enhancing the system's responsiveness.

## Directory
| Directory                                                                | Description                                                                               |
|--------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| [/](https://github.com/pangdogs/core)                                    | Main implementation of service and runtime related functionalities.                       |
| [/define](https://github.com/pangdogs/core/tree/main/define)             | Supports the definition of addins or components using generics, simplifying code writing. |
| [/ec](https://github.com/pangdogs/core/tree/main/ec)                     | Entity Component Framework.                                                               |
| [/ec/pt](https://github.com/pangdogs/core/tree/main/ec/pt)               | Entity Prototype System.                                                                  |
| [/event](https://github.com/pangdogs/core/tree/main/event)               | Local Event System.                                                                       |
| [/event/eventc](https://github.com/pangdogs/core/tree/main/event/eventc) | Local Event Code Generator.                                                               |
| [/extension](https://github.com/pangdogs/core/tree/main/extension)       | Add-In System.                                                                            |
| [/runtime](https://github.com/pangdogs/core/tree/main/runtime)           | Runtime Context.                                                                          |
| [/service](https://github.com/pangdogs/core/tree/main/service)           | Service Context.                                                                          |
| [/utils](https://github.com/pangdogs/core/tree/main/utils)               | Various utility classes and functions.                                                    |

## Examples

For more details, see: [Examples](https://github.com/pangdogs/examples)

## Installation
```
go get -u git.golaxy.org/core
```

## Associated Repositories
- [Golaxy Distributed Service Development Framework](https://github.com/pangdogs/framework)
- [Golaxy Developing a Game Server Scaffold](https://github.com/pangdogs/scaffold)