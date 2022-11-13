//go:generate go run github.com/galaxy-kit/galaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE --default_export=0
package galaxy

type eventUpdate interface {
	Update()
}

type eventLateUpdate interface {
	LateUpdate()
}
