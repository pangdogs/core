package pt

import (
	"fmt"
	"sync"
)

type EntityLib interface {
	Register(prototype string, compTags []string)
	Get(prototype string) EntityPt
	Range(fun func(entityPt EntityPt) bool)
}

func NewEntityLib() EntityLib {
	lib := &_EntityLib{}
	lib.init()
	return lib
}

type _EntityLib struct {
	entityPtMap map[string]EntityPt
	mutex       sync.RWMutex
}

func (lib *_EntityLib) init() {
	if lib.entityPtMap == nil {
		lib.entityPtMap = map[string]EntityPt{}
	}
}

func (lib *_EntityLib) Register(prototype string, compTags []string) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	_, ok := lib.entityPtMap[prototype]
	if ok {
		panic(fmt.Errorf("repeated register entity '%s' invalid", prototype))
	}

	entityPt := EntityPt{
		Prototype: prototype,
	}

	for i := range compTags {
		entityPt.compPts = append(entityPt.compPts, GetCompPt(compTags[i]))
	}

	lib.entityPtMap[prototype] = entityPt
}

func (lib *_EntityLib) Get(prototype string) EntityPt {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	entityPt, ok := lib.entityPtMap[prototype]
	if !ok {
		panic(fmt.Errorf("entity '%s' not registered invalid", prototype))
	}

	return entityPt
}

func (lib *_EntityLib) Range(fun func(entityPt EntityPt) bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	if fun == nil {
		return
	}

	for _, entityPt := range lib.entityPtMap {
		if !fun(entityPt) {
			return
		}
	}
}
