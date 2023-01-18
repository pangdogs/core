package runtime

import "github.com/golaxy-kit/golaxy/localevent"

func UnsafeEntityMgr(entityMgr IEntityMgr) _UnsafeEntityMgr {
	return _UnsafeEntityMgr{
		IEntityMgr: entityMgr,
	}
}

type _UnsafeEntityMgr struct {
	IEntityMgr
}

func (u _UnsafeEntityMgr) EventEntityMgrRemovingEntity() localevent.IEvent {
	return u.eventEntityMgrRemovingEntity()
}
