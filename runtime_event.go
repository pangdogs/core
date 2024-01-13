//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE --default_export=false --default_auto=false
package core

type eventUpdate interface {
	Update()
}

type eventLateUpdate interface {
	LateUpdate()
}
