package runtime

func UnsafeFrame(frame Frame) _UnsafeFrame {
	return _UnsafeFrame{
		Frame: frame,
	}
}

type _UnsafeFrame struct {
	Frame
}

func (uf _UnsafeFrame) SetCurFrames(v uint64) {
	uf.setCurFrames(v)
}

func (uf _UnsafeFrame) RunningBegin() {
	uf.runningBegin()
}

func (uf _UnsafeFrame) RunningEnd() {
	uf.runningEnd()
}

func (uf _UnsafeFrame) FrameBegin() {
	uf.frameBegin()
}

func (uf _UnsafeFrame) FrameEnd() {
	uf.frameEnd()
}

func (uf _UnsafeFrame) UpdateBegin() {
	uf.updateBegin()
}

func (uf _UnsafeFrame) UpdateEnd() {
	uf.updateEnd()
}
