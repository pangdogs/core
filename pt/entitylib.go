package pt

import (
	"fmt"
)

// EntityLib 实体原型库
type EntityLib interface {
	// Register 注册实体原型
	Register(prototype string, compTags []string)

	// Unregister 取消注册实体原型
	Unregister(prototype string)

	// Get 获取实体原型
	Get(prototype string) EntityPt

	// Range 遍历所有已注册的实体原型
	Range(fun func(entityPt EntityPt) bool)
}

// NewEntityLib 创建实体原型库
func NewEntityLib() EntityLib {
	lib := &_EntityLib{}
	lib.init()
	return lib
}

type _EntityLib struct {
	entityPtMap map[string]EntityPt
}

func (lib *_EntityLib) init() {
	if lib.entityPtMap == nil {
		lib.entityPtMap = map[string]EntityPt{}
	}
}

func (lib *_EntityLib) Register(prototype string, compTags []string) {
	_, ok := lib.entityPtMap[prototype]
	if ok {
		panic(fmt.Errorf("repeated register entity '%s' invalid", prototype))
	}

	entityPt := EntityPt{
		Prototype: prototype,
	}

	for i := range compTags {
		entityPt.compPts = append(entityPt.compPts, GetComponentPt(compTags[i]))
	}

	lib.entityPtMap[prototype] = entityPt
}

func (lib *_EntityLib) Unregister(prototype string) {
	delete(lib.entityPtMap, prototype)
}

func (lib *_EntityLib) Get(prototype string) EntityPt {
	entityPt, ok := lib.entityPtMap[prototype]
	if !ok {
		panic(fmt.Errorf("entity '%s' not registered invalid", prototype))
	}

	return entityPt
}

func (lib *_EntityLib) Range(fun func(entityPt EntityPt) bool) {
	if fun == nil {
		return
	}

	for _, entityPt := range lib.entityPtMap {
		if !fun(entityPt) {
			return
		}
	}
}
