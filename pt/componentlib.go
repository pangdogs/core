package pt

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
	"github.com/elliotchance/pie/v2"
	"reflect"
	"slices"
	"sync"
)

// ComponentLib 组件原型库
type ComponentLib interface {
	// Declare 声明组件原型
	Declare(comp any, aliases ...string) ComponentPT
	// Undeclare 取消声明组件原型
	Undeclare(name string)
	// Get 获取组件原型
	Get(name string) (ComponentPT, bool)
	// GetAlias 使用别名获取组件原型
	GetAlias(alias string) []ComponentPT
	// Range 遍历所有已注册的组件原型
	Range(fun generic.Func1[ComponentPT, bool])
}

var compLib = NewComponentLib()

// DefaultComponentLib 默认组件库
func DefaultComponentLib() ComponentLib {
	return compLib
}

// NewComponentLib 创建组件原型库
func NewComponentLib() ComponentLib {
	return &_ComponentLib{
		compMap:  map[string]*ComponentPT{},
		aliasMap: map[string][]*ComponentPT{},
	}
}

type _ComponentLib struct {
	sync.RWMutex
	compMap  map[string]*ComponentPT
	compList []*ComponentPT
	aliasMap map[string][]*ComponentPT
}

// Declare 声明组件原型
func (lib *_ComponentLib) Declare(comp any, aliases ...string) ComponentPT {
	if comp == nil {
		panic(fmt.Errorf("%w: %w: comp is nil", ErrPt, exception.ErrArgs))
	}

	if tfComp, ok := comp.(reflect.Type); ok {
		return lib.register(tfComp, aliases)
	} else {
		return lib.register(reflect.TypeOf(comp), aliases)
	}
}

// Undeclare 取消声明组件原型
func (lib *_ComponentLib) Undeclare(name string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.compMap, name)

	lib.compList = slices.DeleteFunc(lib.compList, func(pt *ComponentPT) bool {
		return pt.Name == name
	})

	lib.clearAliases(name)
}

// Get 获取组件原型
func (lib *_ComponentLib) Get(name string) (ComponentPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	comp, ok := lib.compMap[name]
	if !ok {
		return ComponentPT{}, false
	}

	return *comp, ok
}

// GetAlias 使用别名获取组件原型
func (lib *_ComponentLib) GetAlias(alias string) []ComponentPT {
	lib.RLock()
	defer lib.RUnlock()

	comps := lib.aliasMap[alias]
	ret := make([]ComponentPT, 0, len(comps))

	for _, comp := range comps {
		ret = append(ret, *comp)
	}

	return ret
}

// Range 遍历所有已注册的组件原型
func (lib *_ComponentLib) Range(fun generic.Func1[ComponentPT, bool]) {
	lib.RLock()
	copied := slices.Clone(lib.compList)
	lib.RUnlock()

	for i := range copied {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}

func (lib *_ComponentLib) register(tfComp reflect.Type, aliases []string) ComponentPT {
	lib.Lock()
	defer lib.Unlock()

	for tfComp.Kind() == reflect.Pointer || tfComp.Kind() == reflect.Interface {
		tfComp = tfComp.Elem()
	}

	if tfComp.Name() == "" {
		panic(fmt.Errorf("%w: anonymous component not allowed", ErrPt))
	}

	compName := types.AnyFullName(tfComp)

	if !reflect.PointerTo(tfComp).Implements(reflect.TypeOf((*ec.Component)(nil)).Elem()) {
		panic(fmt.Errorf("%w: component %q not implement ec.Component", ErrPt, compName))
	}

	comp, ok := lib.compMap[compName]
	if ok {
		lib.addAliases(comp, aliases)
		return *comp
	}

	comp = &ComponentPT{
		Name:  compName,
		RType: tfComp,
	}

	lib.compMap[compName] = comp
	lib.compList = append(lib.compList, comp)
	lib.addAliases(comp, aliases)

	return *comp
}

func (lib *_ComponentLib) addAliases(comp *ComponentPT, aliases []string) {
	for _, alias := range aliases {
		if !pie.Any(lib.aliasMap[alias], func(pt *ComponentPT) bool {
			return pt.Name == comp.Name
		}) {
			lib.aliasMap[alias] = append(lib.aliasMap[alias], comp)
		}
	}
}

func (lib *_ComponentLib) clearAliases(compName string) {
	for alias, comps := range lib.aliasMap {
		if !pie.Any(comps, func(pt *ComponentPT) bool {
			return pt.Name == compName
		}) {
			continue
		}
		lib.aliasMap[alias] = pie.Filter(comps, func(pt *ComponentPT) bool {
			return pt.Name != compName
		})
	}
}
