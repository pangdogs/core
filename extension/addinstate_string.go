// Code generated by "stringer -type AddInState"; DO NOT EDIT.

package extension

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AddInState_Loaded-0]
	_ = x[AddInState_Active-1]
	_ = x[AddInState_Inactive-2]
}

const _AddInState_name = "AddInState_LoadedAddInState_ActiveAddInState_Inactive"

var _AddInState_index = [...]uint8{0, 18, 36, 56}

func (i AddInState) String() string {
	if i < 0 || i >= AddInState(len(_AddInState_index)-1) {
		return "AddInState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AddInState_name[_AddInState_index[i]:_AddInState_index[i+1]]
}
