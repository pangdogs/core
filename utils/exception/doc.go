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

// Package exception 提供带调用位置信息的错误与 panic 辅助。
/*
Package exception 用于统一构造框架内部错误。它会把调用文件和行号附加到错误信息中，
并可通过 TraceStack 为 error 附带当前栈信息。

框架内的参数检查与不变量校验通常通过 Panic / Panicf 抛出；需要向上返回 error
时则使用 Error / Errorf 系列函数。
*/
package exception
