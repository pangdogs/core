package pt

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
	"slices"
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

	retry:
		switch pt := comp.(type) {
		case _CompAlias:
			ci.Alias = pt.Alias
			comp = pt.Comp
			goto retry
		case string:
			compPT, ok := lib.compLib.Get(pt)
			if !ok {
				panic(fmt.Errorf("%w: entity %q component %q was not registered", ErrPt, prototype, pt))
			}
			ci.PT = compPT
		default:
			ci.PT = lib.compLib.Register(pt)
		}

		if ci.Alias == "" {
			ci.Alias = ci.PT.Name
		}

		entity.compInfos = append(entity.compInfos, ci)
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

	lib.entityList = slices.DeleteFunc(lib.entityList, func(pt *EntityPT) bool {
		return pt.Prototype == prototype
	})
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
	copied := slices.Clone(lib.entityList)
	lib.RUnlock()

	for i := range copied {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}
