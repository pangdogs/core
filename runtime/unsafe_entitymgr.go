package runtime

import "github.com/galaxy-kit/galaxy/localevent"

func UnsafeEntityMgr(entityMgr IEntityMgr) _UnsafeEntityMgr {
	return _UnsafeEntityMgr{
		IEntityMgr: entityMgr,
	}
}

type _UnsafeEntityMgr struct {
	IEntityMgr
}

func (u _UnsafeEntityMgr) EventEntityMgrNotifyECTreeRemoveEntity() localevent.IEvent {
	return u.eventEntityMgrNotifyECTreeRemoveEntity()
}
