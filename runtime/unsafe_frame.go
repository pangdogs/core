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
func (uf _UnsafeFrame) SetCurFrames(v uint64) {
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
