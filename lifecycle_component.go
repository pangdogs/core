package golaxy

type LifecycleComponentAwake interface {
	Awake()
}

type LifecycleComponentStart interface {
	Start()
}

type LifecycleComponentUpdate = eventUpdate

type LifecycleComponentLateUpdate = eventLateUpdate

type LifecycleComponentShut interface {
	Shut()
}
