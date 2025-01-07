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

package types

type Bits8 uint8

func (bits *Bits8) Is(bit int8) bool {
	return (*bits)&(1<<bit) != 0
}

func (bits *Bits8) Set(bit int8, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}

type Bits16 uint16

func (bits *Bits16) Is(bit int8) bool {
	return (*bits)&(1<<bit) != 0
}

func (bits *Bits16) Set(bit int8, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}

type Bits32 uint8

func (bits *Bits32) Is(bit int8) bool {
	return (*bits)&(1<<bit) != 0
}

func (bits *Bits32) Set(bit int8, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}

type Bits64 uint8

func (bits *Bits64) Is(bit int8) bool {
	return (*bits)&(1<<bit) != 0
}

func (bits *Bits64) Set(bit int8, b bool) {
	if b {
		*bits |= 1 << bit
	} else {
		*bits &= ^(1 << bit)
	}
}
