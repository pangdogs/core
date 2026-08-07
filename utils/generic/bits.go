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

// Bits8 是 8 位标志集合；bit 应位于 [0, 7]。
type Bits8 uint8

// Is 报告指定标志位是否置位。
func (bits Bits8) Is(bit int) bool {
	return (bits)&(1<<bit) != 0
}

// Set 设置或清除指定标志位。
func (bits *Bits8) Set(bit int, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}

// Bits16 是 16 位标志集合；bit 应位于 [0, 15]。
type Bits16 uint16

// Is 报告指定标志位是否置位。
func (bits Bits16) Is(bit int) bool {
	return (bits)&(1<<bit) != 0
}

// Set 设置或清除指定标志位。
func (bits *Bits16) Set(bit int, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}

// Bits32 是 32 位标志集合；bit 应位于 [0, 31]。
type Bits32 uint32

// Is 报告指定标志位是否置位。
func (bits Bits32) Is(bit int) bool {
	return (bits)&(1<<bit) != 0
}

// Set 设置或清除指定标志位。
func (bits *Bits32) Set(bit int, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}

// Bits64 是 64 位标志集合；bit 应位于 [0, 63]。
type Bits64 uint64

// Is 报告指定标志位是否置位。
func (bits Bits64) Is(bit int) bool {
	return (bits)&(1<<bit) != 0
}

// Set 设置或清除指定标志位。
func (bits *Bits64) Set(bit int, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}
