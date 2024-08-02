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

package runtime

// Deprecated: UnsafeFrame 访问帧内部方法
func UnsafeFrame(frame Frame) _UnsafeFrame {
	return _UnsafeFrame{
		Frame: frame,
	}
}

type _UnsafeFrame struct {
	Frame
}

// SetCurFrames 设置当前帧号
func (uf _UnsafeFrame) SetCurFrames(v int64) {
	uf.setCurFrames(v)
}

// RunningBegin 开始运行
func (uf _UnsafeFrame) RunningBegin() {
	uf.runningBegin()
}

// RunningEnd 运行结束
func (uf _UnsafeFrame) RunningEnd() {
	uf.runningEnd()
}

// LoopBegin 开始帧循环
func (uf _UnsafeFrame) LoopBegin() {
	uf.loopBegin()
}

// LoopEnd 帧循环结束
func (uf _UnsafeFrame) LoopEnd() {
	uf.loopEnd()
}

// UpdateBegin 开始帧更新
func (uf _UnsafeFrame) UpdateBegin() {
	uf.updateBegin()
}

// UpdateEnd 帧更新结束
func (uf _UnsafeFrame) UpdateEnd() {
	uf.updateEnd()
}
