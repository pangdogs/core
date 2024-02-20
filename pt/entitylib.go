package pt

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
	"sync"
)

type _CompAlias struct {
	Comp  any
	Alias string
}

// CompAlias 组件与别名，用于注册实体原型时自定义组件别名
func CompAlias(comp any, alias string) _CompAlias {
	return _CompAlias{
		Comp:  comp,
		Alias: alias,
	}
}

// CompInterface 组件与接口，用于注册实体原型时使用接口名作为别名
func CompInterface[FACE any](comp any) _CompAlias {
	return _CompAlias{
		Comp:  comp,
		Alias: types.FullName[FACE](),
	}
}

// EntityLib 实体原型库
type EntityLib interface {
	EntityPTProvider
	// Register 注册实体原型
	Register(prototype string, comps ...any) EntityPT
	// Declare 声明实体原型，要求组件实例已注册
	Declare(prototype string, compNames ...string) EntityPT
	// Deregister 取消注册实体原型
	Deregister(prototype string)
	// Get 获取实体原型
	Get(prototype string) (EntityPT, bool)
	// Range 遍历所有已注册的实体原型
	Range(fun generic.Func1[EntityPT, bool])
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
		var ci _CompInfo

		switch pt := comp.(type) {
		case _CompAlias:
			ci.Alias = pt.Alias
			ci.PT = lib.compLib.Register(pt.Comp, pt.Alias)
		default:
			ci.PT = lib.compLib.Register(pt)
			ci.Alias = ci.PT.Name
		}

		entity.compInfos = append(entity.compInfos, ci)
	}

	lib.entityMap[prototype] = entity
	lib.entityList = append(lib.entityList, entity)

	return *entity
}

// Declare 声明实体原型，要求组件实例已注册
func (lib *_EntityLib) Declare(prototype string, compNames ...string) EntityPT {
	lib.Lock()
	defer lib.Unlock()

	_, ok := lib.entityMap[prototype]
	if ok {
		panic(fmt.Errorf("%w: entity %q is already registered", ErrPt, prototype))
	}

	entity := &EntityPT{
		Prototype: prototype,
	}

	for _, compName := range compNames {
		compPT, ok := lib.compLib.Get(compName)
		if !ok {
			panic(fmt.Errorf("%w: entity %q component %q was not registered", ErrPt, prototype, compName))
		}
		entity.compInfos = append(entity.compInfos, _CompInfo{
			PT:    compPT,
			Alias: compPT.Name,
		})
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
