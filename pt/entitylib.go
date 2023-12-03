package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/generic"
	"sync"
)

// EntityLib 实体原型库
type EntityLib interface {
	EntityPTProvider
	// Register 注册实体原型
	Register(prototype string, comps ...any) EntityPT
	// Declare 声明实体原型，要求组件实例已注册
	Declare(prototype string, compImpls ...string) EntityPT
	// Deregister 取消注册实体原型
	Deregister(prototype string)
	// Get 获取实体原型
	Get(prototype string) (EntityPT, bool)
	// Range 遍历所有已注册的实体原型
	Range(fun generic.Func1[EntityPT, bool])
}

// CompWithName 注册实体原型并指定名称
type CompWithName struct {
	Comp any
	Name string
}

var entityLib = NewEntityLib(DefaultComponentLib())

// DefaultEntityLib 默认实体库
func DefaultEntityLib() EntityLib {
	return entityLib
}

// NewEntityLib 创建实体原型库
func NewEntityLib(compLib ComponentLib) EntityLib {
	if compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", ErrPt, exception.ErrArgs))
	}

	return &_EntityLib{
		compLib:   compLib,
		entityMap: map[string]*EntityPT{},
	}
}

type _EntityLib struct {
	sync.RWMutex
	compLib    ComponentLib
	entityMap  map[string]*EntityPT
	entityList []*EntityPT
}

// GetEntityLib 获取实体原型库
func (lib *_EntityLib) GetEntityLib() EntityLib {
	return lib
}

// Register 注册实体原型
func (lib *_EntityLib) Register(prototype string, comps ...any) EntityPT {
	lib.Lock()
	defer lib.Unlock()

	_, ok := lib.entityMap[prototype]
	if ok {
		panic(fmt.Errorf("%w: entity %q is already registered", ErrPt, prototype))
	}

	entity := &EntityPT{
		Prototype: prototype,
	}

	for _, comp := range comps {
		var compWithName CompWithName

		switch c := comp.(type) {
		case CompWithName:
			compWithName = c
		case *CompWithName:
			compWithName = *c
		default:
			compWithName = CompWithName{
				Comp: c,
			}
		}

		compPT := lib.compLib.Register(compWithName.Comp, compWithName.Name)
		entity.comps = append(entity.comps, compPT)
	}

	lib.entityMap[prototype] = entity
	lib.entityList = append(lib.entityList, entity)

	return *entity
}

// Declare 声明实体原型，要求组件实例已注册
func (lib *_EntityLib) Declare(prototype string, compImpls ...string) EntityPT {
	lib.Lock()
	defer lib.Unlock()

	_, ok := lib.entityMap[prototype]
	if ok {
		panic(fmt.Errorf("%w: entity %q is already registered", ErrPt, prototype))
	}

	entity := &EntityPT{
		Prototype: prototype,
	}

	for _, compImpl := range compImpls {
		comp, ok := lib.compLib.Get(compImpl)
		if !ok {
			panic(fmt.Errorf("%w: entity %q component %q was not registered", ErrPt, prototype, compImpl))
		}
		entity.comps = append(entity.comps, comp)
	}

	lib.entityMap[prototype] = entity
	lib.entityList = append(lib.entityList, entity)

	return *entity
}

// Deregister 取消注册实体原型
func (lib *_EntityLib) Deregister(prototype string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.entityMap, prototype)

	for i, entity := range lib.entityList {
		if entity.Prototype == prototype {
			lib.entityList = append(lib.entityList[:i], lib.entityList[i+1:]...)
			return
		}
	}
}

// Get 获取实体原型
func (lib *_EntityLib) Get(prototype string) (EntityPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	entity, ok := lib.entityMap[prototype]
	if !ok {
		return EntityPT{}, false
	}

	return *entity, ok
}

// Range 遍历所有已注册的实体原型
func (lib *_EntityLib) Range(fun generic.Func1[EntityPT, bool]) {
	lib.RLock()
	entityList := append(make([]*EntityPT, 0, len(lib.entityList)), lib.entityList...)
	lib.RUnlock()

	for _, entity := range entityList {
		if !fun.Exec(*entity) {
			return
		}
	}
}
