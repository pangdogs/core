package golaxy

type LifecycleEntityInit interface {
	Init()
}

type LifecycleEntityInited interface {
	Inited()
}

type LifecycleEntityUpdate = eventUpdate

type LifecycleEntityLateUpdate = eventLateUpdate

type LifecycleEntityShut interface {
	Shut()
}

type LifecycleEntityDestroy interface {
	Destroy()
}
