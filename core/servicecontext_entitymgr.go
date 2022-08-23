package core

import "fmt"

func (servCtx *_ServiceContextBehavior) GetEntity(id uint64) (Entity, bool) {
	entity, ok := servCtx.entityMap.Load(id)
	if !ok {
		return nil, false
	}
	return entity.(Entity), true
}

func (servCtx *_ServiceContextBehavior) GetEntityByPersistID(persistID string) (Entity, bool) {
	entity, ok := servCtx.persistentEntityMap.Load(persistID)
	if !ok {
		return nil, false
	}
	return entity.(Entity), true
}

func (servCtx *_ServiceContextBehavior) RangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	servCtx.entityMap.Range(func(key, value interface{}) bool {
		return fun(value.(Entity))
	})
}

func (servCtx *_ServiceContextBehavior) AddEntity(entity Entity) {
	if entity == nil {
		panic("nil entity")
	}

	if entity.GetID() <= 0 {
		panic("entity id equal 0 invalid")
	}

	if entity.GetPersistID() != "" {
		if _, loaded := servCtx.persistentEntityMap.LoadOrStore(entity.GetPersistID(), entity); loaded {
			panic(fmt.Errorf("repeated persistent entity '{%s}' in this service context", entity.GetPersistID()))
		}
	}

	if _, loaded := servCtx.entityMap.LoadOrStore(entity.GetID(), entity); loaded {
		if entity.GetPersistID() != "" {
			servCtx.persistentEntityMap.Delete(entity.GetPersistID())
		}
		panic(fmt.Errorf("repeated entity '{%d}' in this service context", entity.GetID()))
	}
}

func (servCtx *_ServiceContextBehavior) RemoveEntity(id uint64) {
	entity, loaded := servCtx.entityMap.LoadAndDelete(id)
	if !loaded {
		return
	}

	persistID := entity.(Entity).GetPersistID()
	if persistID != "" {
		servCtx.persistentEntityMap.Delete(persistID)
	}
}
