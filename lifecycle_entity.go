package galaxy

type _EntityInit interface {
	Init()
}

type _EntityInitFin interface {
	InitFin()
}

type _EntityUpdate = eventUpdate

type _EntityLateUpdate = eventLateUpdate

type _EntityShut interface {
	Shut()
}

type _EntityShutFin interface {
	ShutFin()
}
