//go:generate go run github.com/pangdogs/galaxy/core/eventcode --decl_file=$GOFILE --not_import_core gen_emit --package=$GOPACKAGE
package core

// eventComponentDestroySelf [EmitUnExport]
type eventComponentDestroySelf interface {
	onComponentDestroySelf(comp Component)
}
