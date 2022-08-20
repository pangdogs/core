package core

type ComponentAwake interface {
	Awake()
}

type ComponentStart interface {
	Start()
}

type ComponentUpdate = eventUpdate

type ComponentLateUpdate = eventLateUpdate

type ComponentShut interface {
	Shut()
}
