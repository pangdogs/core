//go:generate go run github.com/golaxy-kit/golaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE --default_export=0
package golaxy

type eventUpdate interface {
	Update()
}

type eventLateUpdate interface {
	LateUpdate()
}
