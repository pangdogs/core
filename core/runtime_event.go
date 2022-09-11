//go:generate go run github.com/pangdogs/galaxy/core/eventcode --decl_file=$GOFILE --not_import_core gen_emit --package=$GOPACKAGE --default_export=0
package core

// eventUpdate
type eventUpdate interface {
	Update()
}

// eventLateUpdate
type eventLateUpdate interface {
	LateUpdate()
}
