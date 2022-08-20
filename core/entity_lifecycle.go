package core

type EntityInit interface {
	Init()
}

type EntityInitFin interface {
	InitFin()
}

type EntityUpdate = eventUpdate

type EntityLateUpdate = eventLateUpdate

type EntityShut interface {
	Shut()
}

type EntityShutFin interface {
	ShutFin()
}
