package core

type _ComponentAwake interface {
	Awake()
}

type _ComponentStart interface {
	Start()
}

type _ComponentUpdate = eventUpdate

type _ComponentLateUpdate = eventLateUpdate

type _ComponentShut interface {
	Shut()
}
