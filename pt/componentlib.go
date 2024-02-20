package pt

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
	"github.com/elliotchance/pie/v2"
	"reflect"
	"sync"
)

// ComponentLib 组件原型库
type ComponentLib interface {
	// Register 注册组件原型
	Register(comp any, aliases ...string) ComponentPT
	// Deregister 取消注册组件原型
	Deregister(name string)
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

// Register 注册组件原型
func (lib *_ComponentLib) Register(comp any, aliases ...string) ComponentPT {
	if comp == nil {
		panic(fmt.Errorf("%w: %w: comp is nil", ErrPt, exception.ErrArgs))
	}

	if tfComp, ok := comp.(reflect.Type); ok {
		return lib.register(tfComp, aliases)
	} else {
		return lib.register(reflect.TypeOf(comp), aliases)
	}
}

// Deregister 取消注册组件原型
func (lib *_ComponentLib) Deregister(name string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.compMap, name)

	for i, comp := range lib.compList {
		if comp.Name == name {
			lib.compList = append(lib.compList[:i], lib.compList[i+1:]...)
			break
		}
	}

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
	compList := append(make([]*ComponentPT, 0, len(lib.compList)), lib.compList...)
	lib.RUnlock()

	for _, comp := range compList {
		if !fun.Exec(*comp) {
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
		Name:   compName,
		tfComp: tfComp,
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
