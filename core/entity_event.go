//go:generate go run github.com/pangdogs/galaxy/core/eventcode --decl_file=$GOFILE --not_import_core gen_emit --package=$GOPACKAGE
package core

// eventEntityDestroySelf [EmitUnExport]
type eventEntityDestroySelf interface {
	onEntityDestroySelf(entity Entity)
}
