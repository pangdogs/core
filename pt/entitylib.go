package pt

import (
	"fmt"
	"sync"
)

// EntityLib 实体原型库
type EntityLib interface {
	// Register 注册实体原型。
	//
	//	@param prototype 实体原型名称。
	//	@param compImpls 组件实现列表。
	Register(prototype string, compImpls []string)

	// Deregister 取消注册实体原型。
	//
	//	@param prototype 实体原型名称。
	Deregister(prototype string)

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
	sync.RWMutex
	entityPtMap  map[string]*EntityPt
	entityPtList []*EntityPt
}

func (lib *_EntityLib) init() {
	lib.entityPtMap = map[string]*EntityPt{}
}

// Register 注册实体原型。
//
//	@param prototype 实体原型名称。
//	@param compImpls 组件实现列表。
func (lib *_EntityLib) Register(prototype string, compImpls []string) {
	lib.Lock()
	defer lib.Unlock()

	_, ok := lib.entityPtMap[prototype]
	if ok {
		panic(fmt.Errorf("%w: entity %q is already registered", ErrPt, prototype))
	}

	entityPt := &EntityPt{
		Prototype: prototype,
	}

	for i := range compImpls {
		compPt, ok := AccessComponent(compImpls[i])
		if !ok {
			panic(fmt.Errorf("%w: entity %q component %q not registered", ErrPt, prototype, compImpls[i]))
		}
		entityPt.compPts = append(entityPt.compPts, compPt)
	}

	lib.entityPtMap[prototype] = entityPt
	lib.entityPtList = append(lib.entityPtList, entityPt)
}

// Deregister 取消注册实体原型。
//
//	@param prototype 实体原型名称。
func (lib *_EntityLib) Deregister(prototype string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.entityPtMap, prototype)

	for i, entityPt := range lib.entityPtList {
		if entityPt.Prototype == prototype {
			lib.entityPtList = append(lib.entityPtList[:i], lib.entityPtList[i+1:]...)
			return
		}
	}
}

// Get 获取实体原型
//
//	@param prototype 实体原型名称。
//	@return 实体原型。
//	@return 是否存在。
func (lib *_EntityLib) Get(prototype string) (EntityPt, bool) {
	lib.RLock()
	defer lib.RUnlock()

	entityPt, ok := lib.entityPtMap[prototype]
	if !ok {
		return EntityPt{}, false
	}

	return *entityPt, ok
}

// Range 遍历所有已注册的实体原型
//
//	@param fun 遍历函数。
func (lib *_EntityLib) Range(fun func(entityPt EntityPt) bool) {
	if fun == nil {
		return
	}

	lib.RLock()
	entityPtList := append(make([]*EntityPt, 0, len(lib.entityPtList)), lib.entityPtList...)
	lib.RUnlock()

	for _, entityPt := range entityPtList {
		if !fun(*entityPt) {
			return
		}
	}
}
