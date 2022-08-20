package core

import "fmt"

func (servCtx *ServiceContextBehavior) GetEntity(id uint64) (Entity, bool) {
	entity, ok := servCtx.entityMap.Load(id)
	if !ok {
		return nil, false
	}
	return entity.(Entity), true
}

func (servCtx *ServiceContextBehavior) RangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	servCtx.entityMap.Range(func(key, value interface{}) bool {
		return fun(value.(Entity))
	})
}

func (servCtx *ServiceContextBehavior) AddEntity(entity Entity) {
	if entity == nil {
		panic("nil entity")
	}

	if entity.GetID() <= 0 {
		panic("entity id equal 0 invalid")
	}

	if _, loaded := servCtx.entityMap.LoadOrStore(entity.GetID(), entity); loaded {
		panic(fmt.Errorf("repeated entity '{%d}' in this serv context", entity.GetID()))
	}
}

func (servCtx *ServiceContextBehavior) RemoveEntity(id uint64) {
	servCtx.entityMap.Delete(id)
}
