package core

import "fmt"

// GetEntity ...
func (servCtx *_ServiceContextBehavior) GetEntity(id int64) (Entity, bool) {
	entity, ok := servCtx.entityMap.Load(id)
	if !ok {
		return nil, false
	}
	return entity.(Entity), true
}

// RangeEntities ...
func (servCtx *_ServiceContextBehavior) RangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	servCtx.entityMap.Range(func(key, value interface{}) bool {
		return fun(value.(Entity))
	})
}

// AddEntity ...
func (servCtx *_ServiceContextBehavior) AddEntity(entity Entity) {
	if entity == nil {
		panic("nil entity")
	}

	if entity.GetID() <= 0 {
		panic("entity id equal 0 invalid")
	}

	if _, loaded := servCtx.entityMap.LoadOrStore(entity.GetID(), entity); loaded {
		panic(fmt.Errorf("repeated entity '{%d}' in this service context", entity.GetID()))
	}
}

// RemoveEntity ...
func (servCtx *_ServiceContextBehavior) RemoveEntity(id int64) {
	servCtx.entityMap.Delete(id)
}
