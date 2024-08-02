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

package uid

import "github.com/segmentio/ksuid"

var (
	// Nil is a nil id.
	Nil Id = ""

	// New generates a new id.
	New = func() Id {
		return Id(ksuid.New().String())
	}

	// From generate id from string.
	From = func(str string) Id {
		return Id(str)
	}
)

// Id represents a global unique id.
type Id string

// IsNil checks if an Id is nil.
func (id Id) IsNil() bool {
	return id == Nil
}

// String implements fmt.Stringer
func (id Id) String() string {
	return string(id)
}
