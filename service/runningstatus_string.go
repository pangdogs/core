// Code generated by "stringer -type RunningStatus"; DO NOT EDIT.

package service

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RunningStatus_Birth-0]
	_ = x[RunningStatus_Starting-1]
	_ = x[RunningStatus_Started-2]
	_ = x[RunningStatus_Terminating-3]
	_ = x[RunningStatus_Terminated-4]
	_ = x[RunningStatus_ActivatingAddIn-5]
	_ = x[RunningStatus_AddInActivated-6]
	_ = x[RunningStatus_DeactivatingAddIn-7]
	_ = x[RunningStatus_AddInDeactivated-8]
	_ = x[RunningStatus_EntityPTDeclared-9]
	_ = x[RunningStatus_EntityPTRedeclared-10]
	_ = x[RunningStatus_EntityPTUndeclared-11]
}

const _RunningStatus_name = "RunningStatus_BirthRunningStatus_StartingRunningStatus_StartedRunningStatus_TerminatingRunningStatus_TerminatedRunningStatus_ActivatingAddInRunningStatus_AddInActivatedRunningStatus_DeactivatingAddInRunningStatus_AddInDeactivatedRunningStatus_EntityPTDeclaredRunningStatus_EntityPTRedeclaredRunningStatus_EntityPTUndeclared"

var _RunningStatus_index = [...]uint16{0, 19, 41, 62, 87, 111, 140, 168, 199, 229, 259, 291, 323}

func (i RunningStatus) String() string {
	if i < 0 || i >= RunningStatus(len(_RunningStatus_index)-1) {
		return "RunningStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RunningStatus_name[_RunningStatus_index[i]:_RunningStatus_index[i+1]]
}
