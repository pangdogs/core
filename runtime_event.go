//go:generate go run git.golaxy.org/core/event/eventc event --default_export=false --default_auto=false
package core

type eventUpdate interface {
	Update()
}

type eventLateUpdate interface {
	LateUpdate()
}
