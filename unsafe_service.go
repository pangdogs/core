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

package core

// Deprecated: UnsafeService 暴露服务内部能力，仅供框架集成代码使用。
func UnsafeService(service Service) _UnsafeService {
	return _UnsafeService{
		Service: service,
	}
}

type _UnsafeService struct {
	Service
}

// Options 返回服务当前使用的选项。
func (u _UnsafeService) Options() *ServiceOptions {
	return u.getOptions()
}

// Instance 返回服务的实际实例。
func (u _UnsafeService) Instance() Service {
	return u.getInstance()
}
