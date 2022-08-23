package core

type EntityQuery interface {
	GetEntity(id uint64) (Entity, bool)
	GetEntityByPersistID(persistID string) (Entity, bool)
	RangeEntities(func(entity Entity) bool)
}

type EntityReverseQuery interface {
	ReverseRangeEntities(func(entity Entity) bool)
}

type EntityCountQuery interface {
	GetEntityCount() int
}

type EntityMgr interface {
	EntityQuery
	AddEntity(entity Entity)
	RemoveEntity(id uint64)
}

type EntityMgrEvents interface {
	EventEntityMgrAddEntity() IEvent
	EventEntityMgrRemoveEntity() IEvent
	EventEntityMgrEntityAddComponents() IEvent
	EventEntityMgrEntityRemoveComponent() IEvent
	eventEntityMgrNotifyECTreeRemoveEntity() IEvent
}
