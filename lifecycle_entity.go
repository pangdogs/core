package golaxy

type _EntityInit interface {
	Init()
}

type _EntityStart interface {
	Start()
}

type _EntityUpdate = eventUpdate

type _EntityLateUpdate = eventLateUpdate

type _EntityShut interface {
	Shut()
}

type _EntityDestroy interface {
	Destroy()
}
