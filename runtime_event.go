//go:generate go run git.golaxy.org/core/event/eventcode gen_event --default_export=false --default_auto=false
package core

type eventUpdate interface {
	Update()
}

type eventLateUpdate interface {
	LateUpdate()
}
