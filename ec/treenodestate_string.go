// Code generated by "stringer -type TreeNodeState"; DO NOT EDIT.

package ec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TreeNodeState_Freedom-0]
	_ = x[TreeNodeState_Attaching-1]
	_ = x[TreeNodeState_Attached-2]
	_ = x[TreeNodeState_Detaching-3]
}

const _TreeNodeState_name = "TreeNodeState_FreedomTreeNodeState_AttachingTreeNodeState_AttachedTreeNodeState_Detaching"

var _TreeNodeState_index = [...]uint8{0, 21, 44, 66, 89}

func (i TreeNodeState) String() string {
	if i < 0 || i >= TreeNodeState(len(_TreeNodeState_index)-1) {
		return "TreeNodeState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TreeNodeState_name[_TreeNodeState_index[i]:_TreeNodeState_index[i+1]]
}
