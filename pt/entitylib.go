package pt

import (
	"fmt"
)

// EntityLib 实体原型库
type EntityLib interface {
	// Register 注册实体原型。
	//
	//	@param prototype 实体原型名称。
	//	@param compTags 组件标签列表。
	Register(prototype string, compTags []string)

	// Unregister 取消注册实体原型。
	//
	//	@param prototype 实体原型名称。
	Unregister(prototype string)

	// Get 获取实体原型
	//
	//	@param prototype 实体原型名称。
	//	@return 实体原型。
	//	@return 是否存在。
	Get(prototype string) (EntityPt, bool)

	// Range 遍历所有已注册的实体原型
	//
	//	@param fun 遍历函数。
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
		panic(fmt.Errorf("entity '%s' is already registered", prototype))
	}

	entityPt := EntityPt{
		Prototype: prototype,
	}

	for i := range compTags {
		compPt, ok := GetComponentPt(compTags[i])
		if !ok {
			panic(fmt.Errorf("entity '%s' component '%s' not registered", prototype, compTags[i]))
		}
		entityPt.compPts = append(entityPt.compPts, compPt)
	}

	lib.entityPtMap[prototype] = entityPt
}

func (lib *_EntityLib) Unregister(prototype string) {
	delete(lib.entityPtMap, prototype)
}

func (lib *_EntityLib) Get(prototype string) (EntityPt, bool) {
	entityPt, ok := lib.entityPtMap[prototype]
	return entityPt, ok
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
