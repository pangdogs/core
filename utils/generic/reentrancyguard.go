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

package generic

// ReentrancyGuardBits8 按 8 位操作编号阻止同一调用链中的重复进入。
// 该类型不提供并发保护。
type ReentrancyGuardBits8 Bits8

// Call 在 bit 当前未执行时调用 fun；同一 bit 的重入调用会被忽略。
func (g *ReentrancyGuardBits8) Call(bit int, fun func()) {
	if (*Bits8)(g).Is(bit) {
		return
	}

	(*Bits8)(g).Set(bit, true)
	defer (*Bits8)(g).Set(bit, false)

	fun()
}

// ReentrancyGuardBits16 按 16 位操作编号阻止同一调用链中的重复进入。
// 该类型不提供并发保护。
type ReentrancyGuardBits16 Bits16

// Call 在 bit 当前未执行时调用 fun；同一 bit 的重入调用会被忽略。
func (g *ReentrancyGuardBits16) Call(bit int, fun func()) {
	if (*Bits16)(g).Is(bit) {
		return
	}

	(*Bits16)(g).Set(bit, true)
	defer (*Bits16)(g).Set(bit, false)

	fun()
}

// ReentrancyGuardBits32 按 32 位操作编号阻止同一调用链中的重复进入。
// 该类型不提供并发保护。
type ReentrancyGuardBits32 Bits32

// Call 在 bit 当前未执行时调用 fun；同一 bit 的重入调用会被忽略。
func (g *ReentrancyGuardBits32) Call(bit int, fun func()) {
	if (*Bits32)(g).Is(bit) {
		return
	}

	(*Bits32)(g).Set(bit, true)
	defer (*Bits32)(g).Set(bit, false)

	fun()
}

// ReentrancyGuardBits64 按 64 位操作编号阻止同一调用链中的重复进入。
// 该类型不提供并发保护。
type ReentrancyGuardBits64 Bits64

// Call 在 bit 当前未执行时调用 fun；同一 bit 的重入调用会被忽略。
func (g *ReentrancyGuardBits64) Call(bit int, fun func()) {
	if (*Bits64)(g).Is(bit) {
		return
	}

	(*Bits64)(g).Set(bit, true)
	defer (*Bits64)(g).Set(bit, false)

	fun()
}
