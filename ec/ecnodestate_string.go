// Code generated by "stringer -type ECNodeState"; DO NOT EDIT.

package ec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ECNodeState_Detached-0]
	_ = x[ECNodeState_Attached-1]
	_ = x[ECNodeState_Detaching-2]
}

const _ECNodeState_name = "ECNodeState_DetachedECNodeState_AttachedECNodeState_Detaching"

var _ECNodeState_index = [...]uint8{0, 20, 40, 61}

func (i ECNodeState) String() string {
	if i < 0 || i >= ECNodeState(len(_ECNodeState_index)-1) {
		return "ECNodeState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ECNodeState_name[_ECNodeState_index[i]:_ECNodeState_index[i+1]]
}
