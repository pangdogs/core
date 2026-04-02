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

package extension

import (
	"hash/fnv"
	"reflect"

	"git.golaxy.org/core/utils/types"
)

// GenAddInId 生成插件Id
func GenAddInId(name string) uint64 {
	h := fnv.New64a()
	h.Write(types.String2Bytes(name))
	return h.Sum64()
}

// GenAddInName 生成插件名称
func GenAddInName(addIn any) string {
	addInRT := reflect.TypeOf(addIn)

	for addInRT.Kind() == reflect.Pointer {
		addInRT = addInRT.Elem()
	}

	return types.FullNameRT(addInRT)
}

// GenAddInNameT 生成插件名称
func GenAddInNameT[T any]() string {
	return GenAddInName(types.New[T]())
}
