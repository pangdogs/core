// Code generated by "stringer -type Scope"; DO NOT EDIT.

package ec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Scope_Local-0]
	_ = x[Scope_Global-1]
}

const _Scope_name = "Scope_LocalScope_Global"

var _Scope_index = [...]uint8{0, 11, 23}

func (i Scope) String() string {
	if i < 0 || i >= Scope(len(_Scope_index)-1) {
		return "Scope(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Scope_name[_Scope_index[i]:_Scope_index[i+1]]
}
