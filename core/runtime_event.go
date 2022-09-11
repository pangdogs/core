//go:generate go run github.com/pangdogs/galaxy/core/eventcode --decl_file=$GOFILE --not_import_core gen_emit --package=$GOPACKAGE
package core

// eventUpdate [EmitUnExport]
type eventUpdate interface {
	Update()
}

// eventLateUpdate [EmitUnExport]
type eventLateUpdate interface {
	LateUpdate()
}
