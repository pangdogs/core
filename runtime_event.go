//go:generate go run kit.golaxy.org/golaxy/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE --default_export=false --default_auto=false
package golaxy

type eventUpdate interface {
	Update()
}

type eventLateUpdate interface {
	LateUpdate()
}
